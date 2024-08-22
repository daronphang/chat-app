import { Channel } from 'features/chat/redux/chat.interface';
import { Recipient } from 'features/user/redux/user.interface';
import { getRandomColor } from './formatters';
import { RpcError } from 'grpc-web';
import userPb from 'proto/user/user_pb';
import sessionPb from 'proto/session/session_pb';
import { ConfigState } from 'core/constants/configSlice';

export const isGroupChat = (channel: Channel): boolean => {
  if (channel.channelId.length <= 36) {
    return true;
  }
  return false;
};

// Implementation detail.
export const getRecipientId = (userId: string, channelId: string): string => {
  if (channelId.includes(userId)) {
    const recipientId = channelId.replace(userId, '');
    return recipientId;
  }
  return '';
};

export const getRecipientIds = (userId: string, channels: Channel[]): string[] => {
  const recipientIds: string[] = [];
  channels.forEach(row => {
    if (row.channelId.includes(userId)) {
      const recipientId = row.channelId.replace(userId, '');
      recipientIds.push(recipientId);
    }
  });
  return recipientIds;
};

export const fetchUnknownUsers = async (config: ConfigState, userIds: string[]): Promise<Recipient[] | null> => {
  if (userIds.length === 0) return new Promise(resolve => resolve(null));
  try {
    const payload = new userPb.UserIds();
    payload.setUseridsList(userIds);
    const resp = await config.api.USER_SERVICE.getUsers(payload);
    const recipients: Recipient[] = resp.getUsersList().map(row => {
      return {
        userId: row.getUserid(),
        email: row.getEmail(),
        displayName: row.getDisplayname(),
        isFriend: false,
        friendName: '',
        isOnline: false,
        color: getRandomColor(),
      };
    });
    return new Promise(resolve => resolve(recipients));
  } catch (e) {
    const err = e as RpcError;
    console.error('failed to fetch unknown users', err.message);
    return new Promise(resolve => resolve(null));
  }
};

export const fetchOnlineRecipients = async (config: ConfigState, recipientIds: string[]): Promise<string[]> => {
  if (recipientIds.length === 0) {
    return new Promise(resolve => resolve([]));
  }
  try {
    const payload = new sessionPb.UserIds();
    payload.setUseridsList(recipientIds);
    const resp = await config.api.SESSION_SERVICE.getOnlineUsers(payload);
    return new Promise(resolve => resolve(resp.getUseridsList()));
  } catch (e) {
    const err = e as RpcError;
    console.error('failed to fetch online status of friends', err.message);
    return new Promise(resolve => resolve([]));
  }
};
