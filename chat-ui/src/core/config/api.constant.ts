import { Config } from './config.constant';
import { UserClient } from 'proto/user/UserServiceClientPb';
import { MessageClient } from 'proto/message/MessageServiceClientPb';
import { NotificationClient } from 'proto/notification/NotificationServiceClientPb';

export class BaseApi {
  public USER_SERVICE: UserClient;
  public MESSAGE_SERVICE: MessageClient;
  public NOTIFICATION_SERVICE: NotificationClient;

  constructor() {}

  public initWithConfig(config: Config) {
    this.USER_SERVICE = new UserClient(config.ENVOY_PROXY_ADDRESS);
    this.MESSAGE_SERVICE = new MessageClient(config.ENVOY_PROXY_ADDRESS);
    this.NOTIFICATION_SERVICE = new NotificationClient(config.ENVOY_PROXY_ADDRESS);
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
