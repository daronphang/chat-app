import * as jspb from 'google-protobuf'

import * as google_protobuf_empty_pb from 'google-protobuf/google/protobuf/empty_pb'; // proto import: "google/protobuf/empty.proto"
import * as google_protobuf_wrappers_pb from 'google-protobuf/google/protobuf/wrappers_pb'; // proto import: "google/protobuf/wrappers.proto"
import * as proto_common_common_pb from '../../proto/common/common_pb'; // proto import: "proto/common/common.proto"


export class NewUser extends jspb.Message {
  getEmail(): string;
  setEmail(value: string): NewUser;

  getDisplayname(): string;
  setDisplayname(value: string): NewUser;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): NewUser.AsObject;
  static toObject(includeInstance: boolean, msg: NewUser): NewUser.AsObject;
  static serializeBinaryToWriter(message: NewUser, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): NewUser;
  static deserializeBinaryFromReader(message: NewUser, reader: jspb.BinaryReader): NewUser;
}

export namespace NewUser {
  export type AsObject = {
    email: string,
    displayname: string,
  }
}

export class UserCredentials extends jspb.Message {
  getEmail(): string;
  setEmail(value: string): UserCredentials;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): UserCredentials.AsObject;
  static toObject(includeInstance: boolean, msg: UserCredentials): UserCredentials.AsObject;
  static serializeBinaryToWriter(message: UserCredentials, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): UserCredentials;
  static deserializeBinaryFromReader(message: UserCredentials, reader: jspb.BinaryReader): UserCredentials;
}

export namespace UserCredentials {
  export type AsObject = {
    email: string,
  }
}

export class UserMetadata extends jspb.Message {
  getUserid(): string;
  setUserid(value: string): UserMetadata;

  getEmail(): string;
  setEmail(value: string): UserMetadata;

  getDisplayname(): string;
  setDisplayname(value: string): UserMetadata;

  getCreatedat(): string;
  setCreatedat(value: string): UserMetadata;

  getFriendsList(): Array<Friend>;
  setFriendsList(value: Array<Friend>): UserMetadata;
  clearFriendsList(): UserMetadata;
  addFriends(value?: Friend, index?: number): Friend;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): UserMetadata.AsObject;
  static toObject(includeInstance: boolean, msg: UserMetadata): UserMetadata.AsObject;
  static serializeBinaryToWriter(message: UserMetadata, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): UserMetadata;
  static deserializeBinaryFromReader(message: UserMetadata, reader: jspb.BinaryReader): UserMetadata;
}

export namespace UserMetadata {
  export type AsObject = {
    userid: string,
    email: string,
    displayname: string,
    createdat: string,
    friendsList: Array<Friend.AsObject>,
  }
}

export class UserContact extends jspb.Message {
  getUserid(): string;
  setUserid(value: string): UserContact;

  getEmail(): string;
  setEmail(value: string): UserContact;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): UserContact.AsObject;
  static toObject(includeInstance: boolean, msg: UserContact): UserContact.AsObject;
  static serializeBinaryToWriter(message: UserContact, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): UserContact;
  static deserializeBinaryFromReader(message: UserContact, reader: jspb.BinaryReader): UserContact;
}

export namespace UserContact {
  export type AsObject = {
    userid: string,
    email: string,
  }
}

export class UserContacts extends jspb.Message {
  getUsercontactsList(): Array<UserContact>;
  setUsercontactsList(value: Array<UserContact>): UserContacts;
  clearUsercontactsList(): UserContacts;
  addUsercontacts(value?: UserContact, index?: number): UserContact;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): UserContacts.AsObject;
  static toObject(includeInstance: boolean, msg: UserContacts): UserContacts.AsObject;
  static serializeBinaryToWriter(message: UserContacts, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): UserContacts;
  static deserializeBinaryFromReader(message: UserContacts, reader: jspb.BinaryReader): UserContacts;
}

export namespace UserContacts {
  export type AsObject = {
    usercontactsList: Array<UserContact.AsObject>,
  }
}

export class NewFriend extends jspb.Message {
  getUserid(): string;
  setUserid(value: string): NewFriend;

