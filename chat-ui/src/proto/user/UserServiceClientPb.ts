/**
 * @fileoverview gRPC-Web generated client stub for user
 * @enhanceable
 * @public
 */

// Code generated by protoc-gen-grpc-web. DO NOT EDIT.
// versions:
// 	protoc-gen-grpc-web v1.5.0
// 	protoc              v5.27.1
// source: proto/user/user.proto


/* eslint-disable */
// @ts-nocheck


import * as grpcWeb from 'grpc-web';

import * as proto_common_common_pb from '../../proto/common/common_pb'; // proto import: "proto/common/common.proto"
import * as google_protobuf_empty_pb from 'google-protobuf/google/protobuf/empty_pb'; // proto import: "google/protobuf/empty.proto"
import * as google_protobuf_wrappers_pb from 'google-protobuf/google/protobuf/wrappers_pb'; // proto import: "google/protobuf/wrappers.proto"
import * as proto_user_user_pb from '../../proto/user/user_pb'; // proto import: "proto/user/user.proto"


export class UserClient {
  client_: grpcWeb.AbstractClientBase;
  hostname_: string;
  credentials_: null | { [index: string]: string; };
  options_: null | { [index: string]: any; };

  constructor (hostname: string,
               credentials?: null | { [index: string]: string; },
               options?: null | { [index: string]: any; }) {
    if (!options) options = {};
    if (!credentials) credentials = {};
    options['format'] = 'text';

    this.client_ = new grpcWeb.GrpcWebClientBase(options);
    this.hostname_ = hostname.replace(/\/+$/, '');
    this.credentials_ = credentials;
    this.options_ = options;
  }

  methodDescriptorheartbeat = new grpcWeb.MethodDescriptor(
    '/user.User/heartbeat',
    grpcWeb.MethodType.UNARY,
    google_protobuf_empty_pb.Empty,
    proto_common_common_pb.MessageResponse,
    (request: google_protobuf_empty_pb.Empty) => {
      return request.serializeBinary();
    },
    proto_common_common_pb.MessageResponse.deserializeBinary
  );

  heartbeat(
    request: google_protobuf_empty_pb.Empty,
    metadata?: grpcWeb.Metadata | null): Promise<proto_common_common_pb.MessageResponse>;

  heartbeat(
    request: google_protobuf_empty_pb.Empty,
    metadata: grpcWeb.Metadata | null,
    callback: (err: grpcWeb.RpcError,
               response: proto_common_common_pb.MessageResponse) => void): grpcWeb.ClientReadableStream<proto_common_common_pb.MessageResponse>;

  heartbeat(
    request: google_protobuf_empty_pb.Empty,
    metadata?: grpcWeb.Metadata | null,
    callback?: (err: grpcWeb.RpcError,
               response: proto_common_common_pb.MessageResponse) => void) {
    if (callback !== undefined) {
      return this.client_.rpcCall(
        this.hostname_ +
          '/user.User/heartbeat',
        request,
        metadata || {},
        this.methodDescriptorheartbeat,
        callback);
    }
    return this.client_.unaryCall(
    this.hostname_ +
      '/user.User/heartbeat',
    request,
    metadata || {},
    this.methodDescriptorheartbeat);
  }

  methodDescriptorgetBestServer = new grpcWeb.MethodDescriptor(
    '/user.User/getBestServer',
    grpcWeb.MethodType.UNARY,
    google_protobuf_empty_pb.Empty,
    proto_common_common_pb.MessageResponse,
    (request: google_protobuf_empty_pb.Empty) => {
      return request.serializeBinary();
    },
    proto_common_common_pb.MessageResponse.deserializeBinary
  );

  getBestServer(
    request: google_protobuf_empty_pb.Empty,
    metadata?: grpcWeb.Metadata | null): Promise<proto_common_common_pb.MessageResponse>;

  getBestServer(
    request: google_protobuf_empty_pb.Empty,
    metadata: grpcWeb.Metadata | null,
    callback: (err: grpcWeb.RpcError,
               response: proto_common_common_pb.MessageResponse) => void): grpcWeb.ClientReadableStream<proto_common_common_pb.MessageResponse>;

