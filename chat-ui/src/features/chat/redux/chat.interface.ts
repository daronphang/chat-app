export interface Message {
  messageId: number;
  channelId: string;
  senderId: string;
  messageType: string;
  content: string;
  createdAt: string;
  messageStatus: number; // 0 = pending, 1 = received, 2 = delivered, 3 = read
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
}

export interface WebSocketEvent {
  event: string;
  timestamp: string;
  data: any;
}
