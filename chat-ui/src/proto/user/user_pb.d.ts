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

  getContactsList(): Array<Contact>;
  setContactsList(value: Array<Contact>): UserMetadata;
  clearContactsList(): UserMetadata;
  addContacts(value?: Contact, index?: number): Contact;

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
    contactsList: Array<Contact.AsObject>,
  }
}

export class NewContact extends jspb.Message {
  getUserid(): string;
  setUserid(value: string): NewContact;

  getFriendemail(): string;
  setFriendemail(value: string): NewContact;

  getDisplayname(): string;
  setDisplayname(value: string): NewContact;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): NewContact.AsObject;
  static toObject(includeInstance: boolean, msg: NewContact): NewContact.AsObject;
  static serializeBinaryToWriter(message: NewContact, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): NewContact;
  static deserializeBinaryFromReader(message: NewContact, reader: jspb.BinaryReader): NewContact;
}

export namespace NewContact {
  export type AsObject = {
    userid: string,
    friendemail: string,
    displayname: string,
  }
}

export class Contact extends jspb.Message {
  getUserid(): string;
  setUserid(value: string): Contact;

  getEmail(): string;
  setEmail(value: string): Contact;

  getDisplayname(): string;
  setDisplayname(value: string): Contact;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): Contact.AsObject;
  static toObject(includeInstance: boolean, msg: Contact): Contact.AsObject;
  static serializeBinaryToWriter(message: Contact, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): Contact;
  static deserializeBinaryFromReader(message: Contact, reader: jspb.BinaryReader): Contact;
}

export namespace Contact {
  export type AsObject = {
    userid: string,
    email: string,
    displayname: string,
  }
}

export class Contacts extends jspb.Message {
  getContactsList(): Array<Contact>;
  setContactsList(value: Array<Contact>): Contacts;
  clearContactsList(): Contacts;
  addContacts(value?: Contact, index?: number): Contact;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): Contacts.AsObject;
  static toObject(includeInstance: boolean, msg: Contacts): Contacts.AsObject;
  static serializeBinaryToWriter(message: Contacts, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): Contacts;
  static deserializeBinaryFromReader(message: Contacts, reader: jspb.BinaryReader): Contacts;
}

export namespace Contacts {
  export type AsObject = {
    contactsList: Array<Contact.AsObject>,
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