  getBestServer(
    request: google_protobuf_empty_pb.Empty,
    metadata?: grpcWeb.Metadata | null,
    callback?: (err: grpcWeb.RpcError,
               response: proto_common_common_pb.MessageResponse) => void) {
    if (callback !== undefined) {
      return this.client_.rpcCall(
        this.hostname_ +
          '/user.User/getBestServer',
        request,
        metadata || {},
        this.methodDescriptorgetBestServer,
        callback);
    }
    return this.client_.unaryCall(
    this.hostname_ +
      '/user.User/getBestServer',
    request,
    metadata || {},
    this.methodDescriptorgetBestServer);
  }

  methodDescriptorsignup = new grpcWeb.MethodDescriptor(
    '/user.User/signup',
    grpcWeb.MethodType.UNARY,
    proto_user_user_pb.NewUser,
    proto_user_user_pb.UserMetadata,
    (request: proto_user_user_pb.NewUser) => {
      return request.serializeBinary();
    },
    proto_user_user_pb.UserMetadata.deserializeBinary
  );

  signup(
    request: proto_user_user_pb.NewUser,
    metadata?: grpcWeb.Metadata | null): Promise<proto_user_user_pb.UserMetadata>;

  signup(
    request: proto_user_user_pb.NewUser,
    metadata: grpcWeb.Metadata | null,
    callback: (err: grpcWeb.RpcError,
               response: proto_user_user_pb.UserMetadata) => void): grpcWeb.ClientReadableStream<proto_user_user_pb.UserMetadata>;

  signup(
    request: proto_user_user_pb.NewUser,
    metadata?: grpcWeb.Metadata | null,
    callback?: (err: grpcWeb.RpcError,
               response: proto_user_user_pb.UserMetadata) => void) {
    if (callback !== undefined) {
      return this.client_.rpcCall(
        this.hostname_ +
          '/user.User/signup',
        request,
        metadata || {},
        this.methodDescriptorsignup,
        callback);
    }
    return this.client_.unaryCall(
    this.hostname_ +
      '/user.User/signup',
    request,
    metadata || {},
    this.methodDescriptorsignup);
  }

  methodDescriptorgetUser = new grpcWeb.MethodDescriptor(
    '/user.User/getUser',
    grpcWeb.MethodType.UNARY,
    google_protobuf_wrappers_pb.StringValue,
    proto_user_user_pb.UserMetadata,
    (request: google_protobuf_wrappers_pb.StringValue) => {
      return request.serializeBinary();
    },
    proto_user_user_pb.UserMetadata.deserializeBinary
  );

  getUser(
    request: google_protobuf_wrappers_pb.StringValue,
    metadata?: grpcWeb.Metadata | null): Promise<proto_user_user_pb.UserMetadata>;

  getUser(
    request: google_protobuf_wrappers_pb.StringValue,
    metadata: grpcWeb.Metadata | null,
    callback: (err: grpcWeb.RpcError,
               response: proto_user_user_pb.UserMetadata) => void): grpcWeb.ClientReadableStream<proto_user_user_pb.UserMetadata>;

  getUser(
    request: google_protobuf_wrappers_pb.StringValue,
    metadata?: grpcWeb.Metadata | null,
    callback?: (err: grpcWeb.RpcError,
               response: proto_user_user_pb.UserMetadata) => void) {
    if (callback !== undefined) {
      return this.client_.rpcCall(
        this.hostname_ +
          '/user.User/getUser',
        request,
        metadata || {},
        this.methodDescriptorgetUser,
        callback);
    }
    return this.client_.unaryCall(
    this.hostname_ +
      '/user.User/getUser',
    request,
    metadata || {},
    this.methodDescriptorgetUser);
  }

  methodDescriptorgetUsers = new grpcWeb.MethodDescriptor(
    '/user.User/getUsers',
    grpcWeb.MethodType.UNARY,
    proto_user_user_pb.UserIds,
    proto_user_user_pb.Users,
    (request: proto_user_user_pb.UserIds) => {
      return request.serializeBinary();
    },
    proto_user_user_pb.Users.deserializeBinary
  );

  getUsers(
    request: proto_user_user_pb.UserIds,
    metadata?: grpcWeb.Metadata | null): Promise<proto_user_user_pb.Users>;

