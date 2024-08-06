import { useEffect, useState } from 'react';
import { FontAwesomeIcon } from '@fortawesome/react-fontawesome';

import { Channel } from 'features/chat/redux/chatSlice';
import { useAppDispatch, useAppSelector } from 'core/redux/reduxHooks';
import { addAppListener } from 'core/redux/listenerMiddleware';
import DrawerCard from '../drawerCard/drawerCard';
import styles from './drawer.module.scss';

export default function Drawer() {
  const [channels, setChannels] = useState<JSX.Element[]>([]);
  const dispatch = useAppDispatch();
  const head = useAppSelector(state => state.chat.channelHead);

  useEffect(() => {
    if (head) {
      updateChannelsFromHead(head);
    }
  }, []);

  useEffect(() => {
    dispatch(
      addAppListener({
        predicate: action => {
          if (['chat/setHead', 'chat/addNewMessage', 'chat/addNewChannel'].includes(action.type)) {
            return true;
          }
          return false;
        },
        effect: (action, listenerApi) => {
          const head = listenerApi.getState().chat.channelHead;
          if (head) {
            updateChannelsFromHead(head);
          }
        },
      })
    );
  }, []);

  const updateChannelsFromHead = (head: Channel) => {
    // Need to iterate through linked list of channels for DOM update.
    // This is because the linked list will be reordered and not just appended.
    let cur: Channel | null = head;
    const channels: JSX.Element[] = [];
    while (cur) {
      channels.push(<DrawerCard key={cur.channelId} props={cur} />);
      cur = cur.prev;
    }
    setChannels(channels);
  };

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
      {channels}
    </div>
  );
}
