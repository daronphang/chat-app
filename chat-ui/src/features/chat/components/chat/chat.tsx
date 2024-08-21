import { useEffect, useState } from 'react';
import { Alert } from '@mui/material';
import { RpcError } from 'grpc-web';
import useWebSocket from 'react-use-websocket';

import sessionPb from 'proto/session/session_pb';
import { useAppDispatch, useAppSelector } from 'core/redux/reduxHooks';
import { defaultWsOptions } from 'core/config/ws.constant';
import { updateFriendPresence, updateOnlineRecipients } from 'features/user/redux/userSlice';
import { Channel, Message, WebSocketEvent } from 'features/chat/redux/chat.interface';
import { addChannel, addMessage } from 'features/chat/redux/chatSlice';
import Dialogue from '../dialogue/dialogue';
import Drawer from '../drawerChest/drawerChest';
import Navbar from '../navbar/navbar';
import styles from './chat.module.scss';
import StartUp from '../startUp/startUp';
import { UserPresence } from 'features/user/redux/user.interface';
import { fetchOnlineRecipients, getRecipientIds } from 'core/utils/chat';

export default function Chat() {
  const [loading, setLoading] = useState<boolean>(true);
  const [alert, setAlert] = useState<string>('');
  const config = useAppSelector(state => state.config);
  const user = useAppSelector(state => state.user);
  const chat = useAppSelector(state => state.chat);
  const dispatch = useAppDispatch();
  const { sendJsonMessage } = useWebSocket(config.chatServerWsUrl, {
    ...defaultWsOptions,
    filter: event => {
      const data = JSON.parse(event.data) as WebSocketEvent;
      if (['event/message', 'event/channel', 'event/presence'].includes(data.event)) {
        return true;
      }
      return false;
    },
    onMessage: event => {
      const data = JSON.parse(event.data) as WebSocketEvent;
      if (data.event === 'event/message') {
        handleMessageEvent(data.data, data.eventTimestamp);
      } else if (data.event === 'event/channel') {
        handleChannelEvent(data.data, data.eventTimestamp);
      } else {
        handlePresenceEvent(data.data);
      }
    },
    onError: error => {
      console.error(error);
    },
  });

  useEffect(() => {
    if (loading) {
      return;
    }

    const interval = setInterval(async () => {
      await Promise.all([sendClientHeartbeat()]);
    }, 10000);

    (async () => {
      await Promise.all([
        sendClientHeartbeat(),
        broadcastUserPresence('online', getRecipientIds(user.userId, chat.channels)),
      ]);

      const resp = await fetchOnlineRecipients(config, Object.keys(user.recipients));
      dispatch(updateOnlineRecipients(resp));
    })();

    document.addEventListener('visibilitychange', handleVisibilityChange);

    return () => {
      clearInterval(interval);
      document.removeEventListener('visibilitychange', handleVisibilityChange);
    };
  }, [loading]);

  const handleLoading = (v: boolean) => {
    setLoading(v);
  };

  const handleAlert = (v: string) => {
    setAlert(v);
  };

  const handleMessageEvent = (msg: Message, updatedAt: string) => {
    msg.updatedAt = updatedAt;
    dispatch(addMessage(msg));
  };

  const handleChannelEvent = (channel: Channel, updatedAt: string) => {
    channel.updatedAt = updatedAt;
    channel.messages = [];
    dispatch(addChannel(channel));
  };

  const handlePresenceEvent = (v: UserPresence) => {
    dispatch(updateFriendPresence(v));
  };

  const sendClientHeartbeat = async () => {
    try {
      const payload = new sessionPb.UserSession();
      payload.setUserid(user.userId);
      payload.setServer(config.chatServerAddress);
      await config.api.SESSION_SERVICE.clientHeartbeat(payload);
    } catch (e) {
      const err = e as RpcError;
      console.error('failed to send client heartbeat', err.message);
    }
  };

  const broadcastUserPresence = async (status: string, recipientIds: string[]) => {
    if (recipientIds.length === 0) {
      return;
    }
    try {
      const payload = new sessionPb.UserPresence();
      payload.setUserid(user.userId);
      payload.setStatus(status);
      payload.setRecipientidsList(recipientIds);
      await config.api.SESSION_SERVICE.broadcastUserPresenceEvent(payload);
    } catch (e) {
      const err = e as RpcError;
      console.error('failed to broadcast user presence', err.message);
    }
  };

  const handleVisibilityChange = async () => {
    if (document.hidden) {
      await broadcastUserPresence('offline', getRecipientIds(user.userId, chat.channels));
    } else {
      await broadcastUserPresence('online', getRecipientIds(user.userId, chat.channels));
    }
  };

  return (
    <>
      {loading && <StartUp handleLoading={handleLoading} handleAlert={handleAlert} />}
      {!loading && alert && <Alert severity="error">{alert}</Alert>}
      {!loading && (
        <div className={styles.chat}>
          <Navbar broadcastUserPresence={broadcastUserPresence} />
          <Drawer />
          <Dialogue sendJsonMessage={sendJsonMessage} />
        </div>
      )}
    </>
  );
}