  getUsers(
    request: proto_user_user_pb.UserIds,
    metadata: grpcWeb.Metadata | null,
    callback: (err: grpcWeb.RpcError,
               response: proto_user_user_pb.Users) => void): grpcWeb.ClientReadableStream<proto_user_user_pb.Users>;

  getUsers(
    request: proto_user_user_pb.UserIds,
    metadata?: grpcWeb.Metadata | null,
    callback?: (err: grpcWeb.RpcError,
               response: proto_user_user_pb.Users) => void) {
    if (callback !== undefined) {
      return this.client_.rpcCall(
        this.hostname_ +
          '/user.User/getUsers',
        request,
        metadata || {},
        this.methodDescriptorgetUsers,
        callback);
    }
    return this.client_.unaryCall(
    this.hostname_ +
      '/user.User/getUsers',
    request,
    metadata || {},
    this.methodDescriptorgetUsers);
  }

  methodDescriptorupdateUser = new grpcWeb.MethodDescriptor(
    '/user.User/updateUser',
    grpcWeb.MethodType.UNARY,
    proto_user_user_pb.UserMetadata,
    proto_common_common_pb.MessageResponse,
    (request: proto_user_user_pb.UserMetadata) => {
      return request.serializeBinary();
    },
    proto_common_common_pb.MessageResponse.deserializeBinary
  );

  updateUser(
    request: proto_user_user_pb.UserMetadata,
    metadata?: grpcWeb.Metadata | null): Promise<proto_common_common_pb.MessageResponse>;

  updateUser(
    request: proto_user_user_pb.UserMetadata,
    metadata: grpcWeb.Metadata | null,
    callback: (err: grpcWeb.RpcError,
               response: proto_common_common_pb.MessageResponse) => void): grpcWeb.ClientReadableStream<proto_common_common_pb.MessageResponse>;

  updateUser(
    request: proto_user_user_pb.UserMetadata,
    metadata?: grpcWeb.Metadata | null,
    callback?: (err: grpcWeb.RpcError,
               response: proto_common_common_pb.MessageResponse) => void) {
    if (callback !== undefined) {
      return this.client_.rpcCall(
        this.hostname_ +
          '/user.User/updateUser',
        request,
        metadata || {},
        this.methodDescriptorupdateUser,
        callback);
    }
    return this.client_.unaryCall(
    this.hostname_ +
      '/user.User/updateUser',
    request,
    metadata || {},
    this.methodDescriptorupdateUser);
  }

  methodDescriptorlogin = new grpcWeb.MethodDescriptor(
    '/user.User/login',
    grpcWeb.MethodType.UNARY,
    proto_user_user_pb.UserCredentials,
    proto_user_user_pb.UserMetadata,
    (request: proto_user_user_pb.UserCredentials) => {
      return request.serializeBinary();
    },
    proto_user_user_pb.UserMetadata.deserializeBinary
  );

  login(
    request: proto_user_user_pb.UserCredentials,
    metadata?: grpcWeb.Metadata | null): Promise<proto_user_user_pb.UserMetadata>;

  login(
    request: proto_user_user_pb.UserCredentials,
    metadata: grpcWeb.Metadata | null,
    callback: (err: grpcWeb.RpcError,
               response: proto_user_user_pb.UserMetadata) => void): grpcWeb.ClientReadableStream<proto_user_user_pb.UserMetadata>;

  login(
    request: proto_user_user_pb.UserCredentials,
    metadata?: grpcWeb.Metadata | null,
    callback?: (err: grpcWeb.RpcError,
               response: proto_user_user_pb.UserMetadata) => void) {
    if (callback !== undefined) {
      return this.client_.rpcCall(
        this.hostname_ +
          '/user.User/login',
        request,
        metadata || {},
        this.methodDescriptorlogin,
        callback);
    }
    return this.client_.unaryCall(
    this.hostname_ +
      '/user.User/login',
    request,
    metadata || {},
    this.methodDescriptorlogin);
  }

  methodDescriptorlogout = new grpcWeb.MethodDescriptor(
    '/user.User/logout',
    grpcWeb.MethodType.UNARY,
    google_protobuf_wrappers_pb.StringValue,
    proto_common_common_pb.MessageResponse,
    (request: google_protobuf_wrappers_pb.StringValue) => {
      return request.serializeBinary();
    },
    proto_common_common_pb.MessageResponse.deserializeBinary
  );

