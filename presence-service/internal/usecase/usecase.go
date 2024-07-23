package usecase

import (
	"presence-service/internal/domain"
	pb "protobuf/message"
)

type UseCaseService struct {
	Repository 		domain.Repository
	MessageClient   pb.MessageClient
}

func NewUseCaseService(repo domain.Repository, mc pb.MessageClient) *UseCaseService {
	return &UseCaseService{Repository: repo, MessageClient: mc}
}
