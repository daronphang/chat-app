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

func (ws WebSocketer) DeliverOutboundMsg(ctx context.Context, clientId string, arg domain.Message) error {
	if ctx.Err() != nil {
		return ctx.Err()
	}

	client, ok := hub.clients[clientId]
	if !ok {
		return MissingClientError{
			Message: fmt.Sprintf("client %v is not connected to chat server: unable to deliver message %v", clientId, arg.MessageID),
		}
	}

	data, err := json.Marshal(arg)
	if err != nil {
		return err
	}

	// Each client can have multiple devices connected to the same chat server.
	for _, device := range client.devices {
		device.send <- data
	}

	return nil
}