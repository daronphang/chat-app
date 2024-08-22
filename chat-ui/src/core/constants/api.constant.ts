import { Config } from './config.constant';
import { UserClient } from 'proto/user/UserServiceClientPb';
import { MessageClient } from 'proto/message/MessageServiceClientPb';
import { SessionClient } from 'proto/session/SessionServiceClientPb';

export class BaseApi {
  public USER_SERVICE: UserClient;
  public MESSAGE_SERVICE: MessageClient;
  public SESSION_SERVICE: SessionClient;

  constructor() {}

  public initWithConfig(config: Config) {
    this.USER_SERVICE = new UserClient(config.ENVOY_PROXY_ADDRESS);
    this.MESSAGE_SERVICE = new MessageClient(config.ENVOY_PROXY_ADDRESS);
    this.SESSION_SERVICE = new SessionClient(config.ENVOY_PROXY_ADDRESS);
  }
}

export class DevApi extends BaseApi {}

export class ProdApi extends BaseApi {}

export interface ApiError {
  NETWORK_ERROR: string;
}

export const ApiError = {
  NETWORK_ERROR: 'A network error has occurred, please try again',
};
