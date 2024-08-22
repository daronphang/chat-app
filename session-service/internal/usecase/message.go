package usecase

import (
	"context"
	"session-service/internal/domain"
	"sync"
	"time"

	"google.golang.org/protobuf/types/known/wrapperspb"
)


func (uc * UseCaseService) BroadcastMessageEvent(ctx context.Context, arg domain.Message) error {
	// Get all users associated to channel.
	resp, err := uc.UserClient.GetUsersAssociatedToChannel(ctx, &wrapperspb.StringValue{Value: arg.ChannelID})
	if err != nil {
		return err
	}

	// Instead of sending messages directly to the sender's devices, to push messages 
	// as events into the recipients' queues. This ensures events will never be lost.
	event := domain.BaseEvent{
		Event: domain.EventMessage,
		EventTimestamp: time.Now().UTC().Format(time.RFC3339),
		Data: arg,
	}
	maxGoroutines := 10
	guard := make(chan bool, maxGoroutines)
	wg := sync.WaitGroup{}
	for _, x := range resp.UserContacts {
		userContact := domain.UserContact{
			UserID: x.GetUserId(),
			Email: x.GetEmail(),
		}
		guard <- true
		wg.Add(1)
		go func(u domain.UserContact) {
			defer wg.Done()
			isOffline := true
			userSession, _ := uc.Repository.GetUserSession(ctx, u.UserID)
			if userSession.UserID == u.UserID {
				isOffline = false
			}
			uc.HandleEventRoutingByUserStatus(ctx, u, event, isOffline)
			<- guard
		}(userContact)
	}
	wg.Wait()
	return nil
}