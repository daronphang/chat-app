package usecase

import (
	"context"
	"slices"
	"strings"
	"user-service/internal/domain"

	"github.com/google/uuid"
)

func (uc *UseCaseService) CreateChannel(ctx context.Context, arg domain.NewChannel) (domain.Channel, error) {
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
	return domain.Channel{
		ChannelID: arg.ChannelID,
		ChannelName: arg.ChannelName,
	}, nil
}

func (uc *UseCaseService) GetChannelsAssociatedToUser(ctx context.Context, arg string) ([]domain.Channel, error) {
	rv, err := uc.Repository.GetChannelsAssociatedToUser(ctx, arg)
	if err != nil {
		return nil, err
	}
	return rv, nil
}

func (uc *UseCaseService) GetUsersAssociatedToChannel(ctx context.Context, arg string) ([]string, error) {
	rv, err := uc.Repository.GetUsersAssociatedToChannel(ctx, arg)
	if err != nil {
		return nil, err
	}
	return rv, nil
}