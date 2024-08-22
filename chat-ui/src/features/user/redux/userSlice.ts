import { createSlice, PayloadAction } from '@reduxjs/toolkit';
import { UserMetadata, Recipient, UserPresence } from './user.interface';

const initialState: UserMetadata = {
  userId: '',
  email: '',
  displayName: '',
  recipients: {},
};

export const userSlice = createSlice({
  name: 'user',
  initialState,
  reducers: {
    setUser: (state, action: PayloadAction<UserMetadata>) => {
      state.userId = action.payload.userId;
      state.email = action.payload.email;
      state.displayName = action.payload.displayName;
      state.recipients = action.payload.recipients;
      return state;
    },
    resetUser: state => {
      state.userId = '';
      state.email = '';
      state.displayName = '';
      state.recipients = {};
      return state;
    },
    addRecipient: (state, action: PayloadAction<Recipient>) => {
      state.recipients[action.payload.userId] = action.payload;
      return state;
    },
    addRecipients: (state, action: PayloadAction<Recipient[]>) => {
      action.payload.forEach(row => {
        state.recipients[row.userId] = row;
      });
      return state;
    },
    updateOnlineRecipients: (state, action: PayloadAction<string[]>) => {
      const recipients = Object.values(state.recipients);
      for (let i = 0; i < recipients.length; i++) {
        const recipient = recipients[i];
        if (action.payload.includes(recipient.userId)) {
          recipient.isOnline = true;
        } else {
          recipient.isOnline = false;
        }
      }
      return state;
    },
    updateFriendPresence: (state, action: PayloadAction<UserPresence>) => {
      if (action.payload.clientId in state.recipients) {
        const recipient = state.recipients[action.payload.clientId];
        recipient.isOnline = action.payload.status === 'online' ? true : false;
      }
      return state;
    },
  },
});

export const { setUser, resetUser, addRecipient, addRecipients, updateOnlineRecipients, updateFriendPresence } =
  userSlice.actions;
export default userSlice.reducer;
