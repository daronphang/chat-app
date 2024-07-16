package ws

import (
	"chat-service/internal/domain"
	"context"
	"encoding/json"
	"fmt"
)



type WebSocketer struct {
	Hub *Hub
}

type MissingClientError struct {
	Message string
}

func (e MissingClientError) Error() string {
	return e.Message
}

func New() WebSocketer {
	return WebSocketer{Hub: NewHub()}
}


func (ws WebSocketer) SendMsgToClient(ctx context.Context, arg domain.ReceiverMessage) error {
	// Chat servers communicate with each other via HTTP calls.
	// For sender, closing of channel is not needed once it is done.
	// When websocket connection is closed, the client will be removed.
	if ctx.Err() != nil {
		return ctx.Err()
	}
	client, ok := ws.Hub.clients[arg.ReceiverID]
	if !ok {
		return MissingClientError{
			Message: fmt.Sprintf("client %v is not connected to chat service: %v", arg.ReceiverID, arg.Message.MessageID),
		}
	}

	data, err := json.Marshal(arg.Message)
	if err != nil {
		return err
	}

	client.send <- data
	return nil
}

