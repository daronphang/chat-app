export interface Message {
  messageId: number;
  channelId: string;
  senderId: string;
  messageType: string;
  content: string;
  createdAt: string;
  updatedAt: string;
  messageStatus: number; // 0 = pending, 1 = received, 2 = delivered
}

export interface Channel {
  channelId: string;
  channelName: string;
  createdAt: string;
  updatedAt: string;
  messages: Message[];
  userIds: string[];
  lastMessageId: number;
}

export interface WebSocketEvent {
  event: Event;
  eventTimestamp: string;
  data: any;
}

export interface UnreadChannelHash {
  [key: string]: boolean;
}

export const MessageStatus = {
  PENDING: 0,
  RECEIVED: 1,
  DELIVERED: 2,
  READ: 3,
};

export type Event = 'event/message' | 'event/channel' | 'event/presence';
