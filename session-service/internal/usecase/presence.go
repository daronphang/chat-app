package usecase

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
	"session-service/internal"
	"session-service/internal/config"
	"session-service/internal/domain"

	"go.uber.org/zap"
	"google.golang.org/protobuf/types/known/wrapperspb"
)

var (
	logger, _ = internal.WireLogger()
)

func (uc *UseCaseService) GetOfflineUsers(ctx context.Context, arg []string) []string {
	offlineUsers := make([]string, 0)
	for _, userID := range arg {
		user, err := uc.Repository.GetUser(ctx, userID)
		if err != nil || user.UserID == "" {
			offlineUsers = append(offlineUsers, userID)
		}
	}
	return offlineUsers
} 

func (uc *UseCaseService) GetOnlineUsers(ctx context.Context, arg []string) []string {
	onlineUsers := make([]string, 0)
	for _, userID := range arg {
		user, _ := uc.Repository.GetUser(ctx, userID)
		if user.UserID == userID {
			onlineUsers = append(onlineUsers, userID)
		}
	}
	return onlineUsers
} 

func (uc *UseCaseService) BroadcastUserPresence(ctx context.Context, arg domain.UserPresence) error {
	// Get all the friends of the user.
	relations, err := uc.UserClient.GetUsersAssociatedToChannel(ctx, &wrapperspb.StringValue{Value: arg.UserID})
	if err != nil {
		return err
	}

	maxGoroutines := 10
	guard := make(chan bool, maxGoroutines)
	cfg, _ := config.ProvideConfig()

	for _, friendID := range relations.UserIds {
		guard <- true
		go func (targetID string) {
			// Check if friend is online.
			user, err := uc.Repository.GetUser(ctx, targetID)
			if err != nil || user.UserID == "" {
				<- guard
				return
			}

			// Send request to user devices.
			// User is assumed to be online since TTL has not expired.
			for _, server := range user.Servers {
				// Ignore error and response.
				queryParams := url.Values{
					"clientId": {arg.UserID},
					"targetId": {targetID},
					"status": {arg.Status},
				}
				_, err := http.Get(fmt.Sprintf("%v/%v?%v", server, cfg.ChatServerAPI, queryParams.Encode()))
				if err != nil {
					logger.Error(
						fmt.Sprintf("unable to broadcast client %v %v status to target %v", arg.UserID, arg.Status, targetID),
						zap.String("trace", err.Error()),
					)
				}
			}
			<- guard
		}(friendID)
	}

	// Wait for goroutines to end, else context will be canceled.
	for {
		if len(guard) == 0 {
			break
		}
	}

	return nil
}

func (uc *UseCaseService) ClientHeartbeat(ctx context.Context, arg domain.HeartbeatRequest) error {
	if err := uc.Repository.UpdateUser(ctx, arg); err != nil {
		return err
	}
	return nil
}