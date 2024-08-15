import { Message, MessageStatus } from 'features/chat/redux/chat.interface';
import styles from './content.module.scss';
import { useAppSelector } from 'core/redux/reduxHooks';
import { FontAwesomeIcon } from '@fortawesome/react-fontawesome';
import moment from 'moment';
import { useEffect } from 'react';

interface ContentProps {
  props: Message;
  recipientLastReadMessageId?: number;
}

export default function Content({ props, recipientLastReadMessageId }: ContentProps) {
  const userId = useAppSelector(state => state.user.userId);

  return (
    <div className={`${userId === props.senderId ? styles.senderContent : styles.receiverContent} mb-3 p-2`}>
      {props.content}
      <div className={styles.metadata}>
        <span className="me-2">{moment(props.createdAt).format('MM/DD/YY HH:mm')}</span>
        {userId === props.senderId && props.messageStatus === MessageStatus.PENDING && (
          <FontAwesomeIcon className={styles.icon} size="lg" icon={['fas', 'clock']} />
        )}
        {userId === props.senderId && props.messageStatus === MessageStatus.RECEIVED && (
          <FontAwesomeIcon className={styles.icon} size="lg" icon={['fas', 'check']} />
        )}
        {userId === props.senderId && props.messageStatus === MessageStatus.DELIVERED && (
          <FontAwesomeIcon className={styles.icon} size="lg" icon={['fas', 'check-double']} />
        )}
        {userId === props.senderId && props.messageStatus === MessageStatus.READ && (
          <FontAwesomeIcon className={styles.read} size="lg" icon={['fas', 'check-double']} />
        )}
      </div>
    </div>
  );
}
