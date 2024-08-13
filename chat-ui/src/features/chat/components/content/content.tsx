import { Message } from 'features/chat/redux/chat.interface';
import styles from './content.module.scss';
import { useAppSelector } from 'core/redux/reduxHooks';
import { FontAwesomeIcon } from '@fortawesome/react-fontawesome';
import moment from 'moment';

interface ContentProps {
  props: Message;
}

export default function Content({ props }: ContentProps) {
  const userId = useAppSelector(state => state.user.userId);

  return (
    <div className={`${userId === props.senderId ? styles.senderContent : styles.receiverContent} mb-3 p-2`}>
      {props.content}
      <div className={styles.metadata}>
        <span className="me-2">{moment(props.createdAt).format('MM/DD/YY HH:mm')}</span>
        {props.messageStatus === 0 && <FontAwesomeIcon className={styles.icon} size="lg" icon={['fas', 'clock']} />}
        {props.messageStatus === 1 && <FontAwesomeIcon className={styles.icon} size="lg" icon={['fas', 'check']} />}
        {props.messageStatus === 2 && (
          <FontAwesomeIcon className={styles.icon} size="lg" icon={['fas', 'check-double']} />
        )}
      </div>
    </div>
  );
}
