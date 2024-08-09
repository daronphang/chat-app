package usecase

import (
	cfg "broadcast-service/internal/config"
	"broadcast-service/internal/domain"
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"sync"

	"go.uber.org/zap"
	"google.golang.org/protobuf/types/known/wrapperspb"
)

func (uc *UseCaseService) BroadcastUserPresenceEvent(ctx context.Context, arg domain.UserPresence) error {
	// Get all the friends of the user.
	users, err := uc.UserClient.GetContacts(ctx, &wrapperspb.StringValue{Value: arg.UserID})
	if err != nil {
		return err
	}
	userIds := make([]string, 0)
	for _, contact := range users.Contacts {
		userIds = append(userIds, contact.UserId)
	}

	// Get online users.
	onlineUsers := uc.GetOnlineUsers(ctx, userIds)

	// Push to chat server for online users.
	maxGoroutines := 10
	guard := make(chan bool, maxGoroutines)
	wg := sync.WaitGroup{}
	cfg, _ := cfg.ProvideConfig()
	for _, targetUser := range onlineUsers {
		guard <- true
		wg.Add(1)

		go func(targetUser domain.UserSession) {
			defer wg.Done()
			for _, server := range targetUser.Servers {
				url := fmt.Sprintf(
					"%v/%v",
					server, 
					cfg.ChatServerAPI.PresencePath,
				)
				payload := domain.UserPresenceEvent{
					TargetID: targetUser.UserID,
					ClientID: arg.UserID,
					Status: arg.Status,
				}
				body, _ := json.Marshal(payload)
				_, err := http.Post(url, "application/json", bytes.NewBuffer(body))
				if err != nil {
					logger.Error(
						"unable to broadcast user presence",
						zap.String("body", string(body)),
						zap.String("trace", err.Error()),
					)
				}
			}
		}(targetUser)
	}
	wg.Wait()
	return nil
}