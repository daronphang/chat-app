import DrawerCard from '../drawerCard/drawerCard';
import styles from './drawer.module.scss';
import { FontAwesomeIcon } from '@fortawesome/react-fontawesome';

export default function Drawer() {
  return (
    <div className={`${styles.drawer} p-3`}>
      <div className={`${styles.header} mb-4`}>
        <h3 className={styles.heading}>Chats</h3>
        <div className="flex-spacer"></div>
        <button className="btn-icon ms-3">
          <FontAwesomeIcon size="lg" icon={['fas', 'chalkboard-user']} />
        </button>
        <button className="btn-icon ms-3">
          <FontAwesomeIcon size="lg" icon={['fas', 'users-rectangle']} />
        </button>
      </div>
      <DrawerCard />
      <DrawerCard />
      <DrawerCard />
    </div>
  );
}