  logout(
    request: google_protobuf_wrappers_pb.StringValue,
    metadata?: grpcWeb.Metadata | null): Promise<proto_common_common_pb.MessageResponse>;

  logout(
    request: google_protobuf_wrappers_pb.StringValue,
    metadata: grpcWeb.Metadata | null,
    callback: (err: grpcWeb.RpcError,
               response: proto_common_common_pb.MessageResponse) => void): grpcWeb.ClientReadableStream<proto_common_common_pb.MessageResponse>;

  logout(
    request: google_protobuf_wrappers_pb.StringValue,
    metadata?: grpcWeb.Metadata | null,
    callback?: (err: grpcWeb.RpcError,
               response: proto_common_common_pb.MessageResponse) => void) {
    if (callback !== undefined) {
      return this.client_.rpcCall(
        this.hostname_ +
          '/user.User/logout',
        request,
        metadata || {},
        this.methodDescriptorlogout,
        callback);
    }
    return this.client_.unaryCall(
    this.hostname_ +
      '/user.User/logout',
    request,
    metadata || {},
    this.methodDescriptorlogout);
  }

  methodDescriptoraddFriend = new grpcWeb.MethodDescriptor(
    '/user.User/addFriend',
    grpcWeb.MethodType.UNARY,
    proto_user_user_pb.NewFriend,
    proto_user_user_pb.Friend,
    (request: proto_user_user_pb.NewFriend) => {
      return request.serializeBinary();
    },
    proto_user_user_pb.Friend.deserializeBinary
  );

  addFriend(
    request: proto_user_user_pb.NewFriend,
    metadata?: grpcWeb.Metadata | null): Promise<proto_user_user_pb.Friend>;

  addFriend(
    request: proto_user_user_pb.NewFriend,
    metadata: grpcWeb.Metadata | null,
    callback: (err: grpcWeb.RpcError,
               response: proto_user_user_pb.Friend) => void): grpcWeb.ClientReadableStream<proto_user_user_pb.Friend>;

  addFriend(
    request: proto_user_user_pb.NewFriend,
    metadata?: grpcWeb.Metadata | null,
    callback?: (err: grpcWeb.RpcError,
               response: proto_user_user_pb.Friend) => void) {
    if (callback !== undefined) {
      return this.client_.rpcCall(
        this.hostname_ +
          '/user.User/addFriend',
        request,
        metadata || {},
        this.methodDescriptoraddFriend,
        callback);
    }
    return this.client_.unaryCall(
    this.hostname_ +
      '/user.User/addFriend',
    request,
    metadata || {},
    this.methodDescriptoraddFriend);
  }

  methodDescriptorgetFriends = new grpcWeb.MethodDescriptor(
    '/user.User/getFriends',
    grpcWeb.MethodType.UNARY,
    google_protobuf_wrappers_pb.StringValue,
    proto_user_user_pb.Friends,
    (request: google_protobuf_wrappers_pb.StringValue) => {
      return request.serializeBinary();
    },
    proto_user_user_pb.Friends.deserializeBinary
  );

  getFriends(
    request: google_protobuf_wrappers_pb.StringValue,
    metadata?: grpcWeb.Metadata | null): Promise<proto_user_user_pb.Friends>;

  getFriends(
    request: google_protobuf_wrappers_pb.StringValue,
    metadata: grpcWeb.Metadata | null,
    callback: (err: grpcWeb.RpcError,
               response: proto_user_user_pb.Friends) => void): grpcWeb.ClientReadableStream<proto_user_user_pb.Friends>;

  getFriends(
    request: google_protobuf_wrappers_pb.StringValue,
    metadata?: grpcWeb.Metadata | null,
    callback?: (err: grpcWeb.RpcError,
               response: proto_user_user_pb.Friends) => void) {
    if (callback !== undefined) {
      return this.client_.rpcCall(
        this.hostname_ +
          '/user.User/getFriends',
        request,
        metadata || {},
        this.methodDescriptorgetFriends,
        callback);
    }
    return this.client_.unaryCall(
    this.hostname_ +
      '/user.User/getFriends',
    request,
    metadata || {},
    this.methodDescriptorgetFriends);
  }

