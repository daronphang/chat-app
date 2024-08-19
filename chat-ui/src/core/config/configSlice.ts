import { PayloadAction, createSlice } from '@reduxjs/toolkit';
import { Config, Environment } from './config.constant';
import { ApiError, BaseApi, DevApi, ProdApi } from './api.constant';

export interface ConfigState {
  config: Config;
  api: BaseApi;
  chatServerWsUrl: string;
  apiError: ApiError;
  deviceId: string;
}

const initializeState = (): ConfigState => {
  const config: Config = {
    ENVIRONMENT: (process.env.REACT_APP_ENVIRONMENT as Environment) || 'DEVELOPMENT',
    ENVOY_PROXY_ADDRESS: process.env.REACT_APP_ENVOY_PROXY_ADDRESS || '',
  };

  // API class will not undergo any mutation, does not require hydration,
  // and no updates are required to the UI. Hence, it is safe to store
  // non-serializable class in Redux.
  const api = config.ENVIRONMENT === 'PRODUCTION' ? new ProdApi() : new DevApi();
  api.initWithConfig(config);

  return {
    config,
    api,
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

export const { resetConfig, setChatServerWsUrl, setDeviceId } = configSlice.actions;
export default configSlice.reducer;
