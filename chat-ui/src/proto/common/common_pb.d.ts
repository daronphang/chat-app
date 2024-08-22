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

  getUpdatedat(): string;
  setUpdatedat(value: string): Message;

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
    updatedat: string,
  }
}

export class Channel extends jspb.Message {
  getChannelid(): string;
  setChannelid(value: string): Channel;

  getChannelname(): string;
  setChannelname(value: string): Channel;

  getCreatedat(): string;
  setCreatedat(value: string): Channel;

  getUseridsList(): Array<string>;
  setUseridsList(value: Array<string>): Channel;
  clearUseridsList(): Channel;
  addUserids(value: string, index?: number): Channel;

  getLastmessageid(): number;
  setLastmessageid(value: number): Channel;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): Channel.AsObject;
  static toObject(includeInstance: boolean, msg: Channel): Channel.AsObject;
  static serializeBinaryToWriter(message: Channel, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): Channel;
  static deserializeBinaryFromReader(message: Channel, reader: jspb.BinaryReader): Channel;
}

export namespace Channel {
  export type AsObject = {
    channelid: string,
    channelname: string,
    createdat: string,
    useridsList: Array<string>,
    lastmessageid: number,
  }
}

