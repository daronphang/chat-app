package grpc

import (
	"context"
	"protobuf/proto/common"
	pb "protobuf/proto/user"
	"user-service/internal/domain"
	cv "user-service/internal/validator"

	"github.com/golang/protobuf/ptypes/wrappers"
	"go.uber.org/zap"
	"google.golang.org/grpc/status"
)

func (s *GRPCServer) CreateChannel(ctx context.Context, arg *pb.NewChannel) (*common.Channel, error) {	
	p := domain.NewChannel{
		ChannelName: arg.ChannelName,
		UserIDs: arg.UserIds,
	}
	if err := cv.ValidateStruct(p); err != nil {
		logger.Error("validation error", zap.String("trace", err.Error()))
		return nil, status.Errorf(9, "validation error: %v", err)
	}
	
	rv, err := s.uc.CreateChannel(ctx, p)
	if err != nil {
		return nil, status.Error(10, err.Error())
	}
	return &common.Channel{
		ChannelId: rv.ChannelID,
		ChannelName: rv.ChannelName,
		UserIds: rv.UserIDs,
		CreatedAt: rv.CreatedAt,
	}, nil
}

func (s *GRPCServer) GetUsersAssociatedToChannel(ctx context.Context, arg *wrappers.StringValue) (*pb.UserContacts, error) {	
	if (arg.Value == "") {
		return nil, status.Errorf(9, "channel id is missing")
	}
	
	rv, err := s.uc.GetUsersAssociatedToChannel(ctx, arg.Value)
	if err != nil {
		return nil, status.Error(10, err.Error())
	}

	var userContacts []*pb.UserContact
	for _, x := range rv {
		userContacts = append(userContacts, &pb.UserContact{
			UserId: x.UserID,
			Email: x.Email,
		})
	}

	return &pb.UserContacts{UserContacts: userContacts}, nil
}

func (s *GRPCServer) GetChannelsAssociatedToUser(ctx context.Context, arg *wrappers.StringValue) (*pb.Channels, error) {	
	if (arg.Value == "") {
		return nil, status.Errorf(9, "user id is missing")
	}
	
	rv, err := s.uc.GetChannelsAssociatedToUser(ctx, arg.Value)
	if err != nil {
		return nil, status.Error(10, err.Error())
	}

	var channels []*common.Channel
	for _, x := range rv {
		channels = append(channels, &common.Channel{
			ChannelId: x.ChannelID,
			ChannelName: x.ChannelName,
			CreatedAt: x.CreatedAt,
			UserIds: x.UserIDs,
		})
	}

	return &pb.Channels{Channels: channels}, nil
}

func (s *GRPCServer) RemoveGroup(ctx context.Context, arg *pb.AdminGroupMember) (*common.MessageResponse, error) {
	p := domain.AdminGroupMember{
		UserID: arg.UserId,
		ChannelID: arg.ChannelId,
		ChannelName: arg.ChannelName,
	}
	if err := cv.ValidateStruct(p); err != nil {
		logger.Error("validation error", zap.String("trace", err.Error()))
		return nil, status.Errorf(9, "validation error: %v", err)
	}
	
	if err := s.uc.RemoveGroup(ctx, p); err != nil {
		return nil, status.Error(10, err.Error())
	}
	return &common.MessageResponse{Message: "remove group success"}, nil
}
