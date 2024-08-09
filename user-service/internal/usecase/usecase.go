package usecase

import (
	"context"
	"user-service/internal/domain"
)

type UseCaseService struct {
	Repository 			domain.ExtRepo
	ServiceDiscovery	ServiceDiscovery
	EventBroker			EventBroker
}

type ServiceDiscovery interface {
	GetServersMetdata(ctx context.Context) ([]domain.ServerMetadata, error)
}

type EventBroker interface {
	CreateUserTopic(ctx context.Context, topic string) error 
}

func NewUseCaseService(repo domain.ExtRepo, sc ServiceDiscovery, eb EventBroker) *UseCaseService {
	return &UseCaseService{Repository: repo, ServiceDiscovery: sc, EventBroker: eb}
}
