import * as jspb from 'google-protobuf'

import * as google_protobuf_empty_pb from 'google-protobuf/google/protobuf/empty_pb'; // proto import: "google/protobuf/empty.proto"
import * as proto_common_common_pb from '../../proto/common/common_pb'; // proto import: "proto/common/common.proto"


export class UserSession extends jspb.Message {
  getUserid(): string;
  setUserid(value: string): UserSession;

  getServer(): string;
  setServer(value: string): UserSession;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): UserSession.AsObject;
  static toObject(includeInstance: boolean, msg: UserSession): UserSession.AsObject;
  static serializeBinaryToWriter(message: UserSession, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): UserSession;
  static deserializeBinaryFromReader(message: UserSession, reader: jspb.BinaryReader): UserSession;
}

export namespace UserSession {
  export type AsObject = {
    userid: string,
    server: string,
  }
}

export class UserPresence extends jspb.Message {
  getUserid(): string;
  setUserid(value: string): UserPresence;

  getStatus(): string;
  setStatus(value: string): UserPresence;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): UserPresence.AsObject;
  static toObject(includeInstance: boolean, msg: UserPresence): UserPresence.AsObject;
  static serializeBinaryToWriter(message: UserPresence, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): UserPresence;
  static deserializeBinaryFromReader(message: UserPresence, reader: jspb.BinaryReader): UserPresence;
}

export namespace UserPresence {
  export type AsObject = {
    userid: string,
    status: string,
  }
}

