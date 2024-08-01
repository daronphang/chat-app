import styles from './navbar.module.scss';
import { FontAwesomeIcon } from '@fortawesome/react-fontawesome';
import { Badge } from '@mui/material';

export default function Navbar() {
  return (
    <div className={`${styles.navbar} p-3`}>
      <Badge badgeContent={4} max={99} color="primary" className={styles.badgeWrapper}>
        <button className="btn-icon">
          <FontAwesomeIcon size="lg" icon={['fas', 'rectangle-list']} />
        </button>
      </Badge>
      <div className="flex-spacer"></div>
      <button className="btn-icon">
        <FontAwesomeIcon size="lg" icon={['fas', 'gear']} />
      </button>
    </div>
  );
}
