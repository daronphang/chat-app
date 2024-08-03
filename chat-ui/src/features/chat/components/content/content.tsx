import { Message } from 'features/chat/redux/chatSlice';
import styles from './content.module.scss';
import { useAppSelector } from 'shared/redux/reduxHooks';
import { FontAwesomeIcon } from '@fortawesome/react-fontawesome';

interface ContentProps {
  props: Message;
}

export default function Content({ props }: ContentProps) {
  const userId = useAppSelector(state => state.user.userId);

  return (
    <div className={`${userId === props.senderId ? styles.senderContent : styles.receiverContent} mb-3 p-2`}>
      {props.content}
      <div className={styles.metadata}>
        <span className="me-2">{props.createdAt}</span>
        {props.delivery && props.delivery == 0 && <FontAwesomeIcon size="lg" icon={['fas', 'clock']} />}
        {props.delivery && props.delivery == 1 && <FontAwesomeIcon size="lg" icon={['fas', 'check']} />}
        {props.delivery && props.delivery == 2 && <FontAwesomeIcon size="lg" icon={['fas', 'check-double']} />}
      </div>
    </div>
  );
}
