import { useEffect } from 'react';
import { useAppSelector, useAppDispatch } from 'core/redux/reduxHooks';
import wrappers from 'google-protobuf/google/protobuf/wrappers_pb';
import { useSnackbar } from 'notistack';
import { RpcError } from 'grpc-web';

import { ApiErrors } from 'core/config/api.constant';
import { Channel, Message, initChannels, setHead } from 'features/chat/redux/chatSlice';
import { UserMetadata, addContacts } from 'features/user/redux/userSlice';
import { chunk } from 'shared/utils/formatters';
import { defaultSnackbarOptions } from 'shared/utils/snackbar';
import ChannelDialogue from '../channel/channel';
import Drawer from '../drawer/drawer';
import Navbar from '../navbar/navbar';
import styles from './chat.module.scss';

export default function Chat() {
  const { enqueueSnackbar } = useSnackbar();
  const dispatch = useAppDispatch();
  const api = useAppSelector(state => state.config.api);
  const user = useAppSelector(state => state.user);

  useEffect(() => {
    async function fetchRequestsInParallel() {
      // Fetch contacts and channels.
      const promise = await Promise.all([getUserContacts(user.userId), fetchChannelsAssociatedToUser(user.userId)]);
      const channels = promise[1];

      if (channels.length === 0) return;

      // Buffer channel requests and updates to Redux store.
      const chunks: Channel[][] = chunk(channels, 5);

      for (let chunk of chunks) {
        const promise = await Promise.all(chunk.map(row => fetchLatestMessagesForChannel(row.channelId)));
        // Update messages for each channel.
        for (let i = 0; i < chunk.length; i++) {
          chunk[i].messages = promise[i];
        }
      }

      // Order channels and set head.
      const head = sortChannelsAndReturnHead(channels);
      dispatch(initChannels(channels));
      dispatch(setHead(head));
    }
    fetchRequestsInParallel();

    // Connect to websocket.
  }, []);

  const getUserContacts = async (userId: string): Promise<null> => {
    try {
      const payload = new wrappers.StringValue();
      payload.setValue(userId);
      const resp = await api.USER_SERVICE.getContacts(payload);
      const contacts: UserMetadata[] = resp.getContactsList().map(row => {
        return {
          userId: row.getUserid(),
          email: row.getEmail(),
          displayName: row.getDisplayname(),
          contacts: [],
        };
      });
      dispatch(addContacts(contacts));
      return new Promise(resolve => resolve(null));
    } catch (e) {
      const err = e as RpcError;
      if (err.code === 14) {
        enqueueSnackbar(ApiErrors.NETWORK_ERROR, {
          ...defaultSnackbarOptions(),
          variant: 'error',
        });
      } else {
        console.error(err.message);
        enqueueSnackbar('Unable to retrieve contacts, please refresh the page', {
          ...defaultSnackbarOptions(),
          variant: 'error',
        });
      }

      return new Promise(resolve => resolve(null));
    }
  };

  const fetchChannelsAssociatedToUser = async (userId: string): Promise<Channel[]> => {
    try {
      const payload = new wrappers.StringValue();
      payload.setValue(userId);
      const resp = await api.USER_SERVICE.getChannelsAssociatedToUser(payload);
      const channels = resp.getChannelsList().map(row => {
        return {
          channelId: row.getChannelid(),
          channelName: row.getChannelname(),
          createdAt: row.getCreatedat(),
          messages: [],
          prev: null,
          next: null,
        };
      });
      return new Promise(resolve => resolve(channels));
    } catch (e) {
      const err = e as RpcError;
      if (err.code === 14) {
        enqueueSnackbar(ApiErrors.NETWORK_ERROR, {
          ...defaultSnackbarOptions(),
          variant: 'error',
        });
      } else {
        console.error(err.message);
        enqueueSnackbar('Unable to retrieve user chats, please refresh the page', {
          ...defaultSnackbarOptions(),
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
      const resp = await api.MESSAGE_SERVICE.getLatestMessages(payload);

      // Messages returned are in descending order.
      // Need to sort by ascending i.e. latest message last.
      const messages: Message[] = [];
      const temp = resp.getMessagesList();
      for (let i = temp.length - 1; i >= 0; i--) {
        messages.push({
          messageId: temp[i].getMessageid(),
          channelId: temp[i].getChannelid(),
          senderId: temp[i].getSenderid(),
          messageType: temp[i].getMessagetype(),
          content: temp[i].getContent(),
          createdAt: temp[i].getCreatedat(),
          delivery: 2,
        });
      }
      return new Promise(resolve => resolve(messages));
    } catch (e) {
      const err = e as RpcError;
      if (err.code === 14) {
        enqueueSnackbar(ApiErrors.NETWORK_ERROR, {
          ...defaultSnackbarOptions(),
          variant: 'error',
        });
      } else {
        console.error(err.message);
        enqueueSnackbar('Some messages from chats could not be retrieved, please refresh the page', {
          ...defaultSnackbarOptions(),
          variant: 'error',
        });
      }
      return new Promise(resolve => resolve([]));
    }
  };

  const sortChannelsAndReturnHead = (channels: Channel[]): Channel => {
    // Sort by latest timestamp.
    channels.sort((a, b) => {
      let ts1: Date;
      let ts2: Date;

      if (a.messages.length === 0) {
        ts1 = new Date(a.createdAt);
      } else {
        ts1 = new Date(a.messages[a.messages.length - 1].createdAt);
      }

      if (b.messages.length === 0) {
        ts2 = new Date(b.createdAt);
      } else {
        ts2 = new Date(b.messages[b.messages.length - 1].createdAt);
      }
      return ts2.getTime() - ts1.getTime();
    });

    const head: Channel = channels[0];
    let cur: Channel = head;
    channels.forEach((row, idx) => {
      if (idx !== 0) {
        cur.prev = row;
        cur = cur.prev;
      }
    });
    return head;
  };

  return (
    <div className={styles.chat}>
      <Navbar />
      <Drawer />
      <ChannelDialogue />
    </div>
  );
}
