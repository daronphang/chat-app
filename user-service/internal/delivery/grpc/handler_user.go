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

	friends := make([]*pb.Friend, 0)
	for _, friend := range rv.Friends {
		friends = append(friends, &pb.Friend{
			UserId: friend.UserID,
			Email: friend.Email,
			DisplayName: friend.DisplayName,
		})
	}
	return &pb.UserMetadata{
		UserId: rv.UserID,
		Email: rv.Email,
		DisplayName: rv.DisplayName,
		CreatedAt: rv.CreatedAt,
		Friends: friends,
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

func (s *GRPCServer) AddFriend(ctx context.Context, arg *pb.NewFriend) (*pb.Friend, error) {	
	p := domain.NewFriend{
		UserID: arg.UserId,
		FriendEmail: arg.FriendEmail,
		DisplayName: arg.DisplayName,
	}
	if err := cv.ValidateStruct(p); err != nil {
		logger.Error("validation error", zap.String("trace", err.Error()))
		return nil, status.Errorf(9, "validation error: %v", err)
	}
	
	rv, err := s.uc.AddFriend(ctx, p)
	if err != nil {
		return nil, status.Error(10, err.Error())
	}
	return &pb.Friend{
		UserId: rv.UserID,
		Email: rv.Email,
		DisplayName: rv.DisplayName,
	}, nil
}

func (s *GRPCServer) GetFriends(ctx context.Context, arg *wrappers.StringValue) (*pb.Friends, error) {	
	if (arg.Value == "") {
		return nil, status.Errorf(9, "user id is missing")
	}
	
	rv, err := s.uc.GetFriends(ctx, arg.Value)
	if err != nil {
		return nil, status.Error(10, err.Error())
	}

	var friends []*pb.Friend
	for _, x := range rv {
		friends = append(friends, &pb.Friend{
			UserId: x.UserID,
			Email: x.Email,
			DisplayName: x.DisplayName,
		})
	}

	return &pb.Friends{Friends: friends}, nil
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

func (s *GRPCServer) GetUsersContactsMetadata(ctx context.Context, arg *pb.Users) (*pb.UserContacts, error) {
	rv, err := s.uc.GetUsersContactsMetadata(ctx, arg.UserIds)
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