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

func (ws WebSocketer) SendMsgToClientDevices(ctx context.Context, clientId string, arg domain.Message) error {
	// For sending or receiving message, to broadcast message to all client devices.
	if ctx.Err() != nil {
		return ctx.Err()
	}

	client, ok := hub.clients[clientId]
	if !ok {
		return MissingClientError{
			Message: fmt.Sprintf("client %v is not connected to chat service: %v", clientId, arg.MessageID),
		}
	}

	data, err := json.Marshal(arg)
	if err != nil {
		return err
	}

	for _, device := range client.devices {
		device.send <- data
	}

	return nil
}