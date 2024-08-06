import * as jspb from 'google-protobuf'

import * as google_protobuf_empty_pb from 'google-protobuf/google/protobuf/empty_pb'; // proto import: "google/protobuf/empty.proto"
import * as proto_common_common_pb from '../../proto/common/common_pb'; // proto import: "proto/common/common.proto"


export class User extends jspb.Message {
  getUserid(): string;
  setUserid(value: string): User;

  getServer(): string;
  setServer(value: string): User;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): User.AsObject;
  static toObject(includeInstance: boolean, msg: User): User.AsObject;
  static serializeBinaryToWriter(message: User, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): User;
  static deserializeBinaryFromReader(message: User, reader: jspb.BinaryReader): User;
}

export namespace User {
  export type AsObject = {
    userid: string,
    server: string,
  }
}

export class UserStatus extends jspb.Message {
  getUserid(): string;
  setUserid(value: string): UserStatus;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): UserStatus.AsObject;
  static toObject(includeInstance: boolean, msg: UserStatus): UserStatus.AsObject;
  static serializeBinaryToWriter(message: UserStatus, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): UserStatus;
  static deserializeBinaryFromReader(message: UserStatus, reader: jspb.BinaryReader): UserStatus;
}

export namespace UserStatus {
  export type AsObject = {
    userid: string,
  }
}

export class BroadcastStatus extends jspb.Message {
  getUserid(): string;
  setUserid(value: string): BroadcastStatus;

  getStatus(): string;
  setStatus(value: string): BroadcastStatus;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): BroadcastStatus.AsObject;
  static toObject(includeInstance: boolean, msg: BroadcastStatus): BroadcastStatus.AsObject;
  static serializeBinaryToWriter(message: BroadcastStatus, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): BroadcastStatus;
  static deserializeBinaryFromReader(message: BroadcastStatus, reader: jspb.BinaryReader): BroadcastStatus;
}

export namespace BroadcastStatus {
  export type AsObject = {
    userid: string,
    status: string,
  }
}

