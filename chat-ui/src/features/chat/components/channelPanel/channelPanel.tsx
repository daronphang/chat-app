import { useEffect, useState } from 'react';
import { useAppSelector } from 'core/redux/reduxHooks';
import styles from './channelPanel.module.scss';
import ChannelDrawer from '../channelDrawer/channelDrawer';
import { PriorityQueue } from '@datastructures-js/priority-queue';

interface ChannelSort {
  channelId: string;
  key: number;
  messageSize: number;
}

export default function ChannelPanel() {
  const [drawers, setDrawers] = useState<JSX.Element[]>([]);
  const channelHash = useAppSelector(state => state.chat.channelHash);

  useEffect(() => {
    displayDrawers();
  }, [channelHash]);

  const displayDrawers = () => {
    // Priority queue for sorting.
    const pq = new PriorityQueue<ChannelSort>((a, b) => {
      return b.key - a.key;
    });
    Object.values(channelHash).forEach(row => {
      let timestamp = new Date(row.createdAt);
      if (row.messages.length > 0) {
        timestamp = new Date(row.messages[row.messages.length - 1].createdAt);
      }
      const temp: ChannelSort = {
        channelId: row.channelId,
        key: timestamp.getTime(),
        messageSize: row.messages.length,
      };
      pq.enqueue(temp);
    });

    const newDrawers: JSX.Element[] = [];
    let row: ChannelSort;
    while (pq.size() > 0) {
      row = pq.dequeue();

      if (row.channelId.length > 36 && row.messageSize === 0) {
        continue;
      } else {
        newDrawers.push(<ChannelDrawer key={row.channelId} channelId={row.channelId} />);
      }
    }

    setDrawers(newDrawers);
  };

  return <div className={`${styles.drawerWrapper} p-3 pt-0`}>{drawers}</div>;
}
