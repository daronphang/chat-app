package usecase

import (
	"context"
	"errors"
	"user-service/internal/domain"

	"github.com/google/uuid"
)

var (
	errUserNotExist = errors.New("user does not exist")
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
	friends, err := uc.Repository.GetFriends(ctx, um.UserID)
	if err != nil {
		return domain.UserMetadata{}, err
	}
	um.Friends = friends
	return um, nil
}

func (uc *UseCaseService) UpdateUser(ctx context.Context, arg domain.UserMetadata) error {
	if err := uc.Repository.UpdateUser(ctx, arg); err != nil {
		return err
	}
	return nil
}

func (uc *UseCaseService) AddFriend(ctx context.Context, arg domain.NewFriend) (domain.Friend, error) {
	friend, err := uc.Repository.GetUser(ctx, arg.FriendEmail)
	if err != nil {
		return domain.Friend{}, err
	} else if friend.UserID == "" {
		return domain.Friend{}, errUserNotExist
	}

	arg.FriendID = friend.UserID
	if err := uc.Repository.AddFriend(ctx, arg); err != nil {
		return domain.Friend{}, err
	}
	return domain.Friend{
		UserID: friend.UserID,
		Email: friend.Email,
		DisplayName: arg.DisplayName,
	}, nil
}

func (uc *UseCaseService) GetFriends(ctx context.Context, arg string) ([]domain.Friend, error) {
	rv, err := uc.Repository.GetFriends(ctx, arg)
	if err != nil {
		return nil, err
	}
	return rv, nil
}

func (uc *UseCaseService) GetUsersContactsMetadata(ctx context.Context, arg []string) ([]domain.UserContact, error) {
	rv, err := uc.Repository.GetUsersContactsMetadata(ctx, arg)
	if err != nil {
		return nil, err
	}
	return rv, nil
}