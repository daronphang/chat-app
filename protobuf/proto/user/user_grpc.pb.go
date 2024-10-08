// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v5.27.1
// source: proto/user/user.proto

package user

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
	wrapperspb "google.golang.org/protobuf/types/known/wrapperspb"
	common "protobuf/proto/common"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// UserClient is the client API for User service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type UserClient interface {
	Heartbeat(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*common.MessageResponse, error)
	GetBestServer(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*common.MessageResponse, error)
	Signup(ctx context.Context, in *NewUser, opts ...grpc.CallOption) (*UserMetadata, error)
	GetUser(ctx context.Context, in *wrapperspb.StringValue, opts ...grpc.CallOption) (*UserMetadata, error)
	GetUsers(ctx context.Context, in *UserIds, opts ...grpc.CallOption) (*Users, error)
	UpdateUser(ctx context.Context, in *UserMetadata, opts ...grpc.CallOption) (*common.MessageResponse, error)
	Login(ctx context.Context, in *UserCredentials, opts ...grpc.CallOption) (*UserMetadata, error)
	Logout(ctx context.Context, in *wrapperspb.StringValue, opts ...grpc.CallOption) (*common.MessageResponse, error)
	AddFriend(ctx context.Context, in *NewFriend, opts ...grpc.CallOption) (*Friend, error)
	GetFriends(ctx context.Context, in *wrapperspb.StringValue, opts ...grpc.CallOption) (*Friends, error)
	CreateChannel(ctx context.Context, in *NewChannel, opts ...grpc.CallOption) (*common.Channel, error)
	GetUsersAssociatedToChannel(ctx context.Context, in *wrapperspb.StringValue, opts ...grpc.CallOption) (*UserContacts, error)
	GetChannelsAssociatedToUser(ctx context.Context, in *wrapperspb.StringValue, opts ...grpc.CallOption) (*Channels, error)
	GetUsersAssociatedToTargetUser(ctx context.Context, in *wrapperspb.StringValue, opts ...grpc.CallOption) (*UserIds, error)
	GetUsersContactsMetadata(ctx context.Context, in *UserIds, opts ...grpc.CallOption) (*UserContacts, error)
	AddGroupMembers(ctx context.Context, in *GroupMembers, opts ...grpc.CallOption) (*common.MessageResponse, error)
	RemoveGroupMembers(ctx context.Context, in *GroupMembers, opts ...grpc.CallOption) (*common.MessageResponse, error)
	LeaveGroup(ctx context.Context, in *GroupMembers, opts ...grpc.CallOption) (*common.MessageResponse, error)
	RemoveGroup(ctx context.Context, in *AdminGroupMember, opts ...grpc.CallOption) (*common.MessageResponse, error)
	UpdateLastReadMessage(ctx context.Context, in *LastReadMessage, opts ...grpc.CallOption) (*common.MessageResponse, error)
}

type userClient struct {
	cc grpc.ClientConnInterface
}

func NewUserClient(cc grpc.ClientConnInterface) UserClient {
	return &userClient{cc}
}

