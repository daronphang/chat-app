import { FontAwesomeIcon } from '@fortawesome/react-fontawesome';

import styles from './drawer.module.scss';
import { useEffect, useState } from 'react';
import { FriendHash } from 'features/user/redux/user.interface';
import { Channel } from 'features/chat/redux/chat.interface';
import { useAppSelector } from 'core/redux/reduxHooks';

interface DrawerProps {
  title: string;
  subtitle?: string;
  text: string;
  data: any;
  friends?: FriendHash;
  handleClickDrawer: (data: any) => void;
}

export default function Drawer({ title, subtitle, text, data, friends, handleClickDrawer }: DrawerProps) {
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
      <div className={styles.iconWrapper}>
        {isOnline && <FontAwesomeIcon className={styles.status} size="2xs" icon={['fas', 'circle']} />}
        <FontAwesomeIcon size="3x" icon={['fas', 'circle-user']} />
      </div>
      <div className={`${styles.bodyWrapper}`}>
        <div className={styles.headingWrapper}>
          <div className={`${styles.title} truncated`}>{title}</div>
          <div className="flex-spacer"></div>
          {subtitle && <div className={`${styles.subtitle} me-3`}>{subtitle}</div>}
        </div>
        <div className="truncated">{text}</div>
      </div>
    </div>
  );
}
