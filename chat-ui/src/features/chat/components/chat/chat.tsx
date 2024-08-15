import { useEffect, useState } from 'react';
import { Alert } from '@mui/material';
import { RpcError } from 'grpc-web';
import useWebSocket from 'react-use-websocket';

import { UserIds, UserSession } from 'proto/session/session_pb';
import { useAppDispatch, useAppSelector } from 'core/redux/reduxHooks';
import { defaultWsOptions } from 'core/config/ws.constant';
import { updateOnlineFriends } from 'features/user/redux/userSlice';
import { Channel, Message, WebSocketEvent } from 'features/chat/redux/chat.interface';
import { addChannel, addMessage } from 'features/chat/redux/chatSlice';
import Dialogue from '../dialogue/dialogue';
import Drawer from '../drawerChest/drawerChest';
import Navbar from '../navbar/navbar';
import styles from './chat.module.scss';
import StartUp from '../startUp/startUp';

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
      if (['event/message', 'event/channel'].includes(data.event)) {
        return true;
      }
      return false;
    },
    onMessage: event => {
      const data = JSON.parse(event.data) as WebSocketEvent;
      if (data.event === 'event/message') {
        handleMessageEvent(data.data, data.eventTimestamp);
      } else {
        handleChannelEvent(data.data, data.eventTimestamp);
      }
    },
    onError: error => {
      console.error(error);
    },
  });

  useEffect(() => {
    const interval = setInterval(async () => {
      await Promise.all([sendClientHeartbeat()]);
    }, 10000);

    (async () => {
      await fetchOnlineFriends();
    })();

    () => {
      clearInterval(interval);
    };
  }, []);

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
    dispatch(addChannel(channel));
  };

  const sendClientHeartbeat = async () => {
    try {
      const payload = new UserSession();
      payload.setUserid(user.userId);
      payload.setServer(config.chatServerWsUrl);
      await config.api.SESSION_SERVICE.clientHeartbeat(payload);
    } catch (e) {
      const err = e as RpcError;
      console.error('failed to send client heartbeat', err.message);
    }
  };

  const fetchOnlineFriends = async () => {
    try {
      const payload = new UserIds();
      const friendIds = Object.keys(user.friends);
      payload.setUseridsList(friendIds);
      const resp = await config.api.SESSION_SERVICE.getOnlineUsers(payload);
      dispatch(updateOnlineFriends(resp.getUseridsList()));
    } catch (e) {
      const err = e as RpcError;
      console.error('failed to fetch online status of friends', err.message);
    }
  };

  return (
    <>
      {loading && <StartUp handleLoading={handleLoading} handleAlert={handleAlert} />}
      {!loading && alert && <Alert severity="error">{alert}</Alert>}
      {!loading && (
        <div className={styles.chat}>
          <Navbar />
          <Drawer />
          <Dialogue sendJsonMessage={sendJsonMessage} />
        </div>
      )}
    </>
  );
}
