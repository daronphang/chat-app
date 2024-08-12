import { createSlice, PayloadAction } from '@reduxjs/toolkit';
import { UserMetadata, Friend, UserPresence } from './user.interface';

const initialState: UserMetadata = {
  userId: '',
  email: '',
  displayName: '',
  friends: {},
};

export const userSlice = createSlice({
  name: 'user',
  initialState,
  reducers: {
    setUser: (state, action: PayloadAction<UserMetadata>) => {
      state.userId = action.payload.userId;
      state.email = action.payload.email;
      state.displayName = action.payload.displayName;
      state.friends = action.payload.friends;
      return state;
    },
    removeUser: state => {
      state.userId = '';
      state.email = '';
      state.displayName = '';
      state.friends = {};
      return state;
    },
    addFriend: (state, action: PayloadAction<Friend>) => {
      state.friends[action.payload.userId] = action.payload;
      return state;
    },
    updateOnlineFriends: (state, action: PayloadAction<string[]>) => {
      const friends = Object.values(state.friends);
      for (let i = 0; i < friends.length; i++) {
        const friend = friends[i];
        if (action.payload.includes(friend.userId)) {
          friend.isOnline = true;
        } else {
          friend.isOnline = false;
        }
      }
      return state;
    },
    updateFriendPresence: (state, action: PayloadAction<UserPresence>) => {
      if (action.payload.userId in state.friends) {
        const friend = state.friends[action.payload.userId];
        friend.isOnline = action.payload.status === 'online' ? true : false;
      }
      return state;
    },
  },
});

export const { setUser, removeUser, addFriend, updateOnlineFriends, updateFriendPresence } = userSlice.actions;
export default userSlice.reducer;
