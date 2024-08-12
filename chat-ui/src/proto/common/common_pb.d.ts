import * as jspb from 'google-protobuf'



export class MessageResponse extends jspb.Message {
  getMessage(): string;
  setMessage(value: string): MessageResponse;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): MessageResponse.AsObject;
  static toObject(includeInstance: boolean, msg: MessageResponse): MessageResponse.AsObject;
  static serializeBinaryToWriter(message: MessageResponse, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): MessageResponse;
  static deserializeBinaryFromReader(message: MessageResponse, reader: jspb.BinaryReader): MessageResponse;
}

export namespace MessageResponse {
  export type AsObject = {
    message: string,
  }
}

export class Message extends jspb.Message {
  getMessageid(): number;
  setMessageid(value: number): Message;

  getChannelid(): string;
  setChannelid(value: string): Message;

  getSenderid(): string;
  setSenderid(value: string): Message;

  getMessagetype(): string;
  setMessagetype(value: string): Message;

  getContent(): string;
  setContent(value: string): Message;

  getCreatedat(): string;
  setCreatedat(value: string): Message;

  getMessagestatus(): number;
  setMessagestatus(value: number): Message;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): Message.AsObject;
  static toObject(includeInstance: boolean, msg: Message): Message.AsObject;
  static serializeBinaryToWriter(message: Message, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): Message;
  static deserializeBinaryFromReader(message: Message, reader: jspb.BinaryReader): Message;
}

export namespace Message {
  export type AsObject = {
    messageid: number,
    channelid: string,
    senderid: string,
    messagetype: string,
    content: string,
    createdat: string,
    messagestatus: number,
  }
}