  methodDescriptorcreateChannel = new grpcWeb.MethodDescriptor(
    '/user.User/createChannel',
    grpcWeb.MethodType.UNARY,
    proto_user_user_pb.NewChannel,
    proto_common_common_pb.Channel,
    (request: proto_user_user_pb.NewChannel) => {
      return request.serializeBinary();
    },
    proto_common_common_pb.Channel.deserializeBinary
  );

  createChannel(
    request: proto_user_user_pb.NewChannel,
    metadata?: grpcWeb.Metadata | null): Promise<proto_common_common_pb.Channel>;

  createChannel(
    request: proto_user_user_pb.NewChannel,
    metadata: grpcWeb.Metadata | null,
    callback: (err: grpcWeb.RpcError,
               response: proto_common_common_pb.Channel) => void): grpcWeb.ClientReadableStream<proto_common_common_pb.Channel>;

  createChannel(
    request: proto_user_user_pb.NewChannel,
    metadata?: grpcWeb.Metadata | null,
    callback?: (err: grpcWeb.RpcError,
               response: proto_common_common_pb.Channel) => void) {
    if (callback !== undefined) {
      return this.client_.rpcCall(
        this.hostname_ +
          '/user.User/createChannel',
        request,
        metadata || {},
        this.methodDescriptorcreateChannel,
        callback);
    }
    return this.client_.unaryCall(
    this.hostname_ +
      '/user.User/createChannel',
    request,
    metadata || {},
    this.methodDescriptorcreateChannel);
  }

  methodDescriptorgetUsersAssociatedToChannel = new grpcWeb.MethodDescriptor(
    '/user.User/getUsersAssociatedToChannel',
    grpcWeb.MethodType.UNARY,
    google_protobuf_wrappers_pb.StringValue,
    proto_user_user_pb.UserContacts,
    (request: google_protobuf_wrappers_pb.StringValue) => {
      return request.serializeBinary();
    },
    proto_user_user_pb.UserContacts.deserializeBinary
  );

  getUsersAssociatedToChannel(
    request: google_protobuf_wrappers_pb.StringValue,
    metadata?: grpcWeb.Metadata | null): Promise<proto_user_user_pb.UserContacts>;

  getUsersAssociatedToChannel(
    request: google_protobuf_wrappers_pb.StringValue,
    metadata: grpcWeb.Metadata | null,
    callback: (err: grpcWeb.RpcError,
               response: proto_user_user_pb.UserContacts) => void): grpcWeb.ClientReadableStream<proto_user_user_pb.UserContacts>;

  getUsersAssociatedToChannel(
    request: google_protobuf_wrappers_pb.StringValue,
    metadata?: grpcWeb.Metadata | null,
    callback?: (err: grpcWeb.RpcError,
               response: proto_user_user_pb.UserContacts) => void) {
    if (callback !== undefined) {
      return this.client_.rpcCall(
        this.hostname_ +
          '/user.User/getUsersAssociatedToChannel',
        request,
        metadata || {},
        this.methodDescriptorgetUsersAssociatedToChannel,
        callback);
    }
    return this.client_.unaryCall(
    this.hostname_ +
      '/user.User/getUsersAssociatedToChannel',
    request,
    metadata || {},
    this.methodDescriptorgetUsersAssociatedToChannel);
  }

  methodDescriptorgetChannelsAssociatedToUser = new grpcWeb.MethodDescriptor(
    '/user.User/getChannelsAssociatedToUser',
    grpcWeb.MethodType.UNARY,
    google_protobuf_wrappers_pb.StringValue,
    proto_user_user_pb.Channels,
    (request: google_protobuf_wrappers_pb.StringValue) => {
      return request.serializeBinary();
    },
    proto_user_user_pb.Channels.deserializeBinary
  );

  getChannelsAssociatedToUser(
    request: google_protobuf_wrappers_pb.StringValue,
    metadata?: grpcWeb.Metadata | null): Promise<proto_user_user_pb.Channels>;

