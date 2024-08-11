package usecase

import (
	"chat-service/internal/domain"
	"context"
)

type EventBroker interface {
	PublishNewMessageToQueue(ctx context.Context, partitionKey string, arg domain.Message) error 
}

type ServerClienter interface {
	SendEventToClient(ctx context.Context, clientID string, event domain.BaseEvent) error 
}

type UseCaseService struct {
	EventBroker 		EventBroker
	ServerClienter 		ServerClienter
}

func NewUseCaseService( eb EventBroker, sc ServerClienter) *UseCaseService {
	return &UseCaseService{EventBroker: eb, ServerClienter: sc}
}
