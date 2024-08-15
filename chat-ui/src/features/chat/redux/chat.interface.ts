export interface Message {
  messageId: number;
  channelId: string;
  senderId: string;
  messageType: string;
  content: string;
  createdAt: string;
  messageStatus: number; // 0 = pending, 1 = received, 2 = delivered
  updatedAt: string;
}

export interface Channel {
  channelId: string;
  channelName: string;
  createdAt: string;
  messages: Message[];
  userIds: string[];
  isDraft?: boolean;
  updatedAt: string;
  lastMessageId: number;
}

export interface WebSocketEvent {
  event: Event;
  eventTimestamp: string;
  data: any;
}

export const MessageStatus = {
  PENDING: 0,
  RECEIVED: 1,
  DELIVERED: 2,
  READ: 3,
};

export type Event = 'event/message' | 'event/channel' | 'event/presence';