  getFriendemail(): string;
  setFriendemail(value: string): NewFriend;

  getDisplayname(): string;
  setDisplayname(value: string): NewFriend;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): NewFriend.AsObject;
  static toObject(includeInstance: boolean, msg: NewFriend): NewFriend.AsObject;
  static serializeBinaryToWriter(message: NewFriend, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): NewFriend;
  static deserializeBinaryFromReader(message: NewFriend, reader: jspb.BinaryReader): NewFriend;
}

export namespace NewFriend {
  export type AsObject = {
    userid: string,
    friendemail: string,
    displayname: string,
  }
}

export class Friend extends jspb.Message {
  getUserid(): string;
  setUserid(value: string): Friend;

  getEmail(): string;
  setEmail(value: string): Friend;

  getDisplayname(): string;
  setDisplayname(value: string): Friend;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): Friend.AsObject;
  static toObject(includeInstance: boolean, msg: Friend): Friend.AsObject;
  static serializeBinaryToWriter(message: Friend, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): Friend;
  static deserializeBinaryFromReader(message: Friend, reader: jspb.BinaryReader): Friend;
}

export namespace Friend {
  export type AsObject = {
    userid: string,
    email: string,
    displayname: string,
  }
}

export class Friends extends jspb.Message {
  getFriendsList(): Array<Friend>;
  setFriendsList(value: Array<Friend>): Friends;
  clearFriendsList(): Friends;
  addFriends(value?: Friend, index?: number): Friend;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): Friends.AsObject;
  static toObject(includeInstance: boolean, msg: Friends): Friends.AsObject;
  static serializeBinaryToWriter(message: Friends, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): Friends;
  static deserializeBinaryFromReader(message: Friends, reader: jspb.BinaryReader): Friends;
}

export namespace Friends {
  export type AsObject = {
    friendsList: Array<Friend.AsObject>,
  }
}

export class NewChannel extends jspb.Message {
  getChannelid(): string;
  setChannelid(value: string): NewChannel;

  getUseridsList(): Array<string>;
  setUseridsList(value: Array<string>): NewChannel;
  clearUseridsList(): NewChannel;
  addUserids(value: string, index?: number): NewChannel;

  getChannelname(): string;
  setChannelname(value: string): NewChannel;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): NewChannel.AsObject;
  static toObject(includeInstance: boolean, msg: NewChannel): NewChannel.AsObject;
  static serializeBinaryToWriter(message: NewChannel, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): NewChannel;
  static deserializeBinaryFromReader(message: NewChannel, reader: jspb.BinaryReader): NewChannel;
}

export namespace NewChannel {
  export type AsObject = {
    channelid: string,
    useridsList: Array<string>,
    channelname: string,
  }
}

export class Users extends jspb.Message {
  getUseridsList(): Array<string>;
  setUseridsList(value: Array<string>): Users;
  clearUseridsList(): Users;
  addUserids(value: string, index?: number): Users;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): Users.AsObject;
  static toObject(includeInstance: boolean, msg: Users): Users.AsObject;
  static serializeBinaryToWriter(message: Users, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): Users;
  static deserializeBinaryFromReader(message: Users, reader: jspb.BinaryReader): Users;
}

export namespace Users {
  export type AsObject = {
    useridsList: Array<string>,
  }
}

export class Channel extends jspb.Message {
  getChannelid(): string;
  setChannelid(value: string): Channel;

  getChannelname(): string;
  setChannelname(value: string): Channel;

  getCreatedat(): string;
  setCreatedat(value: string): Channel;

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
  }
}

export class Channels extends jspb.Message {
  getChannelsList(): Array<Channel>;
  setChannelsList(value: Array<Channel>): Channels;
  clearChannelsList(): Channels;
  addChannels(value?: Channel, index?: number): Channel;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): Channels.AsObject;
  static toObject(includeInstance: boolean, msg: Channels): Channels.AsObject;
  static serializeBinaryToWriter(message: Channels, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): Channels;
  static deserializeBinaryFromReader(message: Channels, reader: jspb.BinaryReader): Channels;
}

export namespace Channels {
  export type AsObject = {
    channelsList: Array<Channel.AsObject>,
  }
}