func (c *userClient) Heartbeat(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*common.MessageResponse, error) {
	out := new(common.MessageResponse)
	err := c.cc.Invoke(ctx, "/user.User/heartbeat", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userClient) GetBestServer(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*common.MessageResponse, error) {
	out := new(common.MessageResponse)
	err := c.cc.Invoke(ctx, "/user.User/getBestServer", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userClient) Signup(ctx context.Context, in *NewUser, opts ...grpc.CallOption) (*UserMetadata, error) {
	out := new(UserMetadata)
	err := c.cc.Invoke(ctx, "/user.User/signup", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userClient) GetUser(ctx context.Context, in *wrapperspb.StringValue, opts ...grpc.CallOption) (*UserMetadata, error) {
	out := new(UserMetadata)
	err := c.cc.Invoke(ctx, "/user.User/getUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userClient) GetUsers(ctx context.Context, in *UserIds, opts ...grpc.CallOption) (*Users, error) {
	out := new(Users)
	err := c.cc.Invoke(ctx, "/user.User/getUsers", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userClient) UpdateUser(ctx context.Context, in *UserMetadata, opts ...grpc.CallOption) (*common.MessageResponse, error) {
	out := new(common.MessageResponse)
	err := c.cc.Invoke(ctx, "/user.User/updateUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userClient) Login(ctx context.Context, in *UserCredentials, opts ...grpc.CallOption) (*UserMetadata, error) {
	out := new(UserMetadata)
	err := c.cc.Invoke(ctx, "/user.User/login", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userClient) Logout(ctx context.Context, in *wrapperspb.StringValue, opts ...grpc.CallOption) (*common.MessageResponse, error) {
	out := new(common.MessageResponse)
	err := c.cc.Invoke(ctx, "/user.User/logout", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userClient) AddFriend(ctx context.Context, in *NewFriend, opts ...grpc.CallOption) (*Friend, error) {
	out := new(Friend)
	err := c.cc.Invoke(ctx, "/user.User/addFriend", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userClient) GetFriends(ctx context.Context, in *wrapperspb.StringValue, opts ...grpc.CallOption) (*Friends, error) {
	out := new(Friends)
	err := c.cc.Invoke(ctx, "/user.User/getFriends", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userClient) CreateChannel(ctx context.Context, in *NewChannel, opts ...grpc.CallOption) (*common.Channel, error) {
	out := new(common.Channel)
	err := c.cc.Invoke(ctx, "/user.User/createChannel", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userClient) GetUsersAssociatedToChannel(ctx context.Context, in *wrapperspb.StringValue, opts ...grpc.CallOption) (*UserContacts, error) {
	out := new(UserContacts)
	err := c.cc.Invoke(ctx, "/user.User/getUsersAssociatedToChannel", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userClient) GetChannelsAssociatedToUser(ctx context.Context, in *wrapperspb.StringValue, opts ...grpc.CallOption) (*Channels, error) {
	out := new(Channels)
	err := c.cc.Invoke(ctx, "/user.User/getChannelsAssociatedToUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userClient) GetUsersAssociatedToTargetUser(ctx context.Context, in *wrapperspb.StringValue, opts ...grpc.CallOption) (*UserIds, error) {
	out := new(UserIds)
	err := c.cc.Invoke(ctx, "/user.User/getUsersAssociatedToTargetUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userClient) GetUsersContactsMetadata(ctx context.Context, in *UserIds, opts ...grpc.CallOption) (*UserContacts, error) {
	out := new(UserContacts)
	err := c.cc.Invoke(ctx, "/user.User/getUsersContactsMetadata", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userClient) AddGroupMembers(ctx context.Context, in *GroupMembers, opts ...grpc.CallOption) (*common.MessageResponse, error) {
	out := new(common.MessageResponse)
	err := c.cc.Invoke(ctx, "/user.User/addGroupMembers", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userClient) RemoveGroupMembers(ctx context.Context, in *GroupMembers, opts ...grpc.CallOption) (*common.MessageResponse, error) {
	out := new(common.MessageResponse)
	err := c.cc.Invoke(ctx, "/user.User/removeGroupMembers", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userClient) LeaveGroup(ctx context.Context, in *GroupMembers, opts ...grpc.CallOption) (*common.MessageResponse, error) {
	out := new(common.MessageResponse)
	err := c.cc.Invoke(ctx, "/user.User/leaveGroup", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userClient) RemoveGroup(ctx context.Context, in *AdminGroupMember, opts ...grpc.CallOption) (*common.MessageResponse, error) {
	out := new(common.MessageResponse)
	err := c.cc.Invoke(ctx, "/user.User/removeGroup", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userClient) UpdateLastReadMessage(ctx context.Context, in *LastReadMessage, opts ...grpc.CallOption) (*common.MessageResponse, error) {
	out := new(common.MessageResponse)
	err := c.cc.Invoke(ctx, "/user.User/updateLastReadMessage", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// UserServer is the server API for User service.
// All implementations must embed UnimplementedUserServer
// for forward compatibility
type UserServer interface {
	Heartbeat(context.Context, *emptypb.Empty) (*common.MessageResponse, error)
	GetBestServer(context.Context, *emptypb.Empty) (*common.MessageResponse, error)
	Signup(context.Context, *NewUser) (*UserMetadata, error)
	GetUser(context.Context, *wrapperspb.StringValue) (*UserMetadata, error)
	GetUsers(context.Context, *UserIds) (*Users, error)
	UpdateUser(context.Context, *UserMetadata) (*common.MessageResponse, error)
	Login(context.Context, *UserCredentials) (*UserMetadata, error)
	Logout(context.Context, *wrapperspb.StringValue) (*common.MessageResponse, error)
	AddFriend(context.Context, *NewFriend) (*Friend, error)
	GetFriends(context.Context, *wrapperspb.StringValue) (*Friends, error)
	CreateChannel(context.Context, *NewChannel) (*common.Channel, error)
	GetUsersAssociatedToChannel(context.Context, *wrapperspb.StringValue) (*UserContacts, error)
	GetChannelsAssociatedToUser(context.Context, *wrapperspb.StringValue) (*Channels, error)
	GetUsersAssociatedToTargetUser(context.Context, *wrapperspb.StringValue) (*UserIds, error)
	GetUsersContactsMetadata(context.Context, *UserIds) (*UserContacts, error)
	AddGroupMembers(context.Context, *GroupMembers) (*common.MessageResponse, error)
	RemoveGroupMembers(context.Context, *GroupMembers) (*common.MessageResponse, error)
	LeaveGroup(context.Context, *GroupMembers) (*common.MessageResponse, error)
	RemoveGroup(context.Context, *AdminGroupMember) (*common.MessageResponse, error)
	UpdateLastReadMessage(context.Context, *LastReadMessage) (*common.MessageResponse, error)
	mustEmbedUnimplementedUserServer()
}

// UnimplementedUserServer must be embedded to have forward compatible implementations.
type UnimplementedUserServer struct {
}

func (UnimplementedUserServer) Heartbeat(context.Context, *emptypb.Empty) (*common.MessageResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Heartbeat not implemented")
}
func (UnimplementedUserServer) GetBestServer(context.Context, *emptypb.Empty) (*common.MessageResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetBestServer not implemented")
}
func (UnimplementedUserServer) Signup(context.Context, *NewUser) (*UserMetadata, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Signup not implemented")
}
func (UnimplementedUserServer) GetUser(context.Context, *wrapperspb.StringValue) (*UserMetadata, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUser not implemented")
}
func (UnimplementedUserServer) GetUsers(context.Context, *UserIds) (*Users, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUsers not implemented")
}
func (UnimplementedUserServer) UpdateUser(context.Context, *UserMetadata) (*common.MessageResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateUser not implemented")
}
func (UnimplementedUserServer) Login(context.Context, *UserCredentials) (*UserMetadata, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Login not implemented")
}
func (UnimplementedUserServer) Logout(context.Context, *wrapperspb.StringValue) (*common.MessageResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Logout not implemented")
}
func (UnimplementedUserServer) AddFriend(context.Context, *NewFriend) (*Friend, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddFriend not implemented")
}
func (UnimplementedUserServer) GetFriends(context.Context, *wrapperspb.StringValue) (*Friends, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetFriends not implemented")
}
func (UnimplementedUserServer) CreateChannel(context.Context, *NewChannel) (*common.Channel, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateChannel not implemented")
}
func (UnimplementedUserServer) GetUsersAssociatedToChannel(context.Context, *wrapperspb.StringValue) (*UserContacts, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUsersAssociatedToChannel not implemented")
}
func (UnimplementedUserServer) GetChannelsAssociatedToUser(context.Context, *wrapperspb.StringValue) (*Channels, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetChannelsAssociatedToUser not implemented")
}
func (UnimplementedUserServer) GetUsersAssociatedToTargetUser(context.Context, *wrapperspb.StringValue) (*UserIds, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUsersAssociatedToTargetUser not implemented")
}
func (UnimplementedUserServer) GetUsersContactsMetadata(context.Context, *UserIds) (*UserContacts, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUsersContactsMetadata not implemented")
}
func (UnimplementedUserServer) AddGroupMembers(context.Context, *GroupMembers) (*common.MessageResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddGroupMembers not implemented")
}
func (UnimplementedUserServer) RemoveGroupMembers(context.Context, *GroupMembers) (*common.MessageResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RemoveGroupMembers not implemented")
}
func (UnimplementedUserServer) LeaveGroup(context.Context, *GroupMembers) (*common.MessageResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method LeaveGroup not implemented")
}
func (UnimplementedUserServer) RemoveGroup(context.Context, *AdminGroupMember) (*common.MessageResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RemoveGroup not implemented")
}
func (UnimplementedUserServer) UpdateLastReadMessage(context.Context, *LastReadMessage) (*common.MessageResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateLastReadMessage not implemented")
}
func (UnimplementedUserServer) mustEmbedUnimplementedUserServer() {}

// UnsafeUserServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to UserServer will
// result in compilation errors.
type UnsafeUserServer interface {
	mustEmbedUnimplementedUserServer()
}

func RegisterUserServer(s grpc.ServiceRegistrar, srv UserServer) {
	s.RegisterService(&User_ServiceDesc, srv)
}

func _User_Heartbeat_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(emptypb.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServer).Heartbeat(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/user.User/heartbeat",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServer).Heartbeat(ctx, req.(*emptypb.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _User_GetBestServer_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(emptypb.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServer).GetBestServer(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/user.User/getBestServer",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServer).GetBestServer(ctx, req.(*emptypb.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _User_Signup_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(NewUser)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServer).Signup(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/user.User/signup",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServer).Signup(ctx, req.(*NewUser))
	}
	return interceptor(ctx, in, info, handler)
}

func _User_GetUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(wrapperspb.StringValue)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServer).GetUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/user.User/getUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServer).GetUser(ctx, req.(*wrapperspb.StringValue))
	}
	return interceptor(ctx, in, info, handler)
}

func _User_GetUsers_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserIds)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServer).GetUsers(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/user.User/getUsers",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServer).GetUsers(ctx, req.(*UserIds))
	}
	return interceptor(ctx, in, info, handler)
}

func _User_UpdateUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserMetadata)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServer).UpdateUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/user.User/updateUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServer).UpdateUser(ctx, req.(*UserMetadata))
	}
	return interceptor(ctx, in, info, handler)
}

func _User_Login_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserCredentials)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServer).Login(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/user.User/login",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServer).Login(ctx, req.(*UserCredentials))
	}
	return interceptor(ctx, in, info, handler)
}

func _User_Logout_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(wrapperspb.StringValue)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServer).Logout(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/user.User/logout",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServer).Logout(ctx, req.(*wrapperspb.StringValue))
	}
	return interceptor(ctx, in, info, handler)
}

func _User_AddFriend_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(NewFriend)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServer).AddFriend(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/user.User/addFriend",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServer).AddFriend(ctx, req.(*NewFriend))
	}
	return interceptor(ctx, in, info, handler)
}

func _User_GetFriends_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(wrapperspb.StringValue)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServer).GetFriends(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/user.User/getFriends",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServer).GetFriends(ctx, req.(*wrapperspb.StringValue))
	}
	return interceptor(ctx, in, info, handler)
}