  getChannelsAssociatedToUser(
    request: google_protobuf_wrappers_pb.StringValue,
    metadata: grpcWeb.Metadata | null,
    callback: (err: grpcWeb.RpcError,
               response: proto_user_user_pb.Channels) => void): grpcWeb.ClientReadableStream<proto_user_user_pb.Channels>;

  getChannelsAssociatedToUser(
    request: google_protobuf_wrappers_pb.StringValue,
    metadata?: grpcWeb.Metadata | null,
    callback?: (err: grpcWeb.RpcError,
               response: proto_user_user_pb.Channels) => void) {
    if (callback !== undefined) {
      return this.client_.rpcCall(
        this.hostname_ +
          '/user.User/getChannelsAssociatedToUser',
        request,
        metadata || {},
        this.methodDescriptorgetChannelsAssociatedToUser,
        callback);
    }
    return this.client_.unaryCall(
    this.hostname_ +
      '/user.User/getChannelsAssociatedToUser',
    request,
    metadata || {},
    this.methodDescriptorgetChannelsAssociatedToUser);
  }

  methodDescriptorgetUsersAssociatedToTargetUser = new grpcWeb.MethodDescriptor(
    '/user.User/getUsersAssociatedToTargetUser',
    grpcWeb.MethodType.UNARY,
    google_protobuf_wrappers_pb.StringValue,
    proto_user_user_pb.UserIds,
    (request: google_protobuf_wrappers_pb.StringValue) => {
      return request.serializeBinary();
    },
    proto_user_user_pb.UserIds.deserializeBinary
  );

  getUsersAssociatedToTargetUser(
    request: google_protobuf_wrappers_pb.StringValue,
    metadata?: grpcWeb.Metadata | null): Promise<proto_user_user_pb.UserIds>;

  getUsersAssociatedToTargetUser(
    request: google_protobuf_wrappers_pb.StringValue,
    metadata: grpcWeb.Metadata | null,
    callback: (err: grpcWeb.RpcError,
               response: proto_user_user_pb.UserIds) => void): grpcWeb.ClientReadableStream<proto_user_user_pb.UserIds>;

  getUsersAssociatedToTargetUser(
    request: google_protobuf_wrappers_pb.StringValue,
    metadata?: grpcWeb.Metadata | null,
    callback?: (err: grpcWeb.RpcError,
               response: proto_user_user_pb.UserIds) => void) {
    if (callback !== undefined) {
      return this.client_.rpcCall(
        this.hostname_ +
          '/user.User/getUsersAssociatedToTargetUser',
        request,
        metadata || {},
        this.methodDescriptorgetUsersAssociatedToTargetUser,
        callback);
    }
    return this.client_.unaryCall(
    this.hostname_ +
      '/user.User/getUsersAssociatedToTargetUser',
    request,
    metadata || {},
    this.methodDescriptorgetUsersAssociatedToTargetUser);
  }

  methodDescriptorgetUsersContactsMetadata = new grpcWeb.MethodDescriptor(
    '/user.User/getUsersContactsMetadata',
    grpcWeb.MethodType.UNARY,
    proto_user_user_pb.UserIds,
    proto_user_user_pb.UserContacts,
    (request: proto_user_user_pb.UserIds) => {
      return request.serializeBinary();
    },
    proto_user_user_pb.UserContacts.deserializeBinary
  );

  getUsersContactsMetadata(
    request: proto_user_user_pb.UserIds,
    metadata?: grpcWeb.Metadata | null): Promise<proto_user_user_pb.UserContacts>;

  getUsersContactsMetadata(
    request: proto_user_user_pb.UserIds,
    metadata: grpcWeb.Metadata | null,
    callback: (err: grpcWeb.RpcError,
               response: proto_user_user_pb.UserContacts) => void): grpcWeb.ClientReadableStream<proto_user_user_pb.UserContacts>;

  getUsersContactsMetadata(
    request: proto_user_user_pb.UserIds,
    metadata?: grpcWeb.Metadata | null,
    callback?: (err: grpcWeb.RpcError,
               response: proto_user_user_pb.UserContacts) => void) {
    if (callback !== undefined) {
      return this.client_.rpcCall(
        this.hostname_ +
          '/user.User/getUsersContactsMetadata',
        request,
        metadata || {},
        this.methodDescriptorgetUsersContactsMetadata,
        callback);
    }
    return this.client_.unaryCall(
    this.hostname_ +
      '/user.User/getUsersContactsMetadata',
    request,
    metadata || {},
    this.methodDescriptorgetUsersContactsMetadata);
  }

