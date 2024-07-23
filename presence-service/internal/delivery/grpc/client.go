package grpc

import (
	"presence-service/internal/config"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	pb "protobuf/message"
)

func NewClient(cfg *config.Config) (pb.MessageClient, error) {
	conn, err := grpc.NewClient(
		cfg.MessageClient.HostAddress,
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		return nil, err
	}
	
	client := pb.NewMessageClient(conn)
	return client, nil
}