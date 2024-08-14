package usecase

import (
	"context"
	"message-service/internal/domain"

	pb "protobuf/proto/session"
)

type EventBroker interface {
	PublishEventToUserQueue(ctx context.Context, userID string, arg domain.BaseEvent) error
}

type UseCaseService struct {
	Repository 				domain.Repository
	EventBroker				EventBroker
	SessionClient   	pb.SessionClient
}

func NewUseCaseService(eb EventBroker, repo domain.Repository, sc pb.SessionClient) *UseCaseService {
	return &UseCaseService{EventBroker: eb, Repository: repo, SessionClient: sc}
}
