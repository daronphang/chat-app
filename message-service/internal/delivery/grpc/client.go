package grpc

import (
	"message-service/internal/config"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	pb "protobuf/user"
)

func NewClient(cfg *config.Config) (pb.UserClient, error) {
	conn, err := grpc.NewClient(
		cfg.UserClient.HostAddress,
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		return nil, err
	}
	
	client := pb.NewUserClient(conn)
	return client, nil
}