import { createSlice, PayloadAction } from '@reduxjs/toolkit';

// 1 = sending, 2 = sent, 3 = delivered
export type Delivery = 1 | 2 | 3;

export interface Message {
  messageId: number;
  channelId: string;
  senderId: string;
  messageType: string;
  content: string;
  createdAt: string;
  delivery?: Delivery;
}

export interface Channel {
  channelId: string;
  channelName: string;
  createdAt: string;
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
  curChannel: Channel | null;
  unreadChannels: UnreadChannelHash;
  channelHead: Channel | null;
}

const initialState: ChatState = {
  channelHash: {},
  curChannel: null,
  unreadChannels: {},
  channelHead: null,
};

export const chatSlice = createSlice({
  name: 'chat',
  initialState,
  reducers: {
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
    setHead: (state, action: PayloadAction<Channel>) => {
      state.channelHead = action.payload;
      return state;
    },
    addNewChannel: (state, action: PayloadAction<Channel>) => {
      // Multiple state updates required.
      const key = action.payload.channelId;
      if (key in state.channelHash) {
        return state;
      }

      // Update channel hash.
      state.channelHash[key] = action.payload;

      // Update channel head.
      if (!state.channelHead) {
        state.channelHead = action.payload;
      } else {
        state.channelHead.next = action.payload;
        action.payload.prev = state.channelHead;
        state.channelHead = action.payload;
      }
      return state;
    },
    initChannels: (state, action: PayloadAction<Channel[]>) => {
      action.payload.forEach(row => {
        state.channelHash[row.channelId] = row;
      });
      return state;
    },
    addNewMessage: (state, action: PayloadAction<Message>) => {
      // Multiple state updates required.
      // Channel must first be added if it does not exist.
      const key = action.payload.channelId;
      if (!(key in state.channelHash)) {
        return state;
      }

      // update message array in channel.
      const channel = state.channelHash[key];
      channel.messages.push(action.payload);

      // Update channel head.
      if (!state.channelHead) {
        return state;
      } else if (state.channelHead.channelId === channel.channelId) {
        return state;
      }

      if (channel.next) {
        channel.next.prev = channel.prev;
      }
      if (channel.prev) {
        channel.prev.next = channel.next;
      }
      channel.next = null;
      channel.prev = state.channelHead;
      state.channelHead.next = channel;
      state.channelHead = channel;

      return state;
    },
    updateChannel: (state, action: PayloadAction<Channel>) => {
      const key = action.payload.channelId;
      state.channelHash[key] = action.payload;
      return state;
    },
  },
});

export const {
  setHead,
  addNewChannel,
  initChannels,
  updateChannel,
  addNewMessage,
  setCurChannel,
  addUnreadChannel,
  removeReadChannel,
} = chatSlice.actions;
export default chatSlice.reducer;
