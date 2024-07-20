package grpc

import (
	"context"
	"protobuf/common"
	pb "protobuf/message"

	"google.golang.org/protobuf/types/known/emptypb"
)

func (s *GRPCServer) Heartbeat(_ context.Context, _ *emptypb.Empty) (*common.HeartBeat, error) {
	return &common.HeartBeat{Message: "message-service is alive"}, nil
}

func (s *GRPCServer) GetLatestMessages(_ context.Context, _ *pb.MessageQuery) (*pb.Messages, error) {
	return &pb.Messages{Messages: nil}, nil
}

func (s *GRPCServer) CreateChat() {}

func (s *GRPCServer) CreateGroupChat() {}

func (s *GRPCServer) JoinGroup() {}

func (s *GRPCServer) LeaveGroup() {}

func (s *GRPCServer) DeleteChat() {}
