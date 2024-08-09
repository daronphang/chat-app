package usecase

import (
	"broadcast-service/internal/domain"
	"broadcast-service/internal/util"
	"context"
	"encoding/json"
	"time"

	"go.uber.org/zap"
)

// To be executed in a goroutine.
func (uc * UseCaseService) HandleEventRoutingByUserStatus(ctx context.Context, userID string, event domain.BaseEvent, guard chan bool, isOffline bool) {
	// Push event to user's queue, regardless of online/offline status.
	err := util.ExpBackoff(
		500 * time.Millisecond,
		2,
		3,
		func() error {
			return uc.EventBroker.PublishEventToUserQueue(ctx, userID, event)
		},
	)

	b, _ := json.Marshal(event)
	if err != nil {
		logger.Error(
			"failed to push event to user queue",
			zap.String("payload", string(b)),
			zap.String("trace", err.Error()),
		)
		<- guard
		return
	}

	// If user is offline, to send push notification via message queue.
	// Delivery is not guaranteed.
	if isOffline {
		if err := uc.MessageBroker.PublishNotificationEvent(ctx, event); err != nil {
			logger.Error(
				"failed to notify user",
				zap.String("event", string(b)),
				zap.String("trace", err.Error()),
			)
		}
	}
	<- guard
}