package usecase

import (
	"context"
	"fmt"
	"message-service/internal"
	"message-service/internal/domain"
	"message-service/internal/util"
	"sort"
	"strings"
	"time"

	"github.com/google/uuid"
	"go.uber.org/zap"
)

var (
	logger, _ = internal.WireLogger()
)

func (uc *UseCaseService) GetLatestMessages(ctx context.Context, arg domain.LatestMessagesRequest) ([]domain.Message, error) {
	rv, err := uc.Repository.GetLatestMessages(ctx, arg.ChannelID, arg.LastMessageID)
	if err != nil {
		return nil, err
	}
	return rv, nil
}

func (uc *UseCaseService) SaveMessageAndDeliverToRecipients(ctx context.Context, arg domain.Message) error {
	// Save message in db.
	if err := uc.Repository.CreateMessage(ctx, arg); err != nil {
		return err
	}

	// Get userIds associated to channel.
	userIDs, err := uc.Repository.GetUserIdsAssociatedToChannel(ctx, arg.ChannelID)
	if err != nil {
		return err
	}

	// Push message to the queues of all users in the channel.
	// TODO: Assumption made that the topic is already created.
	// Pushing to queue must be guaranteed.
	maxGoroutines := 10
	guard := make(chan bool, maxGoroutines)
	for _, userID := range userIDs {
		guard <- true
		go func(userID string) {
			errMsg := fmt.Sprintf("failed to push message %v to user %v queue", arg.MessageID, userID)
			err := util.ExpBackoff(
				500 * time.Millisecond,
				2,
				3,
				errMsg,
				func() error {
					return uc.EventBroker.PublishMessage(ctx, userID, userID, arg)
				},
			)
			if err != nil {
				logger.Error(
					errMsg,
					zap.String("trace", err.Error()),
				)
				<- guard
				return
			}

			// If user is offline, to send push notification via queue.
			// Delivery is not guaranteed.
			// resp = fetch()
			if err := uc.MessageBroker.PublishMessage(
				ctx, 
				domain.NotificationQueueConfig.Queue,
				domain.NotificationQueueConfig.RoutingKeys[0],
				arg,
			); err != nil {
				logger.Error(
					fmt.Sprintf("failed to notify user %v for message %v", userID, arg.MessageID),
					zap.String("trace", err.Error()),
				)
			}
			<- guard
		}(userID)
	}
	return nil
}

func (uc *UseCaseService) AddUsersToChannel(ctx context.Context, channelID string, userIDs []string) error {
	if channelID == "" {
		if len(userIDs) == 2 {
			sort.Strings(userIDs)
			channelID = strings.Join(userIDs, "")
		} else {
			channelID = uuid.NewString()
		}
	}
	if err := uc.Repository.AddUserIDsToChannel(ctx, channelID, userIDs); err != nil {
		return err
	}
	return nil
}

func (uc *UseCaseService) JoinGroup(ctx context.Context) {}

func (uc *UseCaseService) LeaveGroup(ctx context.Context) {}

func (uc *UseCaseService) DeleteUserChat(ctx context.Context, client string, channelID string) {}