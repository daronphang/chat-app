package usecase

import (
	"context"
	"user-service/internal/domain"
)

type UseCaseService struct {
	Repository 			domain.ExtRepo
	ServiceDiscovery	ServiceDiscovery
}

type ServiceDiscovery interface {
	GetServersMetdata(ctx context.Context) ([]domain.ServerMetadata, error)
}

func NewUseCaseService(repo domain.ExtRepo, sc ServiceDiscovery) *UseCaseService {
	return &UseCaseService{Repository: repo, ServiceDiscovery: sc}
}
