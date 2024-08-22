package grpc

import (
	"context"
	"fmt"
	"protobuf/proto/common"
	pb "protobuf/proto/session"
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
	return &common.MessageResponse{Message: "session-service is alive"}, nil
}

func (s *GRPCServer) ClientHeartbeat(ctx context.Context, arg *pb.UserSession) (*common.MessageResponse, error) {
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

func (s *GRPCServer) GetOnlineUsers(ctx context.Context, arg *pb.UserIds) (*pb.UserIds, error) {
	rv := s.uc.GetOnlineUsers(ctx, arg.UserIds) 
	onlineUsers := make([]string, 0)
	for _, x := range rv {
		onlineUsers = append(onlineUsers, x.UserID)
	}

	return &pb.UserIds{UserIds: onlineUsers}, nil
}

// When user comes online or offline.
func (s *GRPCServer) BroadcastUserPresenceEvent(ctx context.Context, arg *pb.UserPresence) (*common.MessageResponse, error) {
	p := domain.UserPresence{
		UserID: arg.UserId,
		Status: arg.Status,
		RecipientIDs: arg.RecipientIds,
	}
	if err := cv.ValidateStruct(p); err != nil {
		logger.Error("validation error", zap.String("trace", err.Error()))
		return nil, status.Errorf(9, "validation error: %v", err)
	}
	if err := s.uc.BroadcastUserPresenceEvent(ctx, p); err != nil {
		logger.Error(
			fmt.Sprintf("failed to broadcast user %v status for %v", p.UserID, p.Status),
			zap.String("trace", err.Error()),
		)
	}
	return &common.MessageResponse{Message: "user status broadcasted"}, nil
}

func (s *GRPCServer) BroadcastMessageEvent(ctx context.Context, arg *common.Message) (*common.MessageResponse, error) {
	p := domain.Message{
		MessageID: arg.MessageId,
		ChannelID: arg.ChannelId,
		SenderID: arg.SenderId,
		MessageType: arg.MessageType,
		Content: arg.Content,
		MessageStatus: int(arg.MessageStatus),
		CreatedAt: arg.CreatedAt,
	}
	if err := cv.ValidateStruct(p); err != nil {
		logger.Error("validation error", zap.String("trace", err.Error()))
		return nil, status.Errorf(9, "validation error: %v", err)
	}

	if err := s.uc.BroadcastMessageEvent(ctx, p); err != nil {
		logger.Error(
			fmt.Sprintf("failed to broadcast message event for user %v and message %v", p.SenderID, p.MessageID),
			zap.String("trace", err.Error()),
		)
		return nil, status.Errorf(9, "validation error: %v", err)
	}
	return &common.MessageResponse{Message: "user status broadcasted"}, nil
}

func (s *GRPCServer) BroadcastChannelEvent(ctx context.Context, arg *common.Channel) (*common.MessageResponse, error) {
	p := domain.Channel{
		ChannelID: arg.ChannelId,
		ChannelName: arg.ChannelName,
		CreatedAt: arg.CreatedAt,
		UserIDs: arg.UserIds,
	}
	if err := cv.ValidateStruct(p); err != nil {
		logger.Error("validation error", zap.String("trace", err.Error()))
		return nil, status.Errorf(9, "validation error: %v", err)
	}

	if err := s.uc.BroadcastChannelEvent(ctx, p); err != nil {
		logger.Error(
			fmt.Sprintf("failed to broadcast channel event for %v", p.ChannelID),
			zap.String("trace", err.Error()),
		)
		return nil, status.Errorf(9, "validation error: %v", err)
	}
	return &common.MessageResponse{Message: "user status broadcasted"}, nil
}


