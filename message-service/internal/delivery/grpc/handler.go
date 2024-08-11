package grpc

import (
	"context"
	"message-service/internal"
	"message-service/internal/domain"
	cv "message-service/internal/validator"
	"protobuf/proto/common"
	pb "protobuf/proto/message"

	"github.com/golang/protobuf/ptypes/wrappers"
	"go.uber.org/zap"

	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
)

var (
	logger, _ = internal.WireLogger()
)

func (s *GRPCServer) Heartbeat(_ context.Context, _ *emptypb.Empty) (*common.MessageResponse, error) {
	return &common.MessageResponse{Message: "message-service is alive"}, nil
}

func (s *GRPCServer) GetLatestMessages(ctx context.Context, arg *wrappers.StringValue) (*pb.Messages, error) {
	if (arg.Value == "") {
		return nil, status.Errorf(9, "channel id is missing")
	}
	
	msgs, err := s.uc.GetLatestMessages(ctx, arg.Value)
	if err != nil {
		return nil, status.Errorf(10, "failed to fetch latest messages: %v", err)
	}
	rv := &pb.Messages{}
	for _, msg := range msgs {
		rv.Messages = append(rv.Messages, &common.Message{
			MessageId: msg.MessageID,
			ChannelId: msg.ChannelID,
			SenderId: msg.SenderID,
			MessageType: msg.MessageType,
			Content: msg.Content,
			CreatedAt: msg.CreatedAt,
			MessageStatus: string(msg.MessageStatus),
		})
	}

	return rv, nil
}

func (s *GRPCServer) GetPreviousMessages(ctx context.Context, arg *pb.PrevMessageRequest) (*pb.Messages, error) {
	p := domain.PrevMessageRequest{
		ChannelID: arg.ChannelId,
		LastMessageID: arg.LastMessageId,
	}
	if err := cv.ValidateStruct(p); err != nil {
		logger.Error("validation error", zap.String("trace", err.Error()))
		return nil, status.Errorf(9, "validation error: %v", err)
	}
	
	msgs, err := s.uc.GetPreviousMessages(ctx, p)
	if err != nil {
		return nil, status.Errorf(10, "failed to fetch latest messages: %v", err)
	}
	rv := &pb.Messages{}
	for _, msg := range msgs {
		rv.Messages = append(rv.Messages, &common.Message{
			MessageId: msg.MessageID,
			ChannelId: msg.ChannelID,
			SenderId: msg.SenderID,
			MessageType: msg.MessageType,
			Content: msg.Content,
			CreatedAt: msg.CreatedAt,
			MessageStatus: string(msg.MessageStatus),
		})
	}

	return rv, nil
}