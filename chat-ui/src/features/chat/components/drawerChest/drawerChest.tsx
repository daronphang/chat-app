import { useState } from 'react';

import NewChat from '../newChat/newChat';
import Toolbar, { Display } from '../toolbar/toolbar';
import ChannelPanel from '../channelPanel/channelPanel';
import styles from './drawerChest.module.scss';

export default function DrawerChest() {
  const [display, setDisplay] = useState<Display>('channel');

  const handleClickToolbarButton = (v: Display) => {
    setDisplay(v);
  };

  const handleClickBack = () => {
    setDisplay('channel');
  };

  return (
    <div className={`${styles.drawer} p-3`}>
      {display === 'channel' && (
        <div>
          <Toolbar handleClickToolbarButton={handleClickToolbarButton} />
          <div className="mb-4"></div>
          <ChannelPanel />
        </div>
      )}
      {display === 'newChat' && <NewChat handleClickBack={handleClickBack} />}
    </div>
  );
}
