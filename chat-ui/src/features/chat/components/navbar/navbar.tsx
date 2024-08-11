import { FontAwesomeIcon } from '@fortawesome/react-fontawesome';
import { Badge } from '@mui/material';
import { RpcError } from 'grpc-web';
import { enqueueSnackbar } from 'notistack';

import { UserPresence } from 'proto/notification/notification_pb';
import { useAppDispatch, useAppSelector } from 'core/redux/reduxHooks';
import { defaultSnackbarOptions } from 'core/config/snackbar.constant';
import styles from './navbar.module.scss';
import { useEffect } from 'react';
import { removeUser } from 'features/user/redux/userSlice';
import { useNavigate } from 'react-router-dom';
import { RoutePath } from 'core/config/route.constant';

type Status = 'online' | 'offline';

export default function Navbar() {
  const config = useAppSelector(state => state.config);
  const user = useAppSelector(state => state.user);
  const dispatch = useAppDispatch();
  const navigate = useNavigate();

  useEffect(() => {
    (async () => {
      await broadcastUserPresence('online');
    })();
  }, []);

  const handleLogOut = async () => {
    await broadcastUserPresence('offline');
    dispatch(removeUser());
    navigate(RoutePath.LOGIN);
  };

  const broadcastUserPresence = async (status: Status) => {
    try {
      const payload = new UserPresence();
      payload.setUserid(user.userId);
      payload.setStatus(status);
      await config.api.NOTIFICATION_SERVICE.broadcastUserPresenceEvent(payload);
    } catch (e) {
      const err = e as RpcError;
      if (err.code === 14) {
        enqueueSnackbar(config.apiError.NETWORK_ERROR, {
          ...defaultSnackbarOptions,
          variant: 'error',
        });
      } else {
        console.error('failed to send broadcast user presence', err.message);
      }
    }
  };
  return (
    <div className={`${styles.navbar} p-3`}>
      <Badge badgeContent={4} max={99} color="primary" className={styles.badgeWrapper}>
        <button className="btn-icon">
          <FontAwesomeIcon size="lg" icon={['fas', 'rectangle-list']} />
        </button>
      </Badge>
      <div className="flex-spacer"></div>
      <button className="btn-icon">
        <FontAwesomeIcon size="lg" icon={['fas', 'gear']} />
      </button>
    </div>
  );
}
