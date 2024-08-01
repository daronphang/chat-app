import styles from './convo.module.scss';
import { FontAwesomeIcon } from '@fortawesome/react-fontawesome';

export interface ConvoProps {
  displayName: string;
  channelId: string;
}

interface Conversation {}

function Convo({ displayName, channelId }: ConvoProps) {
  return (
    <div>
      <div className={`${styles.header} p-3`}>
        <FontAwesomeIcon size="2x" icon={['fas', 'circle-user']} />
        <div className={`${styles.heading} ms-3`}>{displayName}</div>
      </div>
      <div className={`${styles.convo} p-3`}></div>
    </div>
  );
}

export default Convo;
