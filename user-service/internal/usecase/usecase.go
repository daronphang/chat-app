package usecase

import (
	"context"
	"user-service/internal/domain"
)

type UseCaseService struct {
	Repository 			domain.Repository
	ServiceDiscovery	ServiceDiscovery
}

type ServiceDiscovery interface {
	GetServersMetdata(ctx context.Context) ([]domain.ServerMetadata, error)
}

func NewUseCaseService(repo domain.Repository, sc ServiceDiscovery) *UseCaseService {
	return &UseCaseService{Repository: repo, ServiceDiscovery: sc}
}
