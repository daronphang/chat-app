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
	channel := domain.Channel{
		ChannelID: "",
		ChannelName: arg.ChannelName,
		UserIDs: arg.UserIDs,
		CreatedAt: time.Now().Format(time.RFC3339),
	}

	if len(channel.UserIDs) == 2 {
		slices.Sort(channel.UserIDs)
		channel.ChannelID = strings.Join(arg.UserIDs, "")
	} else {
		channel.ChannelID = uuid.NewString()
	}

	closure := uc.Repository.ExecWithTx(ctx, func(qtx domain.Repository) (interface{}, error) {
		if err := qtx.CreateUserToChannelAssociation(ctx, channel); err != nil {
			return nil, err
		}

		if len(arg.UserIDs) == 2 {
			return nil, nil
		}

		if err := qtx.CreateGroupChannel(ctx, channel); err != nil {
			return nil, err
		}
		return nil, nil
	})
	_, err := closure()
	if err != nil {
		return domain.Channel{}, err
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

func (uc *UseCaseService) GetUsersAssociatedToTargetUser(ctx context.Context, arg string) ([]string, error) {
	rv, err := uc.Repository.GetUsersAssociatedToTargetUser(ctx, arg)
	if err != nil {
		return nil, err
	}
	return rv, nil
}

func (uc *UseCaseService) AddGroupMembers(ctx context.Context, arg domain.GroupMembers) error {
	// Check if group exists.
	group, err := uc.Repository.GetGroupChannel(ctx, arg.ChannelID)
	if err != nil {
		return err
	}

	channel := domain.Channel{
		ChannelID: arg.ChannelID,
		ChannelName: group.ChannelName,
		UserIDs: arg.UserIDs,
		CreatedAt: time.Now().Format(time.RFC3339),
	}

	if err := uc.Repository.CreateUserToChannelAssociation(ctx, channel); err != nil {
		return err
	}
	return nil
}

func (uc *UseCaseService) RemoveGroupMembers(ctx context.Context, arg domain.GroupMembers) error {
	if err := uc.Repository.RemoveGroupMembers(ctx, arg); err != nil {
		return err
	}
	return nil
}

func (uc *UseCaseService) LeaveGroup(ctx context.Context, arg domain.GroupMembers) error {
	if err := uc.Repository.RemoveGroupMembers(ctx, arg); err != nil {
		return err
	}
	return nil
}

func (uc *UseCaseService) RemoveGroup(ctx context.Context, arg domain.AdminGroupMember) error {
	// TODO: Validate if admin.

	if err := uc.Repository.RemoveGroup(ctx, arg.ChannelID); err != nil {
		return err
	}
	return nil
}

func (uc *UseCaseService) UpdateLastReadMessage(ctx context.Context, arg domain.LastReadMessage) error {
	if err := uc.Repository.UpdateLastReadMessage(ctx, arg); err != nil {
		return err
	}
	return nil
}