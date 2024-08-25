import { PayloadAction, createSlice } from '@reduxjs/toolkit';
import { Config, Environment } from './config.constant';
import { ApiError, BaseApi, DevApi, ProdApi } from './api.constant';

export interface ConfigState {
  config: Config;
  api: BaseApi;
  chatServerAddress: string;
  chatServerWsUrl: string;
  apiError: ApiError;
  deviceId: string;
}

declare global {
  interface Window {
    REACT_APP_ENVOY_PROXY_ADDRESS: string;
    REACT_APP_CHAT_WEBSOCKET_API: string;
  }
}

const initializeState = (): ConfigState => {
  const config: Config = {
    ENVIRONMENT: (process.env.REACT_APP_ENVIRONMENT as Environment) || 'DEVELOPMENT',
    ENVOY_PROXY_ADDRESS: window.REACT_APP_ENVOY_PROXY_ADDRESS || process.env.REACT_APP_ENVOY_PROXY_ADDRESS || '',
    CHAT_WEBSOCKET_API: window.REACT_APP_CHAT_WEBSOCKET_API || process.env.REACT_APP_CHAT_WEBSOCKET_API || '',
  };

  // API class will not undergo any mutation, does not require hydration,
  // and no updates are required to the UI. Hence, it is safe to store
  // non-serializable class in Redux.
  const api = config.ENVIRONMENT === 'PRODUCTION' ? new ProdApi() : new DevApi();
  api.initWithConfig(config);

  return {
    config,
    api,
    chatServerAddress: '',
    chatServerWsUrl: '',
    apiError: ApiError,
    deviceId: '',
  };
};

export const configSlice = createSlice({
  name: 'config',
  initialState: initializeState(),
  reducers: {
    resetConfig: state => {
      state.chatServerWsUrl = '';
      state.chatServerAddress = '';
    },
    setChatServerAddress: (state, action: PayloadAction<string>) => {
      state.chatServerAddress = action.payload;
      return state;
    },
    setChatServerWsUrl: (state, action: PayloadAction<string>) => {
      state.chatServerWsUrl = action.payload;
      return state;
    },
    setDeviceId: (state, action: PayloadAction<string>) => {
      state.deviceId = action.payload;
    },
  },
});

export const { resetConfig, setChatServerAddress, setChatServerWsUrl, setDeviceId } = configSlice.actions;
export default configSlice.reducer;
