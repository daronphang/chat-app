import { configureStore } from '@reduxjs/toolkit';
import chatReducer from 'features/chat/redux/chatSlice';
import userReducer from 'features/user/redux/userSlice';
import configReducer from 'core/config/configSlice';
import { listenerMiddleware } from './listenerMiddleware';

export const store = configureStore({
  reducer: {
    config: configReducer,
    chat: chatReducer,
    user: userReducer,
  },
  middleware: getDefaultMiddleware =>
    getDefaultMiddleware({
      serializableCheck: {
        ignoredActions: ['listenerMiddleware/add'],
        ignoredPaths: ['config.api'],
      },
    }).concat(listenerMiddleware.middleware),
});

// Infer the `RootState` and `AppDispatch` types from the store itself
export type RootState = ReturnType<typeof store.getState>;
// Inferred type: {posts: PostsState, comments: CommentsState, users: UsersState}
export type AppDispatch = typeof store.dispatch;
