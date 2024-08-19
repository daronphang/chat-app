import { FontAwesomeIcon } from '@fortawesome/react-fontawesome';
import { Badge } from '@mui/material';

import { useAppDispatch, useAppSelector } from 'core/redux/reduxHooks';
import { useEffect, useState } from 'react';
import { resetUser } from 'features/user/redux/userSlice';
import { resetChat } from 'features/chat/redux/chatSlice';
import { useNavigate } from 'react-router-dom';
import { RoutePath } from 'core/config/route.constant';
import { addAppListener } from 'core/redux/listenerMiddleware';
import { getRecipientIds } from 'core/utils/chat';
import styles from './navbar.module.scss';
import { resetConfig } from 'core/config/configSlice';

interface NavbarProps {
  broadcastUserPresence: (status: string, recipientIds: string[]) => void;
}

export default function Navbar({ broadcastUserPresence }: NavbarProps) {
  const user = useAppSelector(state => state.user);
  const chat = useAppSelector(state => state.chat);
  const [unreadChannels, setUnreadChannels] = useState<number>(0);
  const dispatch = useAppDispatch();
  const navigate = useNavigate();

  useEffect(() => {
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
    await broadcastUserPresence('offline', getRecipientIds(user.userId, chat.channels));
    dispatch(resetUser());
    dispatch(resetChat());
    dispatch(resetConfig());
    navigate(RoutePath.LOGIN);
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
