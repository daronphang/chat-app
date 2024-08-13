package usecase

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	cfg "notification-service/internal/config"
	"notification-service/internal/domain"
	"sync"

	"go.uber.org/zap"
	"google.golang.org/protobuf/types/known/wrapperspb"
)

func (uc *UseCaseService) BroadcastUserPresenceEvent(ctx context.Context, arg domain.UserPresence) error {
	// Get all the friends of the user.
	// Performance wise, this should not have a huge impact even if the user has many friends.
	// This is because the number of online users should be minimal.
	// Otherwise, this can be improved by running an algorithm to determine the closest friends.
	users, err := uc.UserClient.GetFriends(ctx, &wrapperspb.StringValue{Value: arg.UserID})
	if err != nil {
		return err
	}
	
	// Notify users who are online only.
	maxGoroutines := 10
	guard := make(chan bool, maxGoroutines)
	wg := sync.WaitGroup{}
	cfg, _ := cfg.ProvideConfig()
	for _, friend := range users.Friends {
		guard <- true
		wg.Add(1)
		go func(friendID string) {
			defer wg.Done()
			userSession, err := uc.Repository.GetUserSession(ctx, friendID)

			// If user is offline, to skip.
			if err != nil || userSession.UserID != friendID {
				<- guard
				return
			}

			// Notify all user's devices.
			for _, server := range userSession.Servers {
				u, err := url.Parse(server)
				if err != nil {
					continue
				}
				
				chatServerURL := fmt.Sprintf(
					"%v://%v/%v",
					u.Scheme,
					u.Host,
					cfg.ChatServerAPI.PresencePath,
				)
				payload := domain.UserPresenceEvent{
					TargetID: friendID,
					ClientID: arg.UserID,
					Status: arg.Status,
				}
				body, _ := json.Marshal(payload)
				_, err = http.Post(chatServerURL, "application/json", bytes.NewBuffer(body))
				if err != nil {
					logger.Error(
						"unable to broadcast user presence",
						zap.String("body", string(body)),
						zap.String("trace", err.Error()),
					)
				}
			}
			<- guard
		}(friend.UserId)
	}
	wg.Wait()
	return nil
}