package usecase

import (
	"context"
	"session-service/internal"
	"session-service/internal/domain"
)

var (
	logger, _ = internal.WireLogger()
)

func (uc *UseCaseService) ClientHeartbeat(ctx context.Context, arg domain.HeartbeatRequest) error {
	if err := uc.Repository.UpdateUserSession(ctx, arg); err != nil {
		return err
	}
	return nil
}

func (uc *UseCaseService) GetOfflineUsers(ctx context.Context, arg []string) []domain.UserSession {
	offlineUsers := make([]domain.UserSession, 0)
	for _, userID := range arg {
		userSession, err := uc.Repository.GetUserSession(ctx, userID)
		if err != nil || userSession.UserID == "" {
			offlineUsers = append(offlineUsers, userSession)
		}
	}
	return offlineUsers
} 

func (uc *UseCaseService) GetOnlineUsers(ctx context.Context, arg []string) []domain.UserSession {
	onlineUsers := make([]domain.UserSession, 0)
	for _, userID := range arg {
		userSession, _ := uc.Repository.GetUserSession(ctx, userID)
		if userSession.UserID == userID {
			onlineUsers = append(onlineUsers, userSession)
		}
	}
	return onlineUsers
} 

