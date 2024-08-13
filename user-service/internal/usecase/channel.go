package usecase

import (
	"context"
	"slices"
	"strings"
	"time"
	"user-service/internal/domain"

	"github.com/google/uuid"
)

func (uc *UseCaseService) CreateChannel(ctx context.Context, arg domain.NewChannel) (domain.Channel, error) {
	arg.CreatedAt = time.Now().UTC().Format(time.RFC3339)
	if len(arg.UserIDs) == 2 {
		slices.Sort(arg.UserIDs)
		arg.ChannelID = strings.Join(arg.UserIDs, "")
	} else {
		arg.ChannelID = uuid.NewString()
	}

	closure := uc.Repository.ExecWithTx(ctx, func(qtx domain.Repository) (interface{}, error) {
		if err := qtx.CreateUserToChannelAssociation(ctx, arg); err != nil {
			return nil, err
		}

		if len(arg.UserIDs) == 2 {
			return nil, nil
		}

		if err := qtx.CreateGroupChannel(ctx, arg); err != nil {
			return nil, err
		}
		return nil, nil
	})
	_, err := closure()
	if err != nil {
		return domain.Channel{}, err
	}

	channel := domain.Channel{
		ChannelID: arg.ChannelID,
		ChannelName: arg.ChannelName,
		CreatedAt: arg.CreatedAt,
		UserIDs: arg.UserIDs,
	}
	return channel, nil
}

func (uc *UseCaseService) GetChannelsAssociatedToUser(ctx context.Context, arg string) ([]domain.Channel, error) {
	rv, err := uc.Repository.GetChannelsAssociatedToUser(ctx, arg)
	if err != nil {
		return nil, err
	}
	return rv, nil
}

func (uc *UseCaseService) GetUsersAssociatedToChannel(ctx context.Context, arg string) ([]domain.UserContact, error) {
	rv, err := uc.Repository.GetUsersAssociatedToChannel(ctx, arg)
	if err != nil {
		return nil, err
	}
	return rv, nil
}

func (uc *UseCaseService) BroadcastChannelEventToUsers(ctx context.Context, arg []string) error {
	
	return nil
}

func (uc *UseCaseService) GetUsersAssociatedToTargetUser(ctx context.Context, arg string) ([]string, error) {
	rv, err := uc.Repository.GetUsersAssociatedToTargetUser(ctx, arg)
	if err != nil {
		return nil, err
	}
	return rv, nil
}