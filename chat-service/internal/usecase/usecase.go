package usecase

import (
	"chat-service/internal/domain"
	"context"
)

type EventBroker interface {
	PublishMessage(ctx context.Context, arg domain.Message) error
}

type ServerClienter interface {
	SendMsgToClient(ctx context.Context, arg domain.ReceiverMessage) error
}

type UseCaseService struct {
	EventBroker EventBroker
	ServerClienter ServerClienter
}

func NewUseCaseService( eb EventBroker, sc ServerClienter) *UseCaseService {
	return &UseCaseService{EventBroker: eb, ServerClienter: sc}
}
