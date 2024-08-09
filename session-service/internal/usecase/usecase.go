package usecase

import (
	pb "protobuf/proto/user"
	"session-service/internal/domain"
)

type UseCaseService struct {
	Repository 		domain.Repository
	UserClient   pb.UserClient
}

func NewUseCaseService(repo domain.Repository, u pb.UserClient) *UseCaseService {
	return &UseCaseService{Repository: repo, UserClient: u}
}
