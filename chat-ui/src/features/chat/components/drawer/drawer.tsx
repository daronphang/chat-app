import { FontAwesomeIcon } from '@fortawesome/react-fontawesome';

import styles from './drawer.module.scss';
import { useEffect, useState } from 'react';
import { FriendHash } from 'features/user/redux/userSlice';
import { Channel } from 'features/chat/redux/chatSlice';
import { useAppSelector } from 'core/redux/reduxHooks';

interface DrawerProps {
  title: string;
  text: string;
  data: any;
  friends?: FriendHash;
  handleClickDrawer: (data: any) => void;
}

export default function Drawer({ title, text, data, friends, handleClickDrawer }: DrawerProps) {
  const [isOnline, setIsOnline] = useState<boolean>(false);
  const user = useAppSelector(state => state.user);

  useEffect(() => {
    updateOnlineStatus();
  }, [friends]);

  const isInstanceOfChannel = (v: any): v is Channel => {
    return 'channelName' in v;
  };

  const updateOnlineStatus = () => {
    if (!friends || !isInstanceOfChannel(data) || data.userIds.length !== 2) {
      return;
    }

    const friendId = data.userIds.filter(row => row !== user.userId)[0];
    if (friendId in friends) {
      const friend = friends[friendId];
      setIsOnline(friend.isOnline ? true : false);
    }
  };

  return (
    <div onClick={() => handleClickDrawer(data)} className={`${styles.drawer} gap-3`}>
      <FontAwesomeIcon className={styles.iconWrapper} size="3x" icon={['fas', 'circle-user']} />
      <div className={`${styles.bodyWrapper}`}>
        <div className={styles.header}>{title}</div>
        <div className="truncated">{text}</div>
      </div>
    </div>
  );
}
