import { useState } from 'react';
import { useAppSelector, useAppDispatch } from 'shared/redux/reduxHooks';
import ChannelDialogue from '../channel/channel';
import Drawer from '../drawer/drawer';
import Navbar from '../navbar/navbar';
import styles from './chat.module.scss';
import { Channel } from 'features/chat/redux/chatSlice';

export default function Chat() {
  const dispatch = useAppDispatch();

  return (
    <div className={styles.chat}>
      <Navbar />
      <Drawer />
      <ChannelDialogue />
    </div>
  );
}
