package grpc

import (
	"context"
	"protobuf/proto/common"
	pb "protobuf/proto/user"
	"user-service/internal"
	"user-service/internal/domain"
	cv "user-service/internal/validator"

	"github.com/golang/protobuf/ptypes/wrappers"
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

func (s *GRPCServer) Signup(ctx context.Context, arg *pb.NewUser) (*pb.UserMetadata, error) {
	p := domain.NewUser{
		Email: arg.Email,
		DisplayName: arg.DisplayName,
	}
	if err := cv.ValidateStruct(p); err != nil {
		logger.Error("validation error", zap.String("trace", err.Error()))
		return nil, status.Errorf(9, "validation error: %v", err)
	}
	
	rv, err := s.uc.Signup(ctx, p); 
	if err != nil {
		return nil, status.Error(10, err.Error())
	}
	return &pb.UserMetadata{
		UserId: rv.UserID,
		Email: rv.Email,
		DisplayName: rv.DisplayName,
		CreatedAt: rv.CreatedAt,
	}, nil
}

func (s *GRPCServer) Login(ctx context.Context, arg *pb.UserCredentials) (*pb.UserMetadata, error) {	
	p := domain.UserCredentials{
		Email: arg.Email,
	}
	if err := cv.ValidateStruct(p); err != nil {
		logger.Error("validation error", zap.String("trace", err.Error()))
		return nil, status.Errorf(9, "validation error: %v", err)
	}
	
	rv, err := s.uc.Login(ctx, p); 
	if err != nil {
		return nil, status.Error(16, err.Error())
	}

	contacts := make([]*pb.Contact, 0)
	for _, contact := range rv.Contacts {
		contacts = append(contacts, &pb.Contact{
			UserId: contact.UserID,
			Email: contact.Email,
			DisplayName: contact.DisplayName,
		})
	}
	return &pb.UserMetadata{
		UserId: rv.UserID,
		Email: rv.Email,
		DisplayName: rv.DisplayName,
		CreatedAt: rv.CreatedAt,
		Contacts: contacts,
	}, nil
}

func (s *GRPCServer) UpdateUser(ctx context.Context, arg *pb.UserMetadata) (*common.MessageResponse, error) {	
	p := domain.UserMetadata{
		UserID: arg.UserId,
		Email: arg.Email,
		DisplayName: arg.DisplayName,
	}
	if err := cv.ValidateStruct(p); err != nil {
		logger.Error("validation error", zap.String("trace", err.Error()))
		return nil, status.Errorf(9, "validation error: %v", err)
	}
	
	if err := s.uc.UpdateUser(ctx, p); err != nil {
		return nil, status.Error(10, err.Error())
	}
	return &common.MessageResponse{Message: "user updated"}, nil
}


func (s *GRPCServer) GetBestServer(ctx context.Context,  _ *emptypb.Empty) (*common.MessageResponse, error) {
	server, err := s.uc.GetBestServer(ctx)
	if err != nil {
		return nil, status.Errorf(10, err.Error())
	}
	return &common.MessageResponse{Message: server}, nil
}

func (s *GRPCServer) CreateContact(ctx context.Context, arg *pb.NewContact) (*pb.Contact, error) {	
	p := domain.NewContact{
		UserID: arg.UserId,
		FriendEmail: arg.FriendEmail,
		DisplayName: arg.DisplayName,
	}
	if err := cv.ValidateStruct(p); err != nil {
		logger.Error("validation error", zap.String("trace", err.Error()))
		return nil, status.Errorf(9, "validation error: %v", err)
	}
	
	rv, err := s.uc.CreateContact(ctx, p)
	if err != nil {
		return nil, status.Error(10, err.Error())
	}
	return &pb.Contact{
		UserId: rv.UserID,
		Email: rv.Email,
		DisplayName: rv.DisplayName,
	}, nil
}

func (s *GRPCServer) GetContacts(ctx context.Context, arg *wrappers.StringValue) (*pb.Contacts, error) {	
	if (arg.Value == "") {
		return nil, status.Errorf(9, "user id is missing")
	}
	
	rv, err := s.uc.GetContacts(ctx, arg.Value)
	if err != nil {
		return nil, status.Error(10, err.Error())
	}

	var contacts []*pb.Contact
	for _, x := range rv {
		contacts = append(contacts, &pb.Contact{
			UserId: x.UserID,
			Email: x.Email,
			DisplayName: x.DisplayName,
		})
	}

	return &pb.Contacts{Contacts: contacts}, nil
}

func (s *GRPCServer) CreateChannel(ctx context.Context, arg *pb.NewChannel) (*pb.Channel, error) {	
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
	return &pb.Channel{
		ChannelId: rv.ChannelID,
		ChannelName: rv.ChannelName,
	}, nil
}

func (s *GRPCServer) GetUsersAssociatedToChannel(ctx context.Context, arg *wrappers.StringValue) (*pb.Users, error) {	
	if (arg.Value == "") {
		return nil, status.Errorf(9, "channel id is missing")
	}
	
	rv, err := s.uc.GetUsersAssociatedToChannel(ctx, arg.Value)
	if err != nil {
		return nil, status.Error(10, err.Error())
	}
	return &pb.Users{UserIds: rv}, nil
}

func (s *GRPCServer) GetChannelsAssociatedToUser(ctx context.Context, arg *wrappers.StringValue) (*pb.Channels, error) {	
	if (arg.Value == "") {
		return nil, status.Errorf(9, "user id is missing")
	}
	
	rv, err := s.uc.GetChannelsAssociatedToUser(ctx, arg.Value)
	if err != nil {
		return nil, status.Error(10, err.Error())
	}

	var channels []*pb.Channel
	for _, x := range rv {
		channels = append(channels, &pb.Channel{
			ChannelId: x.ChannelID,
			ChannelName: x.ChannelName,
			CreatedAt: x.CreatedAt,
		})
	}

	return &pb.Channels{Channels: channels}, nil
}

func (s *GRPCServer) GetUsersAssociatedToTargetUser(ctx context.Context, arg *wrappers.StringValue) (*pb.Users, error) {
	if (arg.Value == "") {
		return nil, status.Errorf(9, "user id is missing")
	}

	rv, err := s.uc.GetUsersAssociatedToTargetUser(ctx, arg.Value)
	if err != nil {
		return nil, status.Error(10, err.Error())
	}
	return &pb.Users{UserIds: rv}, nil
}