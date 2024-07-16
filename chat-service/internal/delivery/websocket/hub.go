package ws

import (
	"chat-service/internal/domain"
	"chat-service/internal/usecase"
	cv "chat-service/internal/validator"
	"context"
	"fmt"

	"github.com/gorilla/websocket"
	"go.uber.org/zap"
)

type Client struct {
	hub *Hub
	// Websocket connection.
	conn *websocket.Conn
	// Buffered channel of outbound messages.
	send chan []byte
	// ClientID.
	clientID string
}

type Hub struct {
	// Registered clients.
	clients map[string]*Client

	// Inbound messages from the clients.
	sendMsg chan []byte

	// Register requests from the clients.
	register chan *Client

	// Unregister requests from clients.
	unregister chan *Client
}

func NewHub() *Hub {
	return &Hub{
		sendMsg:  make(chan []byte),
		register:   make(chan *Client),
		unregister: make(chan *Client),
		clients:    make(map[string]*Client),
	}
}

func (h *Hub) Run(ctx context.Context, uc *usecase.UseCaseService) {
	for {
		select {
		case <- ctx.Done():
			return
		case client := <-h.register:
			h.clients[client.clientID] = client
		case client := <-h.unregister:
			if _, ok := h.clients[client.clientID]; ok {
				delete(h.clients, client.clientID)
				close(client.send)
			}
		case msg := <-h.sendMsg:
			// Validate sender message, push to queue and send ack.
			// If msg failed to deliver, message returned to the sender
			// will not include a msgID.
			v := new(domain.Message)
			if err := cv.UnmarshalAndValidate(msg, v); err != nil {
				logger.Error(
					fmt.Sprintf("validation failed for sender message: %v", string(msg)),
				)
				break
			}

			rv, err := uc.PushSenderMsgToQueue(ctx, *v)
			var rm domain.ReceiverMessage
			if err != nil {
				logger.Error(
					"failed to push sender message to queue", 
					zap.String("payload", string(msg)),
					zap.String("trace", err.Error()),
				)
				rm = domain.ReceiverMessage{Message: *v, ReceiverID: v.SenderID}
			} else {
				rm = domain.ReceiverMessage{Message: rv, ReceiverID: rv.SenderID}
			}
			
			if err := uc.SendMsgToClient(ctx, rm); err != nil {
				logger.Error(
					"failed to send ack message to sender",
					zap.String("payload", string(msg)),
					zap.String("trace", err.Error()),
				)
			}
		}
	}
}
