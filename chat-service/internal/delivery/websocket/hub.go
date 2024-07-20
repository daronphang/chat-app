package ws

import (
	"chat-service/internal/config"
	"chat-service/internal/delivery/kafka"
	"chat-service/internal/domain"
	"chat-service/internal/usecase"
	cv "chat-service/internal/validator"
	"context"
	"fmt"

	"github.com/gorilla/websocket"
	"go.uber.org/zap"
)

var (
	hub *Hub
)

type Device struct {
	deviceID	string
	clientID 	string
	hub 		*Hub
	conn 		*websocket.Conn
	send 		chan []byte // Buffered channel of outbound messages.
}

type Client struct {
	clientID 	string
	devices 	map[string]*Device
}

type Hub struct {
	clients 			map[string]*Client
	sendMsg 			chan []byte	// Inbound messages from the clients.
	registerDevice 		chan *Device
	unregisterDevice 	chan *Device
	uc 					*usecase.UseCaseService
}

func NewHub(uc *usecase.UseCaseService) *Hub {
	hub = &Hub{
		sendMsg:  			make(chan []byte),
		registerDevice:   	make(chan *Device),
		unregisterDevice: 	make(chan *Device),
		clients:    		make(map[string]*Client),
		uc: 				uc,
	}
	return hub
}

func (h *Hub) handleSenderMsg(ctx context.Context, msg []byte) {
	// Validate sender message, push to queue and send ack.
	// If msg failed to deliver, message returned to the sender
	// will not include a msgID.
	v := new(domain.Message)
	if err := cv.UnmarshalAndValidate(msg, v); err != nil {
		logger.Error(
			fmt.Sprintf("validation failed for sender message: %v", string(msg)),
		)
		return
	}

	rv, err := h.uc.AckSenderMsg(ctx, *v)
	if err != nil {
		logger.Error(
			"failed to ack sender message", 
			zap.String("payload", string(msg)),
			zap.String("trace", err.Error()),
		)
	}
	
	if err := h.uc.SendMsgToClientDevices(ctx, rv.SenderID, rv); err != nil {
		logger.Error(
			"failed to send ack to sender",
			zap.String("payload", string(msg)),
			zap.String("trace", err.Error()),
		)
	}
}

func (h *Hub) createNewClient(ctx context.Context, cfg *config.Config, device *Device) error {
	// For each new client, to consume messages from client topic.
	// One reader opened for all devices.
	// Create topic first.
	tcfg := domain.UserTopicConfig
	tcfg.Topic = device.clientID
	tcfg.ConsumerGroupID = device.deviceID
	if err := kafka.CreateKafkaTopics(cfg, tcfg); err != nil {
		return err
	}

	// Create new client.
	client := &Client{
		clientID: device.clientID,
		devices: make(map[string]*Device),
	}
	h.clients[client.clientID] = client	
	client.devices[device.deviceID] = device

	// Consume from user topic.
	consumer := kafka.NewConsumer(cfg, device.deviceID, device.clientID)
	go consumer.ConsumeMsgFromUserTopic(ctx, h.uc)
	return nil
}

func (h *Hub) Run(ctx context.Context, cfg *config.Config) {
	for {
		select {
		case <- ctx.Done():
			return
		case device := <-h.registerDevice:
			client, ok := h.clients[device.clientID]
			if !ok {
				if err := h.createNewClient(ctx, cfg, device); err != nil {
					logger.Error(
						fmt.Sprintf("unable to create client for %v", device.clientID),
						zap.String("trace", err.Error()),
					)

					// Close websocket connection.
					device.conn.Close()
				}
			} else {
				client.devices[device.deviceID] = device
			}
		case device := <-h.unregisterDevice:
			client, ok := h.clients[device.clientID]
			if !ok {
				break
			}
			close(device.send)

			if len(client.devices) == 1 {
				// Remove client.
				delete(h.clients, client.clientID)
			} else {
				// Remove device as client still has other devices connected.
				delete(client.devices, device.deviceID)
			}
		case msg := <-h.sendMsg:
			maxGoroutines := 10
			guard := make(chan bool, maxGoroutines)
			guard <- true
			go func() {
				h.handleSenderMsg(ctx, msg)
				<- guard
			}()

		}
	}
}