import { Tooltip } from '@mui/material';
import { FontAwesomeIcon } from '@fortawesome/react-fontawesome';
import { useEffect, useState } from 'react';

import { useAppDispatch, useAppSelector } from 'core/redux/reduxHooks';
import { Friend } from 'features/user/redux/user.interface';
import { Channel } from 'features/chat/redux/chat.interface';
import { addChannel, setCurChannelId } from 'features/chat/redux/chatSlice';
import Search from 'shared/components/search/search';
import Drawer from '../drawer/drawer';
import styles from './newChat.module.scss';

interface NewChatProps {
  handleClickBack: () => void;
  createNewChannel: (arg: Channel) => Promise<Channel | null>;
  broadcastChannelEvent: (arg: Channel) => void;
}

export default function NewChat({ handleClickBack, createNewChannel, broadcastChannelEvent }: NewChatProps) {
  const [drawers, setDrawers] = useState<JSX.Element[]>([]);
  const user = useAppSelector(state => state.user);
  const chat = useAppSelector(state => state.chat);
  const config = useAppSelector(state => state.config);
  const dispatch = useAppDispatch();

  useEffect(() => {
    const friends: Friend[] = Object.values(user.friends);
    displayFriends(friends);
  }, []);

  const handleClickDrawer = async (v: Friend) => {
    // If channel exists, to route to that channel.
    // Else, create a new channel.
    let key = [v.userId, user.userId].sort().join('');
    const idx = chat.channels.findIndex(row => row.channelId === key);

    if (idx === -1) {
      const newChannel: Channel = {
        channelId: key,
        channelName: user.displayName, // For the benefit of the recipient.
        messages: [],
        userIds: [v.userId, user.userId],
        isDraft: true,
        createdAt: new Date().toISOString(),
        updatedAt: new Date().toISOString(),
      };
      const resp = await createNewChannel(newChannel);
      if (!resp) return;

      await broadcastChannelEvent(resp);

      resp.channelName = v.displayName;
      dispatch(addChannel(resp));
    } else {
      const channel = chat.channels[idx];
      key = channel.channelId;

      // For existing 1-on-1 chats with no messages, to set isDraft to true
      // and move channel to the front.
      if (channel.messages.length === 0 && channel.userIds.length == 2) {
        const updatedChannel: Channel = {
          ...channel,
          createdAt: new Date().toISOString(),
          updatedAt: new Date().toISOString(),
          isDraft: true,
        };
        dispatch(addChannel(updatedChannel));
      }
    }
    dispatch(setCurChannelId(key));
    handleClickBack();
  };

  const displayFriends = (friends: Friend[]) => {
    friends.sort((a, b) => a.displayName.localeCompare(b.displayName));
    setDrawers(
      friends.map(row => (
        <Drawer key={row.userId} data={row} title={row.displayName} text="" handleClickDrawer={handleClickDrawer} />
      ))
    );
  };

  const handleSearchResult = (v: any[]) => {
    displayFriends(v);
  };

  return (
    <div>
      <div className={styles.headerWrapper}>
        <Tooltip title="Back" placement="bottom">
          <button className="btn-icon ms-3" onClick={() => handleClickBack()}>
            <FontAwesomeIcon size="lg" icon={['fas', 'arrow-left']} />
          </button>
        </Tooltip>
        <div className={`${styles.heading} ms-3`}>New Chat</div>
      </div>
      <div className="mb-4"></div>
      <Search
        sourceData={Object.values(user.friends)}
        excludedFields={['userId', 'email', 'isOnline']}
        handleSearchResult={handleSearchResult}
      />
      <div className="mb-3"></div>
      {drawers}
    </div>
  );
}
