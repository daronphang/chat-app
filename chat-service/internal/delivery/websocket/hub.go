package ws

import (
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
	send 		chan []byte // Buffered channel of outbound data to client.
	consumer	*kafka.KafkaConsumer
}

type Client struct {
	clientID 	string
	devices 	map[string]*Device
}

type Hub struct {
	clients 			map[string]*Client
	receive 			chan []byte	// Buffered channel of inbound data from client.
	registerDevice 		chan *Device
	unregisterDevice 	chan *Device
	usecase 			*usecase.UseCaseService
}

func ProvideHub() *Hub {
	return hub
}

func NewHub(uc *usecase.UseCaseService) *Hub {
	hub = &Hub{
		receive:  			make(chan []byte),
		registerDevice:   	make(chan *Device),
		unregisterDevice: 	make(chan *Device),
		clients:    		make(map[string]*Client),
		usecase: 			uc,
	}
	return hub
}

func (h *Hub) Close() {
	for _, client := range h.clients {
		for _, device := range client.devices {
			device.consumer.Close()
		}
	}
}

func (h *Hub) handleReceiveMessage(ctx context.Context, msg []byte) {
	// Validate message from client.
	v := new(domain.Message)
	if err := cv.UnmarshalAndValidate(msg, v); err != nil {
		logger.Error(
			fmt.Sprintf("validation failed for inbound message: %v", string(msg)),
		)
		return
	}

	// Save message.
	rv, err := h.usecase.SaveNewMessage(ctx, *v)
	if err != nil {
		logger.Error(
			"failed to send inbound message", 
			zap.String("payload", string(msg)),
			zap.String("trace", err.Error()),
		)
		return
	}

	// Send acknowledgement back to client.
	event := domain.BaseEvent{
		Event: domain.EventMessage,
		Data: rv,
	}
	if err := h.usecase.SendEventToClient(ctx, rv.SenderID, event); err != nil {
		logger.Error(
			"failed to ack inbound message to sender",
			zap.String("payload", string(msg)),
			zap.String("trace", err.Error()),
		)
	}
}

func (h *Hub) createNewClient(device *Device) {
	client := &Client{
		clientID: device.clientID,
		devices: make(map[string]*Device),
	}
	h.clients[client.clientID] = client	
	client.devices[device.deviceID] = device
}

func (h *Hub) Run(ctx context.Context, brokers []string) {
	for {
		select {
		case <- ctx.Done():
			return
		case device := <-h.registerDevice:
			client, ok := h.clients[device.clientID]
			if !ok {
				h.createNewClient(device)
			} else {
				client.devices[device.deviceID] = device
			}
			// Take note of Kafka dependency in this layer.
			// Each device will consume from the user topic at different 
			// offset values. Hence, a new reader is required.
			device.consumer = kafka.NewConsumer(brokers, device.deviceID, device.clientID)
			go device.consumer.ConsumeFromUserTopic(ctx, h.usecase)
		case device := <-h.unregisterDevice:
			client, ok := h.clients[device.clientID]
			if !ok {
				break
			}
			
			close(device.send)
			device.consumer.Close()

			if len(client.devices) == 1 {
				// Remove client.
				delete(h.clients, client.clientID)
			} else {
				// Remove device as client still has other devices connected.
				delete(client.devices, device.deviceID)
			}
		case msg := <-h.receive:
			// Hub handles all inbound messages for all websocket connections.
			maxGoroutines := 10
			guard := make(chan bool, maxGoroutines)
			guard <- true
			go func() {
				h.handleReceiveMessage(ctx, msg)
				<- guard
			}()
		}
	}
}