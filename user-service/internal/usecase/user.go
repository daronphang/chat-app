package usecase

import (
	"context"
	"user-service/internal/domain"

	"github.com/google/uuid"
)

func (uc *UseCaseService) CreateUser(ctx context.Context, arg domain.NewUser) (domain.UserMetadata, error) {
	arg.UserID = uuid.NewString()
	rv, err := uc.Repository.CreateUser(ctx, arg)
	if err != nil {
		return domain.UserMetadata{}, err
	}
	return rv, nil
}

func (uc *UseCaseService) UpdateUser(ctx context.Context, arg domain.UserMetadata) error {
	if err := uc.Repository.UpdateUser(ctx, arg); err != nil {
		return err
	}
	return nil
}

func (uc *UseCaseService) Login(ctx context.Context) (domain.UserMetadata, error) {
	return domain.UserMetadata{}, nil
}

