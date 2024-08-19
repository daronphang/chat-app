import { useEffect, useRef, useState } from 'react';
import { useAppSelector } from 'core/redux/reduxHooks';
import ChannelDrawer from '../channelDrawer/channelDrawer';
import styles from './channelPanel.module.scss';

export default function ChannelPanel() {
  const [drawers, setDrawers] = useState<JSX.Element[]>([]);
  const channels = useAppSelector(state => state.chat.channels);
  const curChannelId = useAppSelector(state => state.chat.curChannelId);
  const initiatedConvos = useRef<Set<string>>(new Set());

  useEffect(() => {
    if (curChannelId) {
      initiatedConvos.current.add(curChannelId);
    }
    displayDrawers();
  }, [channels]);

  const displayDrawers = () => {
    // Channels will be moved up to the front for the following:
    // 1. New messages
    // 2. New channel
    // Other events (name/picture change, new members, message/channel updates, etc) will be ignored.
    // Channels will already be sorted by Redux.
    const newDrawers: JSX.Element[] = [];
    channels.forEach(row => {
      if (row.channelId.length > 36 && row.messages.length === 0 && !initiatedConvos.current.has(row.channelId)) {
        // Hide 1-on-1 chats with no messages if user did not initiate a conversation.
      } else {
        newDrawers.push(<ChannelDrawer key={row.channelId} channel={row} />);
      }
    });

    setDrawers(newDrawers);
  };

  return <div className={`${styles.drawerWrapper} p-3 pt-0`}>{drawers}</div>;
}
