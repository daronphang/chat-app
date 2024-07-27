package grpc

import (
	"context"
	"protobuf/common"
	pb "protobuf/user"
	"user-service/internal"
	"user-service/internal/domain"
	cv "user-service/internal/validator"

	"go.uber.org/zap"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
)

var (
	logger, _ = internal.WireLogger()
)

func (s *GRPCServer) Heartbeat(_ context.Context, _ *emptypb.Empty) (*common.MessageResponse, error) {
	return &common.MessageResponse{Message: "user-service is alive"}, nil
}

func (s *GRPCServer) CreateUser(ctx context.Context, arg *pb.NewUser) (*pb.UserMetadata, error) {
	p := domain.NewUser{
		Email: arg.Email,
		DisplayName: arg.DisplayName,
	}
	if err := cv.ValidateStruct(p); err != nil {
		logger.Error("validation error", zap.String("trace", err.Error()))
		return nil, status.Errorf(9, "validation error: %v", err)
	}
	
	rv, err := s.uc.CreateUser(ctx, p); 
	if err != nil {
		return nil, status.Error(2, err.Error())
	}
	return &pb.UserMetadata{
		UserId: rv.UserID,
		Email: rv.Email,
		DisplayName: rv.DisplayName,
		CreatedAt: rv.CreatedAt,
	}, nil
}

func (s *GRPCServer) GetBestServer(ctx context.Context,  _ *emptypb.Empty) (*common.MessageResponse, error) {
	server, err := s.uc.GetBestServer(ctx)
	if err != nil {
		return nil, status.Errorf(2, err.Error())
	}
	return &common.MessageResponse{Message: server}, nil
}
