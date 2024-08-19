package usecase

import (
	"context"
	pb "protobuf/proto/user"
	"session-service/internal/domain"
	"sync"
	"time"
)

func (uc * UseCaseService) BroadcastChannelEvent(ctx context.Context, arg domain.Channel) error {
	// Get contacts of all users in the channel.
	resp, err := uc.UserClient.GetUsersContactsMetadata(ctx, &pb.UserIds{UserIds: arg.UserIDs})
	if err != nil {
		return err
	}

	event := domain.BaseEvent{
		Event: domain.EventChannel,
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