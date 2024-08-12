import { useAppDispatch, useAppSelector } from 'core/redux/reduxHooks';
import { useSnackbar } from 'notistack';
import { useEffect } from 'react';
import wrappers from 'google-protobuf/google/protobuf/wrappers_pb';
import empty from 'google-protobuf/google/protobuf/empty_pb';
import { RpcError } from 'grpc-web';

import { setChatServerWsUrl } from 'core/config/configSlice';
import { defaultSnackbarOptions } from 'core/config/snackbar.constant';
import { Channel, Message } from 'features/chat/redux/chat.interface';
import { initChannels } from 'features/chat/redux/chatSlice';
import { chunk } from 'shared/utils/formatters';
import styles from './startUp.module.scss';

interface StartUpProps {
  handleLoading: (v: boolean) => void;
  handleAlert: (v: string) => void;
}

export default function StartUp({ handleLoading, handleAlert }: StartUpProps) {
  const { enqueueSnackbar } = useSnackbar();
  const dispatch = useAppDispatch();
  const config = useAppSelector(state => state.config);
  const user = useAppSelector(state => state.user);

  useEffect(() => {
    async function handleStartUp() {
      const promise = await Promise.all([fetchChannelsAssociatedToUser(user.userId), fetchChatServerWsUrl()]);

      const channels = promise[0];
      if (channels.length === 0) return;
      await loadMessagesInChannels(channels);
    }
    handleStartUp();
    setTimeout(() => {
      handleLoading(false);
    }, 2000);
  }, []);

  const loadMessagesInChannels = async (channels: Channel[]) => {
    // Buffer channel requests and updates to Redux store.
    const chunks: Channel[][] = chunk(channels, 5);

    for (let chunk of chunks) {
      const promise = await Promise.all(chunk.map(row => fetchLatestMessagesForChannel(row.channelId)));
      // Update messages for each channel.
      for (let i = 0; i < chunk.length; i++) {
        chunk[i].messages = promise[i];
      }
    }
    dispatch(initChannels(channels));
  };

  const fetchChannelsAssociatedToUser = async (userId: string): Promise<Channel[]> => {
    try {
      const payload = new wrappers.StringValue();
      payload.setValue(userId);
      const resp = await config.api.USER_SERVICE.getChannelsAssociatedToUser(payload);
      const channels: Channel[] = resp.getChannelsList().map(row => {
        // Warning: implementation detail required to extract friend id.
        const userIds: string[] = [user.userId];
        if (row.getChannelid().includes(user.userId)) {
          userIds.push(row.getChannelid().replace(user.userId, ''));
        }

        return {
          channelId: row.getChannelid(),
          channelName: row.getChannelname(),
          createdAt: row.getCreatedat(),
          messages: [],
          userIds,
        };
      });
      return new Promise(resolve => resolve(channels));
    } catch (e) {
      const err = e as RpcError;
      if (err.code === 14) {
        enqueueSnackbar(config.apiError.NETWORK_ERROR, {
          ...defaultSnackbarOptions,
          variant: 'error',
        });
      } else {
        console.error(err.message);
        enqueueSnackbar('Unable to retrieve user chats, please refresh the page', {
          ...defaultSnackbarOptions,
          variant: 'error',
        });
      }
      return new Promise(resolve => resolve([]));
    }
  };

  const fetchLatestMessagesForChannel = async (channelId: string): Promise<Message[]> => {
    try {
      const payload = new wrappers.StringValue();
      payload.setValue(channelId);
      const resp = await config.api.MESSAGE_SERVICE.getLatestMessages(payload);

      // Messages are fetched in ascending order.
      const messages: Message[] = resp.getMessagesList().map(row => {
        return {
          messageId: row.getMessageid(),
          channelId: row.getChannelid(),
          senderId: row.getSenderid(),
          messageType: row.getMessagetype(),
          content: row.getContent(),
          createdAt: row.getCreatedat(),
          messageStatus: row.getMessagestatus(),
        };
      });
      return new Promise(resolve => resolve(messages));
    } catch (e) {
      const err = e as RpcError;
      if (err.code === 14) {
        enqueueSnackbar(config.apiError.NETWORK_ERROR, {
          ...defaultSnackbarOptions,
          variant: 'error',
        });
      } else {
        console.error(err.message);
        enqueueSnackbar('Some messages from chats could not be retrieved, please refresh the page', {
          ...defaultSnackbarOptions,
          variant: 'error',
        });
      }
      return new Promise(resolve => resolve([]));
    }
  };

  const fetchChatServerWsUrl = async () => {
    try {
      const payload = new empty.Empty();
      const resp = await config.api.USER_SERVICE.getBestServer(payload);
      dispatch(setChatServerWsUrl(resp.getMessage()));
    } catch (e) {
      const err = e as RpcError;
      if (err.code === 14) {
        enqueueSnackbar(config.apiError.NETWORK_ERROR, {
          ...defaultSnackbarOptions,
          variant: 'error',
        });
      } else {
        console.error(err.message);
        enqueueSnackbar('Could not connect to chat server, please refresh the page', {
          ...defaultSnackbarOptions,
          variant: 'error',
        });
        handleAlert('Could not connect to chat server, please refresh the page');
      }
    }
  };

  return (
    <div className={styles.wrapper}>
      <div className={styles.loader}></div>
      <div className={`mt-3`}>Loading your chats...</div>
    </div>
  );
}
