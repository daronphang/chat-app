import { createSlice, PayloadAction } from '@reduxjs/toolkit';

export interface UserMetadata {
  userId: string;
  email: string;
  displayName: String;
  contacts: UserMetadata[];
}

const initialState: UserMetadata = {
  userId: 'user123',
  email: '',
  displayName: '',
  contacts: [],
};

export const userSlice = createSlice({
  name: 'user',
  initialState,
  reducers: {
    addContact: (state, action: PayloadAction<UserMetadata>) => {
      state.contacts.push(action.payload);
      return state;
    },
  },
});

export const { addContact } = userSlice.actions;
export default userSlice.reducer;
