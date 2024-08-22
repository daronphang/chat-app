package usecase

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	cfg "session-service/internal/config"
	"session-service/internal/domain"
	"sync"

	"go.uber.org/zap"
)

func (uc *UseCaseService) BroadcastUserPresenceEvent(ctx context.Context, arg domain.UserPresence) error {
	if arg.Status == "offline" {
		if err := uc.Repository.RemoveUserSession(ctx, arg.UserID); err != nil {
			logger.Error(
				fmt.Sprintf("unable to remove user session for %v", arg.UserID),
				zap.String("trace", err.Error()),
			)
		}
	}

	// Performance wise, this should not have a huge impact even if the user has many friends.
	// This is because the number of online users should be minimal.
	// Otherwise, this can be improved by running an algorithm to determine the closest friends.
	// Notify users who are online only.
	maxGoroutines := 10
	guard := make(chan bool, maxGoroutines)
	wg := sync.WaitGroup{}
	cfg, _ := cfg.ProvideConfig()
	for _, recipientID := range arg.RecipientIDs {
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
			for _, serverAddress := range userSession.Servers {
				// u, err := url.Parse(server)
				// if err != nil {
				// 	continue
				// }
				
				chatServerURL := fmt.Sprintf(
					"%v://%v:%v/%v",
					"http",
					serverAddress, // u.Host
					cfg.ChatServer.Port,
					cfg.ChatServer.PresencePath,
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
		}(recipientID)
	}
	wg.Wait()
	return nil
}