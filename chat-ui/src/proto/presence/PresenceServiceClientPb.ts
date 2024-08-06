/**
 * @fileoverview gRPC-Web generated client stub for presence
 * @enhanceable
 * @public
 */

// Code generated by protoc-gen-grpc-web. DO NOT EDIT.
// versions:
// 	protoc-gen-grpc-web v1.5.0
// 	protoc              v5.27.1
// source: proto/presence/presence.proto


/* eslint-disable */
// @ts-nocheck


import * as grpcWeb from 'grpc-web';

import * as proto_common_common_pb from '../../proto/common/common_pb'; // proto import: "proto/common/common.proto"
import * as google_protobuf_empty_pb from 'google-protobuf/google/protobuf/empty_pb'; // proto import: "google/protobuf/empty.proto"
import * as proto_presence_presence_pb from '../../proto/presence/presence_pb'; // proto import: "proto/presence/presence.proto"


export class PresenceClient {
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
    '/presence.Presence/heartbeat',
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
          '/presence.Presence/heartbeat',
        request,
        metadata || {},
        this.methodDescriptorheartbeat,
        callback);
    }
    return this.client_.unaryCall(
    this.hostname_ +
      '/presence.Presence/heartbeat',
    request,
    metadata || {},
    this.methodDescriptorheartbeat);
  }

  methodDescriptorclientHeartbeat = new grpcWeb.MethodDescriptor(
    '/presence.Presence/clientHeartbeat',
    grpcWeb.MethodType.UNARY,
    proto_presence_presence_pb.User,
    proto_common_common_pb.MessageResponse,
    (request: proto_presence_presence_pb.User) => {
      return request.serializeBinary();
    },
    proto_common_common_pb.MessageResponse.deserializeBinary
  );

  clientHeartbeat(
    request: proto_presence_presence_pb.User,
    metadata?: grpcWeb.Metadata | null): Promise<proto_common_common_pb.MessageResponse>;

  clientHeartbeat(
    request: proto_presence_presence_pb.User,
    metadata: grpcWeb.Metadata | null,
    callback: (err: grpcWeb.RpcError,
               response: proto_common_common_pb.MessageResponse) => void): grpcWeb.ClientReadableStream<proto_common_common_pb.MessageResponse>;

  clientHeartbeat(
    request: proto_presence_presence_pb.User,
    metadata?: grpcWeb.Metadata | null,
    callback?: (err: grpcWeb.RpcError,
               response: proto_common_common_pb.MessageResponse) => void) {
    if (callback !== undefined) {
      return this.client_.rpcCall(
        this.hostname_ +
          '/presence.Presence/clientHeartbeat',
        request,
        metadata || {},
        this.methodDescriptorclientHeartbeat,
        callback);
    }
    return this.client_.unaryCall(
    this.hostname_ +
      '/presence.Presence/clientHeartbeat',
    request,
    metadata || {},
    this.methodDescriptorclientHeartbeat);
  }

  methodDescriptorbroadcastStatus = new grpcWeb.MethodDescriptor(
    '/presence.Presence/broadcastStatus',
    grpcWeb.MethodType.UNARY,
    proto_presence_presence_pb.BroadcastStatus,
    proto_common_common_pb.MessageResponse,
    (request: proto_presence_presence_pb.BroadcastStatus) => {
      return request.serializeBinary();
    },
    proto_common_common_pb.MessageResponse.deserializeBinary
  );

  broadcastStatus(
    request: proto_presence_presence_pb.BroadcastStatus,
    metadata?: grpcWeb.Metadata | null): Promise<proto_common_common_pb.MessageResponse>;

  broadcastStatus(
    request: proto_presence_presence_pb.BroadcastStatus,
    metadata: grpcWeb.Metadata | null,
    callback: (err: grpcWeb.RpcError,
               response: proto_common_common_pb.MessageResponse) => void): grpcWeb.ClientReadableStream<proto_common_common_pb.MessageResponse>;

  broadcastStatus(
    request: proto_presence_presence_pb.BroadcastStatus,
    metadata?: grpcWeb.Metadata | null,
    callback?: (err: grpcWeb.RpcError,
               response: proto_common_common_pb.MessageResponse) => void) {
    if (callback !== undefined) {
      return this.client_.rpcCall(
        this.hostname_ +
          '/presence.Presence/broadcastStatus',
        request,
        metadata || {},
        this.methodDescriptorbroadcastStatus,
        callback);
    }
    return this.client_.unaryCall(
    this.hostname_ +
      '/presence.Presence/broadcastStatus',
    request,
    metadata || {},
    this.methodDescriptorbroadcastStatus);
  }

}

