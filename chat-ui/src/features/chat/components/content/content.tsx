import { Message, MessageStatus } from 'features/chat/redux/chat.interface';
import styles from './content.module.scss';
import { useAppSelector } from 'core/redux/reduxHooks';
import { FontAwesomeIcon } from '@fortawesome/react-fontawesome';
import moment from 'moment';
import { useEffect, useState } from 'react';

interface ContentProps {
  message: Message;
  isGroup: boolean;
  recipientLastReadMessageId?: number;
}

export default function Content({ message, isGroup, recipientLastReadMessageId }: ContentProps) {
  const user = useAppSelector(state => state.user);
  const [sender, setSender] = useState<string>('');
  const [color, setColor] = useState<string>('#000000');

  useEffect(() => {
    if (isGroup && message.senderId in user.recipients) {
      setSender(resolveSenderIdToName(message.senderId));
      setColor(user.recipients[message.senderId].color);
    }
  }, []);

  // Applicable only for group chats.
  const resolveSenderIdToName = (senderId: string): string => {
    if (user.userId === senderId) {
      return '';
    } else if (senderId in user.recipients) {
      const sender = user.recipients[senderId];
      if (sender.isFriend) {
        return sender.friendName;
      }
      return sender.email;
    }
    return 'Unknown User';
  };

  return (
    <div className={`${user.userId === message.senderId ? styles.senderContent : styles.receiverContent} mb-3 p-2`}>
      {sender && (
        <div style={{ color: color }}>
          <strong>{sender}</strong>
        </div>
      )}
      {message.content}
      <div className={styles.metadata}>
        <span className="me-2">{moment(message.createdAt).format('MM/DD/YY HH:mm')}</span>
        {user.userId === message.senderId && message.messageStatus === MessageStatus.PENDING && (
          <FontAwesomeIcon className={styles.icon} size="lg" icon={['fas', 'clock']} />
        )}
        {user.userId === message.senderId && message.messageStatus === MessageStatus.RECEIVED && (
          <FontAwesomeIcon className={styles.icon} size="lg" icon={['fas', 'check']} />
        )}
        {user.userId === message.senderId && message.messageStatus === MessageStatus.DELIVERED && (
          <FontAwesomeIcon className={styles.icon} size="lg" icon={['fas', 'check-double']} />
        )}
        {user.userId === message.senderId && message.messageStatus === MessageStatus.READ && (
          <FontAwesomeIcon className={styles.read} size="lg" icon={['fas', 'check-double']} />
        )}
      </div>
    </div>
  );
}
