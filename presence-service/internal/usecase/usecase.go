package usecase

import (
	"presence-service/internal/domain"
	pb "protobuf/proto/user"
)

type UseCaseService struct {
	Repository 		domain.Repository
	UserClient   pb.UserClient
}

func NewUseCaseService(repo domain.Repository, u pb.UserClient) *UseCaseService {
	return &UseCaseService{Repository: repo, UserClient: u}
}