  methodDescriptoraddGroupMembers = new grpcWeb.MethodDescriptor(
    '/user.User/addGroupMembers',
    grpcWeb.MethodType.UNARY,
    proto_user_user_pb.GroupMembers,
    proto_common_common_pb.MessageResponse,
    (request: proto_user_user_pb.GroupMembers) => {
      return request.serializeBinary();
    },
    proto_common_common_pb.MessageResponse.deserializeBinary
  );

  addGroupMembers(
    request: proto_user_user_pb.GroupMembers,
    metadata?: grpcWeb.Metadata | null): Promise<proto_common_common_pb.MessageResponse>;

  addGroupMembers(
    request: proto_user_user_pb.GroupMembers,
    metadata: grpcWeb.Metadata | null,
    callback: (err: grpcWeb.RpcError,
               response: proto_common_common_pb.MessageResponse) => void): grpcWeb.ClientReadableStream<proto_common_common_pb.MessageResponse>;

  addGroupMembers(
    request: proto_user_user_pb.GroupMembers,
    metadata?: grpcWeb.Metadata | null,
    callback?: (err: grpcWeb.RpcError,
               response: proto_common_common_pb.MessageResponse) => void) {
    if (callback !== undefined) {
      return this.client_.rpcCall(
        this.hostname_ +
          '/user.User/addGroupMembers',
        request,
        metadata || {},
        this.methodDescriptoraddGroupMembers,
        callback);
    }
    return this.client_.unaryCall(
    this.hostname_ +
      '/user.User/addGroupMembers',
    request,
    metadata || {},
    this.methodDescriptoraddGroupMembers);
  }

  methodDescriptorremoveGroupMembers = new grpcWeb.MethodDescriptor(
    '/user.User/removeGroupMembers',
    grpcWeb.MethodType.UNARY,
    proto_user_user_pb.GroupMembers,
    proto_common_common_pb.MessageResponse,
    (request: proto_user_user_pb.GroupMembers) => {
      return request.serializeBinary();
    },
    proto_common_common_pb.MessageResponse.deserializeBinary
  );

  removeGroupMembers(
    request: proto_user_user_pb.GroupMembers,
    metadata?: grpcWeb.Metadata | null): Promise<proto_common_common_pb.MessageResponse>;

  removeGroupMembers(
    request: proto_user_user_pb.GroupMembers,
    metadata: grpcWeb.Metadata | null,
    callback: (err: grpcWeb.RpcError,
               response: proto_common_common_pb.MessageResponse) => void): grpcWeb.ClientReadableStream<proto_common_common_pb.MessageResponse>;

  removeGroupMembers(
    request: proto_user_user_pb.GroupMembers,
    metadata?: grpcWeb.Metadata | null,
    callback?: (err: grpcWeb.RpcError,
               response: proto_common_common_pb.MessageResponse) => void) {
    if (callback !== undefined) {
      return this.client_.rpcCall(
        this.hostname_ +
          '/user.User/removeGroupMembers',
        request,
        metadata || {},
        this.methodDescriptorremoveGroupMembers,
        callback);
    }
    return this.client_.unaryCall(
    this.hostname_ +
      '/user.User/removeGroupMembers',
    request,
    metadata || {},
    this.methodDescriptorremoveGroupMembers);
  }

  methodDescriptorleaveGroup = new grpcWeb.MethodDescriptor(
    '/user.User/leaveGroup',
    grpcWeb.MethodType.UNARY,
    proto_user_user_pb.GroupMembers,
    proto_common_common_pb.MessageResponse,
    (request: proto_user_user_pb.GroupMembers) => {
      return request.serializeBinary();
    },
    proto_common_common_pb.MessageResponse.deserializeBinary
  );

  leaveGroup(
    request: proto_user_user_pb.GroupMembers,
    metadata?: grpcWeb.Metadata | null): Promise<proto_common_common_pb.MessageResponse>;

