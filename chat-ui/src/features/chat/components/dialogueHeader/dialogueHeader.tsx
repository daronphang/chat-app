import { useEffect, useState } from 'react';
import { FontAwesomeIcon } from '@fortawesome/react-fontawesome';
import { Channel } from 'features/chat/redux/chat.interface';
import styles from './dialogueHeader.module.scss';
import { useAppSelector } from 'core/redux/reduxHooks';
import { getRecipientId } from 'core/utils/chat';

interface DialogueHeaderProps {
  channel: Channel;
}

export default function DialogueHeader({ channel }: DialogueHeaderProps) {
  const [title, setTitle] = useState<string>('');
  const [isOnline, setIsOnline] = useState<boolean>(false);
  const user = useAppSelector(state => state.user);

  useEffect(() => {
    setTitle(channel.channelName);
  }, [channel]);

  useEffect(() => {
    const recipientId = getRecipientId(user.userId, channel.channelId);
    if (recipientId in user.recipients) {
      const recipient = user.recipients[recipientId];
      if (recipient.isFriend) {
        setIsOnline(recipient.isOnline);
      }
    }
  }, [channel.channelId, user.recipients]);

  return (
    <div className={`${styles.headerWrapper} p-3`}>
      <FontAwesomeIcon className={styles.iconWrapper} size="3x" icon={['fas', 'circle-user']} />
      <div className={`${styles.textWrapper} ms-3`}>
        <div className={`${styles.heading}`}>{title}</div>
        {isOnline && <div>Online</div>}
      </div>
    </div>
  );
}
