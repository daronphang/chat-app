import { Tooltip } from '@mui/material';
import { FontAwesomeIcon } from '@fortawesome/react-fontawesome';

import styles from './toolbar.module.scss';

export type Display = 'channel' | 'newChat' | 'newGroup' | 'newUser';

interface ToolbarProps {
  handleClickToolbarButton: (v: Display) => void;
}

export default function Toolbar({ handleClickToolbarButton }: ToolbarProps) {
  return (
    <div className={`${styles.toolbarWrapper}`}>
      <h3 className={styles.heading}>Chats</h3>
      <div className="flex-spacer"></div>
      <Tooltip title="New Friend" placement="bottom">
        <button className="btn-icon ms-3" onClick={() => handleClickToolbarButton('newUser')}>
          <FontAwesomeIcon size="lg" icon={['fas', 'user-plus']} />
        </button>
      </Tooltip>
      <Tooltip title="New Chat" placement="bottom">
        <button className="btn-icon ms-3" onClick={() => handleClickToolbarButton('newChat')}>
          <FontAwesomeIcon size="lg" icon={['fas', 'chalkboard-user']} />
        </button>
      </Tooltip>
      <Tooltip title="New Group" placement="bottom">
        <button className="btn-icon ms-3" onClick={() => handleClickToolbarButton('newGroup')}>
          <FontAwesomeIcon size="lg" icon={['fas', 'users-rectangle']} />
        </button>
      </Tooltip>
    </div>
  );
}
