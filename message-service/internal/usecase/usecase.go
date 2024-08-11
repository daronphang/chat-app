package usecase

import (
	"context"
	"message-service/internal/domain"

	pb "protobuf/proto/notification"
)

type EventBroker interface {
	PublishEventToUserQueue(ctx context.Context, userID string, arg domain.BaseEvent) error
}

type UseCaseService struct {
	Repository 				domain.Repository
	EventBroker				EventBroker
	NotificationClient   	pb.NotificationClient
}

func NewUseCaseService(eb EventBroker, repo domain.Repository, nc pb.NotificationClient) *UseCaseService {
	return &UseCaseService{EventBroker: eb, Repository: repo, NotificationClient: nc}
}
