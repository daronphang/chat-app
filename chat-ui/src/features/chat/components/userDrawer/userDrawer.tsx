import { FontAwesomeIcon } from '@fortawesome/react-fontawesome';
import { useEffect, useState } from 'react';
import { Friend } from 'features/user/redux/user.interface';
import { Channel } from 'features/chat/redux/chat.interface';
import { useAppSelector } from 'core/redux/reduxHooks';
import styles from './userDrawer.module.scss';

interface DrawerProps {
  title: string;
  subtitle?: string;
  text: string;
  data: Friend;
  handleClickDrawer: (data: Friend) => void;
}

export default function UserDrawer({ title, subtitle, text, data, handleClickDrawer }: DrawerProps) {
  const [isOnline, setIsOnline] = useState<boolean>(false);
  const user = useAppSelector(state => state.user);

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
        <div className={styles.headingWrapper}>
          <div className="truncated">{text}</div>
          <div className="flex-spacer"></div>
        </div>
      </div>
    </div>
  );
}
