import { createSlice, PayloadAction } from '@reduxjs/toolkit';

export interface UserMetadata {
  userId: string;
  email: string;
  displayName: String;
  contacts: UserMetadata[];
}

const initialState: UserMetadata = {
  userId: '',
  email: '',
  displayName: '',
  contacts: [],
};

export const userSlice = createSlice({
  name: 'user',
  initialState,
  reducers: {
    setUser: (state, action: PayloadAction<UserMetadata>) => {
      state.userId = action.payload.userId;
      state.email = action.payload.email;
      state.displayName = action.payload.displayName;
      return state;
    },
    removeUser: (state, action) => {
      state.userId = '';
      state.email = '';
      state.displayName = '';
      state.contacts = [];
      return state;
    },
    addContacts: (state, action: PayloadAction<UserMetadata[]>) => {
      state.contacts = action.payload;
      return state;
    },
  },
});

export const { setUser, removeUser, addContacts } = userSlice.actions;
export default userSlice.reducer;