  leaveGroup(
    request: proto_user_user_pb.GroupMembers,
    metadata: grpcWeb.Metadata | null,
    callback: (err: grpcWeb.RpcError,
               response: proto_common_common_pb.MessageResponse) => void): grpcWeb.ClientReadableStream<proto_common_common_pb.MessageResponse>;

  leaveGroup(
    request: proto_user_user_pb.GroupMembers,
    metadata?: grpcWeb.Metadata | null,
    callback?: (err: grpcWeb.RpcError,
               response: proto_common_common_pb.MessageResponse) => void) {
    if (callback !== undefined) {
      return this.client_.rpcCall(
        this.hostname_ +
          '/user.User/leaveGroup',
        request,
        metadata || {},
        this.methodDescriptorleaveGroup,
        callback);
    }
    return this.client_.unaryCall(
    this.hostname_ +
      '/user.User/leaveGroup',
    request,
    metadata || {},
    this.methodDescriptorleaveGroup);
  }

  methodDescriptorremoveGroup = new grpcWeb.MethodDescriptor(
    '/user.User/removeGroup',
    grpcWeb.MethodType.UNARY,
    proto_user_user_pb.AdminGroupMember,
    proto_common_common_pb.MessageResponse,
    (request: proto_user_user_pb.AdminGroupMember) => {
      return request.serializeBinary();
    },
    proto_common_common_pb.MessageResponse.deserializeBinary
  );

  removeGroup(
    request: proto_user_user_pb.AdminGroupMember,
    metadata?: grpcWeb.Metadata | null): Promise<proto_common_common_pb.MessageResponse>;

  removeGroup(
    request: proto_user_user_pb.AdminGroupMember,
    metadata: grpcWeb.Metadata | null,
    callback: (err: grpcWeb.RpcError,
               response: proto_common_common_pb.MessageResponse) => void): grpcWeb.ClientReadableStream<proto_common_common_pb.MessageResponse>;

  removeGroup(
    request: proto_user_user_pb.AdminGroupMember,
    metadata?: grpcWeb.Metadata | null,
    callback?: (err: grpcWeb.RpcError,
               response: proto_common_common_pb.MessageResponse) => void) {
    if (callback !== undefined) {
      return this.client_.rpcCall(
        this.hostname_ +
          '/user.User/removeGroup',
        request,
        metadata || {},
        this.methodDescriptorremoveGroup,
        callback);
    }
    return this.client_.unaryCall(
    this.hostname_ +
      '/user.User/removeGroup',
    request,
    metadata || {},
    this.methodDescriptorremoveGroup);
  }

  methodDescriptorupdateLastReadMessage = new grpcWeb.MethodDescriptor(
    '/user.User/updateLastReadMessage',
    grpcWeb.MethodType.UNARY,
    proto_user_user_pb.LastReadMessage,
    proto_common_common_pb.MessageResponse,
    (request: proto_user_user_pb.LastReadMessage) => {
      return request.serializeBinary();
    },
    proto_common_common_pb.MessageResponse.deserializeBinary
  );

  updateLastReadMessage(
    request: proto_user_user_pb.LastReadMessage,
    metadata?: grpcWeb.Metadata | null): Promise<proto_common_common_pb.MessageResponse>;

  updateLastReadMessage(
    request: proto_user_user_pb.LastReadMessage,
    metadata: grpcWeb.Metadata | null,
    callback: (err: grpcWeb.RpcError,
               response: proto_common_common_pb.MessageResponse) => void): grpcWeb.ClientReadableStream<proto_common_common_pb.MessageResponse>;

  updateLastReadMessage(
    request: proto_user_user_pb.LastReadMessage,
    metadata?: grpcWeb.Metadata | null,
    callback?: (err: grpcWeb.RpcError,
               response: proto_common_common_pb.MessageResponse) => void) {
    if (callback !== undefined) {
      return this.client_.rpcCall(
        this.hostname_ +
          '/user.User/updateLastReadMessage',
        request,
        metadata || {},
        this.methodDescriptorupdateLastReadMessage,
        callback);
    }
    return this.client_.unaryCall(
    this.hostname_ +
      '/user.User/updateLastReadMessage',
    request,
    metadata || {},
    this.methodDescriptorupdateLastReadMessage);
  }

}

