import { FontAwesomeIcon } from '@fortawesome/react-fontawesome';
import { Channel } from 'features/chat/redux/chatSlice';
import styles from './drawerCard.module.scss';
import { useAppDispatch } from 'shared/redux/reduxHooks';
import { setCurChannel } from 'features/chat/redux/chatSlice';

interface DrawerCardProps {
  props: Channel;
}

export default function DrawerCard({ props }: DrawerCardProps) {
  const dispatch = useAppDispatch();

  const handleClick = () => {
    dispatch(setCurChannel(props));
  };

  return (
    <div onClick={handleClick} className={`${styles.drawerCard} gap-3`}>
      <FontAwesomeIcon size="3x" icon={['fas', 'circle-user']} />
      <div className={`${styles.bodyWrapper}`}>
        <div className={styles.header}>{props.channelName}</div>
        <div className="truncated">{props.messages[props.messages.length - 1].content}</div>
      </div>
    </div>
  );
}