func _User_CreateChannel_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(NewChannel)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServer).CreateChannel(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/user.User/createChannel",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServer).CreateChannel(ctx, req.(*NewChannel))
	}
	return interceptor(ctx, in, info, handler)
}

func _User_GetUsersAssociatedToChannel_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(wrapperspb.StringValue)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServer).GetUsersAssociatedToChannel(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/user.User/getUsersAssociatedToChannel",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServer).GetUsersAssociatedToChannel(ctx, req.(*wrapperspb.StringValue))
	}
	return interceptor(ctx, in, info, handler)
}

func _User_GetChannelsAssociatedToUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(wrapperspb.StringValue)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServer).GetChannelsAssociatedToUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/user.User/getChannelsAssociatedToUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServer).GetChannelsAssociatedToUser(ctx, req.(*wrapperspb.StringValue))
	}
	return interceptor(ctx, in, info, handler)
}

func _User_GetUsersAssociatedToTargetUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(wrapperspb.StringValue)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServer).GetUsersAssociatedToTargetUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/user.User/getUsersAssociatedToTargetUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServer).GetUsersAssociatedToTargetUser(ctx, req.(*wrapperspb.StringValue))
	}
	return interceptor(ctx, in, info, handler)
}

func _User_GetUsersContactsMetadata_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserIds)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServer).GetUsersContactsMetadata(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/user.User/getUsersContactsMetadata",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServer).GetUsersContactsMetadata(ctx, req.(*UserIds))
	}
	return interceptor(ctx, in, info, handler)
}

