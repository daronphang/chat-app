import { useEffect, useState } from 'react';
import { Alert } from '@mui/material';

import Dialogue from '../dialogue/dialogue';
import Drawer from '../drawerChest/drawerChest';
import Navbar from '../navbar/navbar';
import styles from './chat.module.scss';
import StartUp from '../startUp/startUp';
import { useAppDispatch, useAppSelector } from 'core/redux/reduxHooks';
import { UserIds, UserSession } from 'proto/notification/notification_pb';
import { RpcError } from 'grpc-web';
import { enqueueSnackbar } from 'notistack';
import { defaultSnackbarOptions } from 'core/config/snackbar.constant';
import { updateOnlineFriends } from 'features/user/redux/userSlice';

export default function Chat() {
  const [loading, setLoading] = useState<boolean>(true);
  const [alert, setAlert] = useState<string>('');
  const config = useAppSelector(state => state.config);
  const user = useAppSelector(state => state.user);
  const dispatch = useAppDispatch();

  useEffect(() => {
    const interval = setInterval(async () => {
      await Promise.all([sendClientHeartbeat()]);
    }, 5000);

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

  const sendClientHeartbeat = async () => {
    if (config.chatServerWsUrl && user.userId) {
      try {
        const payload = new UserSession();
        payload.setUserid(user.userId);
        payload.setServer(config.chatServerWsUrl);
        await config.api.NOTIFICATION_SERVICE.clientHeartbeat(payload);
      } catch (e) {
        const err = e as RpcError;
        if (err.code === 14) {
          enqueueSnackbar(config.apiError.NETWORK_ERROR, {
            ...defaultSnackbarOptions,
            variant: 'error',
          });
        } else {
          console.error('failed to send client heartbeat', err.message);
        }
      }
    }
  };

  const fetchOnlineFriends = async () => {
    try {
      const payload = new UserIds();
      const friendIds = Object.keys(user.friends);
      payload.setUseridsList(friendIds);
      const resp = await config.api.NOTIFICATION_SERVICE.getOnlineUsers(payload);
      dispatch(updateOnlineFriends(resp.getUseridsList()));
    } catch (e) {
      const err = e as RpcError;
      if (err.code === 14) {
        enqueueSnackbar(config.apiError.NETWORK_ERROR, {
          ...defaultSnackbarOptions,
          variant: 'error',
        });
      } else {
        console.error('failed to fetch online status of friends', err.message);
      }
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
          <Dialogue />
        </div>
      )}
    </>
  );
}
