import { Options } from 'react-use-websocket';

export const defaultWsOptions: Options = {
  shouldReconnect: () => true,
  share: true,
  reconnectAttempts: 15,
  reconnectInterval: attemptNumber => attemptNumber * 2 * 1000,
  heartbeat: true,
};
