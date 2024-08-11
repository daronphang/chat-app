package grpc

import (
	"message-service/internal/config"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	pb "protobuf/proto/notification"
)

func NewClient(cfg *config.Config) (pb.NotificationClient, error) {
	conn, err := grpc.NewClient(
		cfg.BroadcastClient.HostAddress,
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		return nil, err
	}
	
	client := pb.NewNotificationClient(conn)
	return client, nil
}