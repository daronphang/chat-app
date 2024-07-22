package usecase

import (
	"chat-service/internal/domain"
	"context"
)

type EventBroker interface {
	PublishMessage(ctx context.Context, partitionKey string, topic string, arg interface{}) error
}

type ServerClienter interface {
	DeliverOutboundMsg(ctx context.Context, clientId string, arg domain.Message) error
}

type UseCaseService struct {
	EventBroker 	EventBroker
	ServerClienter 	ServerClienter
}

func NewUseCaseService( eb EventBroker, sc ServerClienter) *UseCaseService {
	return &UseCaseService{EventBroker: eb, ServerClienter: sc}
}
