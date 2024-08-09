package ws

import (
	"chat-service/internal/domain"
	"context"
	"encoding/json"
	"fmt"
)

type WebSocketer struct {}

type MissingClientError struct {
	Message string
}

func (e MissingClientError) Error() string {
	return e.Message
}

func New() WebSocketer {
	return WebSocketer{}
}

func (ws WebSocketer) SendEventToClient(ctx context.Context, clientID string, event domain.Event, payload interface{}) error {
	if ctx.Err() != nil {
		return ctx.Err()
	}
	
	client, ok := hub.clients[clientID]
	if !ok {
		return MissingClientError{
			Message: fmt.Sprintf("client %v is not connected to chat server, unable to send event: %v", clientID, payload),
		}
	}

	data, err := json.Marshal(domain.WebSocketEvent{
		Event: event,
		Data: payload,
	})
	if err != nil {
		return err
	}

	for _, device := range client.devices {
		device.send <- data
	}

	return nil
}