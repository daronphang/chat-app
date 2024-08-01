import styles from './drawerCard.module.scss';
import { FontAwesomeIcon } from '@fortawesome/react-fontawesome';

export default function DrawerCard() {
  return (
    <div className={`${styles.drawerCard} gap-3`}>
      <FontAwesomeIcon size="3x" icon={['fas', 'circle-user']} />
      <div className={`${styles.bodyWrapper}`}>
        <div className={styles.header}>Hello world</div>
        <div className="truncated">Some very long message here, hello there how are you hello there how are you?</div>
      </div>
    </div>
  );
}