func _User_AddGroupMembers_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GroupMembers)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServer).AddGroupMembers(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/user.User/addGroupMembers",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServer).AddGroupMembers(ctx, req.(*GroupMembers))
	}
	return interceptor(ctx, in, info, handler)
}

func _User_RemoveGroupMembers_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GroupMembers)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServer).RemoveGroupMembers(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/user.User/removeGroupMembers",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServer).RemoveGroupMembers(ctx, req.(*GroupMembers))
	}
	return interceptor(ctx, in, info, handler)
}

func _User_LeaveGroup_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GroupMembers)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServer).LeaveGroup(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/user.User/leaveGroup",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServer).LeaveGroup(ctx, req.(*GroupMembers))
	}
	return interceptor(ctx, in, info, handler)
}

func _User_RemoveGroup_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AdminGroupMember)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServer).RemoveGroup(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/user.User/removeGroup",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServer).RemoveGroup(ctx, req.(*AdminGroupMember))
	}
	return interceptor(ctx, in, info, handler)
}

func _User_UpdateLastReadMessage_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LastReadMessage)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServer).UpdateLastReadMessage(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/user.User/updateLastReadMessage",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServer).UpdateLastReadMessage(ctx, req.(*LastReadMessage))
	}
	return interceptor(ctx, in, info, handler)
}

