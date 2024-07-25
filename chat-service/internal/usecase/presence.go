package usecase

import (
	"chat-service/internal/domain"
	"context"
)

func (uc *UseCaseService) BroadcastPresenceStatus(ctx context.Context, arg domain.PresenceStatus) error {
	if err := uc.ServerClienter.BroadcastPresenceStatus(ctx, arg); err != nil {
		return err
	}
	return nil
}