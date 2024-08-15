import { FontAwesomeIcon } from '@fortawesome/react-fontawesome';
import { Badge } from '@mui/material';
import { RpcError } from 'grpc-web';
import { enqueueSnackbar } from 'notistack';

import { UserPresence } from 'proto/session/session_pb';
import { useAppDispatch, useAppSelector } from 'core/redux/reduxHooks';
import { defaultSnackbarOptions } from 'core/config/snackbar.constant';
import styles from './navbar.module.scss';
import { useEffect, useState } from 'react';
import { removeUser } from 'features/user/redux/userSlice';
import { useNavigate } from 'react-router-dom';
import { RoutePath } from 'core/config/route.constant';
import { addAppListener } from 'core/redux/listenerMiddleware';

type Status = 'online' | 'offline';

export default function Navbar() {
  const config = useAppSelector(state => state.config);
  const user = useAppSelector(state => state.user);
  const chat = useAppSelector(state => state.chat);
  const [unreadChannels, setUnreadChannels] = useState<number>(0);
  const dispatch = useAppDispatch();
  const navigate = useNavigate();

  useEffect(() => {
    (async () => {
      await broadcastUserPresence('online');
    })();

    setUnreadChannels(Object.keys(chat.unreadChannels).length);

    dispatch(
      addAppListener({
        predicate: action => {
          if (['chat/setChannelAsRead', 'chat/addUnreadChannel'].includes(action.type)) return true;
          return false;
        },
        effect: (action, listenerApi) => {
          const temp = listenerApi.getState().chat.unreadChannels;
          setUnreadChannels(Object.keys(temp).length);
        },
      })
    );
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
      await config.api.SESSION_SERVICE.broadcastUserPresenceEvent(payload);
    } catch (e) {
      const err = e as RpcError;
      console.error('failed to broadcast user presence', err.message);
    }
  };
  return (
    <div className={`${styles.navbar} p-3`}>
      <Badge badgeContent={unreadChannels} max={99} color="primary" className={styles.badgeWrapper}>
        <button className="btn-icon">
          <FontAwesomeIcon size="lg" icon={['fas', 'rectangle-list']} />
        </button>
      </Badge>
      <div className="flex-spacer"></div>
      <button className="btn-icon" onClick={handleLogOut}>
        <FontAwesomeIcon size="lg" icon={['fas', 'gear']} />
      </button>
    </div>
  );
}
