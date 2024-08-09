import * as jspb from 'google-protobuf'

import * as google_protobuf_empty_pb from 'google-protobuf/google/protobuf/empty_pb'; // proto import: "google/protobuf/empty.proto"
import * as google_protobuf_wrappers_pb from 'google-protobuf/google/protobuf/wrappers_pb'; // proto import: "google/protobuf/wrappers.proto"
import * as proto_common_common_pb from '../../proto/common/common_pb'; // proto import: "proto/common/common.proto"


export class PrevMessageRequest extends jspb.Message {
  getChannelid(): string;
  setChannelid(value: string): PrevMessageRequest;

  getLastmessageid(): string;
  setLastmessageid(value: string): PrevMessageRequest;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): PrevMessageRequest.AsObject;
  static toObject(includeInstance: boolean, msg: PrevMessageRequest): PrevMessageRequest.AsObject;
  static serializeBinaryToWriter(message: PrevMessageRequest, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): PrevMessageRequest;
  static deserializeBinaryFromReader(message: PrevMessageRequest, reader: jspb.BinaryReader): PrevMessageRequest;
}

export namespace PrevMessageRequest {
  export type AsObject = {
    channelid: string,
    lastmessageid: string,
  }
}

export class Messages extends jspb.Message {
  getMessagesList(): Array<proto_common_common_pb.Message>;
  setMessagesList(value: Array<proto_common_common_pb.Message>): Messages;
  clearMessagesList(): Messages;
  addMessages(value?: proto_common_common_pb.Message, index?: number): proto_common_common_pb.Message;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): Messages.AsObject;
  static toObject(includeInstance: boolean, msg: Messages): Messages.AsObject;
  static serializeBinaryToWriter(message: Messages, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): Messages;
  static deserializeBinaryFromReader(message: Messages, reader: jspb.BinaryReader): Messages;
}

export namespace Messages {
  export type AsObject = {
    messagesList: Array<proto_common_common_pb.Message.AsObject>,
  }
}

