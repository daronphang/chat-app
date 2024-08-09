package grpc

import (
	"context"
	"fmt"
	"protobuf/proto/common"
	pb "protobuf/proto/presence"
	"session-service/internal"
	"session-service/internal/domain"
	cv "session-service/internal/validator"

	"go.uber.org/zap"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
)

var (
	logger, _ = internal.WireLogger()
)

func (s *GRPCServer) Heartbeat(_ context.Context, _ *emptypb.Empty) (*common.MessageResponse, error) {
	return &common.MessageResponse{Message: "presence-service is alive"}, nil
}

func (s *GRPCServer) ClientHeartbeat(ctx context.Context, arg *pb.User) (*common.MessageResponse, error) {
	p := domain.HeartbeatRequest{
		UserID: arg.UserId,
		Server: arg.Server,
	}
	if err := cv.ValidateStruct(p); err != nil {
		logger.Error("validation error", zap.String("trace", err.Error()))
		return nil, status.Errorf(9, "validation error: %v", err)
	}
	
	if err := s.uc.ClientHeartbeat(ctx, p); err != nil {
		return nil, status.Error(2, err.Error())
	}
	return &common.MessageResponse{Message: "heartbeat updated"}, nil
}

// When user comes online or offline.
func (s *GRPCServer) BroadcastUserPresence(ctx context.Context, arg *pb.BroadcastStatus) (*common.MessageResponse, error) {
	p := domain.UserPresence{
		UserID: arg.UserId,
		Status: arg.Status,
	}
	if err := cv.ValidateStruct(p); err != nil {
		logger.Error("validation error", zap.String("trace", err.Error()))
		return nil, status.Errorf(9, "validation error: %v", err)
	}

	if err := s.uc.BroadcastUserPresence(ctx, p); err != nil {
		logger.Error(
			fmt.Sprintf("failed to broadcast user %v status for %v", p.UserID, p.Status),
			zap.String("trace", err.Error()),
		)
		return nil, status.Errorf(9, "validation error: %v", err)
	}
	return &common.MessageResponse{Message: "user status broadcasted"}, nil
}
