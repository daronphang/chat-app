import { createSlice, PayloadAction } from '@reduxjs/toolkit';

export interface Message {
  messageId: string;
  channelId: string;
  senderId: string;
  messageType: string;
  content: string;
  createdAt: string;
  delivery?: number; // 0 = sending, 1 = sent, 2 = delivered
}

export interface Channel {
  channelId: string;
  channelName: string;
  messages: Message[];
  prev: Channel | null;
  next: Channel | null;
}

interface ChannelHash {
  [key: string]: Channel;
}

interface UnreadChannelHash {
  [key: string]: boolean;
}

interface ChatState {
  channelHash: ChannelHash;
  head: Channel | null;
  curChannel: Channel | null;
  unreadChannels: UnreadChannelHash;
}

const initialState: ChatState = {
  channelHash: {},
  head: {
    channelId: 'abc123',
    channelName: 'hello world!!',
    messages: [
      {
        channelId: 'abc123',
        senderId: 'user123',
        content: 'hello how are you?',
        messageId: 'msg123',
        messageType: 'text',
        createdAt: new Date().toISOString(),
        delivery: 2,
      },
      {
        channelId: 'abc123',
        senderId: 'anotheruser123',
        content: 'I am very well',
        messageId: 'msg1234',
        messageType: 'text',
        createdAt: new Date().toISOString(),
      },
    ],
    prev: null,
    next: null,
  },
  curChannel: null,
  unreadChannels: {},
};

export const chatSlice = createSlice({
  name: 'chat',
  initialState,
  reducers: {
    setHead: (state, action: PayloadAction<Channel>) => {
      state.head = action.payload;
      return state;
    },
    setCurChannel: (state, action: PayloadAction<Channel>) => {
      state.curChannel = action.payload;
      return state;
    },
    addUnreadChannel: (state, action: PayloadAction<string>) => {
      state.unreadChannels[action.payload] = true;
      return state;
    },
    removeReadChannel: (state, action: PayloadAction<string>) => {
      delete state.unreadChannels[action.payload];
      return state;
    },
    addNewChannel: (state, action: PayloadAction<Channel>) => {
      // Multiple state updates required.
      const key = action.payload.channelId;
      if (!(key in state.channelHash)) {
        state.channelHash[key] = action.payload;

        // Update linked list head.
        if (state.head) {
          action.payload.next = state.head;
          state.head.prev = action.payload;
        }
        state.head = action.payload;
      }
      return state;
    },
    addNewMessage: (state, action: PayloadAction<Message>) => {
      // Multiple state updates required.
      // Channel must first be added if it does not exist.
      const key = action.payload.channelId;
      if (key in state.channelHash) {
        const channel = state.channelHash[key];
        channel.messages.push(action.payload);

        // Update linked list head.
        if (channel.prev) {
          channel.prev.next = channel.next;

          if (channel.next) {
            channel.next.prev = channel.prev;
          }
        }

        channel.prev = null;

        if (state.head) {
          channel.next = state.head;
          state.head.prev = channel;
        }
        state.head = channel;
      }
      return state;
    },
  },
});

export const { setHead, addNewChannel, addNewMessage, setCurChannel, addUnreadChannel, removeReadChannel } =
  chatSlice.actions;
export default chatSlice.reducer;
