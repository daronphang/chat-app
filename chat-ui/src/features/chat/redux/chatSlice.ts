import { createSlice, PayloadAction } from '@reduxjs/toolkit';

export interface Message {
  messageId: string;
  channelId: string;
  senderId: string;
  messageType: string;
  content: string;
  createdAt: string;
}

interface;
