import { useState } from 'react';
import Convo, { ConvoProps } from '../convo/convo';
import Drawer from '../drawer/drawer';
import Navbar from '../navbar/navbar';
import styles from './chat.module.scss';

export default function Chat() {
  const [activeConvo, setActiveConvo] = useState<ConvoProps>({ displayName: 'John Doe', channelId: 'test123' });

  return (
    <div className={styles.chat}>
      <Navbar />
      <Drawer />
      <Convo displayName={activeConvo.displayName} channelId={activeConvo.channelId} />
    </div>
  );
}
