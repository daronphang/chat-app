package usecase

import (
	"context"
	"user-service/internal/domain"

	"github.com/google/uuid"
)

func (uc *UseCaseService) Signup(ctx context.Context, arg domain.NewUser) (domain.UserMetadata, error) {
	arg.UserID = uuid.NewString()
	rv, err := uc.Repository.CreateUser(ctx, arg)
	if err != nil {
		return domain.UserMetadata{}, err
	}

	// Create user topic.
	if err := uc.EventBroker.CreateUserTopic(ctx, arg.UserID); err != nil {
		return domain.UserMetadata{}, err
	}

	return rv, nil
}

func (uc *UseCaseService) Login(ctx context.Context, arg domain.UserCredentials) (domain.UserMetadata, error) {
	um, err := uc.Repository.GetUser(ctx, arg.Email)
	if err != nil {
		return domain.UserMetadata{}, err
	}
	contacts, err := uc.Repository.GetContacts(ctx, um.UserID)
	if err != nil {
		return domain.UserMetadata{}, err
	}
	um.Contacts = contacts
	return um, nil
}

func (uc *UseCaseService) UpdateUser(ctx context.Context, arg domain.UserMetadata) error {
	if err := uc.Repository.UpdateUser(ctx, arg); err != nil {
		return err
	}
	return nil
}

func (uc *UseCaseService) CreateContact(ctx context.Context, arg domain.NewContact) (domain.Contact, error) {
	friend, err := uc.Repository.GetUser(ctx, arg.FriendEmail)
	if err != nil {
		return domain.Contact{}, err
	}

	arg.FriendID = friend.UserID

	if err := uc.Repository.CreateContact(ctx, arg); err != nil {
		return domain.Contact{}, err
	}
	return domain.Contact{
		UserID: friend.UserID,
		Email: friend.Email,
		DisplayName: arg.DisplayName,
	}, nil
}

func (uc *UseCaseService) GetContacts(ctx context.Context, arg string) ([]domain.Contact, error) {
	rv, err := uc.Repository.GetContacts(ctx, arg)
	if err != nil {
		return nil, err
	}
	return rv, nil
}