// User_ServiceDesc is the grpc.ServiceDesc for User service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var User_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "user.User",
	HandlerType: (*UserServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "heartbeat",
			Handler:    _User_Heartbeat_Handler,
		},
		{
			MethodName: "getBestServer",
			Handler:    _User_GetBestServer_Handler,
		},
		{
			MethodName: "signup",
			Handler:    _User_Signup_Handler,
		},
		{
			MethodName: "getUser",
			Handler:    _User_GetUser_Handler,
		},
		{
			MethodName: "getUsers",
			Handler:    _User_GetUsers_Handler,
		},
		{
			MethodName: "updateUser",
			Handler:    _User_UpdateUser_Handler,
		},
		{
			MethodName: "login",
			Handler:    _User_Login_Handler,
		},
		{
			MethodName: "logout",
			Handler:    _User_Logout_Handler,
		},
		{
			MethodName: "addFriend",
			Handler:    _User_AddFriend_Handler,
		},
		{
			MethodName: "getFriends",
			Handler:    _User_GetFriends_Handler,
		},
		{
			MethodName: "createChannel",
			Handler:    _User_CreateChannel_Handler,
		},
		{
			MethodName: "getUsersAssociatedToChannel",
			Handler:    _User_GetUsersAssociatedToChannel_Handler,
		},
		{
			MethodName: "getChannelsAssociatedToUser",
			Handler:    _User_GetChannelsAssociatedToUser_Handler,
		},
		{
			MethodName: "getUsersAssociatedToTargetUser",
			Handler:    _User_GetUsersAssociatedToTargetUser_Handler,
		},
		{
			MethodName: "getUsersContactsMetadata",
			Handler:    _User_GetUsersContactsMetadata_Handler,
		},
		{
			MethodName: "addGroupMembers",
			Handler:    _User_AddGroupMembers_Handler,
		},
		{
			MethodName: "removeGroupMembers",
			Handler:    _User_RemoveGroupMembers_Handler,
		},
		{
			MethodName: "leaveGroup",
			Handler:    _User_LeaveGroup_Handler,
		},
		{
			MethodName: "removeGroup",
			Handler:    _User_RemoveGroup_Handler,
		},
		{
			MethodName: "updateLastReadMessage",
			Handler:    _User_UpdateLastReadMessage_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/user/user.proto",
}
