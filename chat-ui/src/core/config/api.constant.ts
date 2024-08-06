import { Config } from './config.constant';
import { UserClient } from 'proto/user/UserServiceClientPb';
import { MessageClient } from 'proto/message/MessageServiceClientPb';

export class BaseApi {
  public USER_SERVICE: UserClient;
  public MESSAGE_SERVICE: MessageClient;

  constructor() {}

  public initWithConfig(config: Config) {
    this.USER_SERVICE = new UserClient(config.ENVOY_PROXY_ADDRESS);
    this.MESSAGE_SERVICE = new MessageClient(config.ENVOY_PROXY_ADDRESS);
  }
}

export class DevApi extends BaseApi {}

export class ProdApi extends BaseApi {}

export const ApiErrors = {
  NETWORK_ERROR: 'A network error has occurred, please try again',
};
