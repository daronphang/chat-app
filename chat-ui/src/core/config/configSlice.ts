import { createSlice, PayloadAction } from '@reduxjs/toolkit';
import { Config, Environment } from './config.constant';
import { BaseApi, DevApi, ProdApi } from './api.constant';

interface ConfigState {
  config: Config;
  api: BaseApi;
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
  };
};

export const configSlice = createSlice({
  name: 'config',
  initialState: initializeState(),
  reducers: {},
});

export default configSlice.reducer;
