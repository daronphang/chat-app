import { Tooltip } from '@mui/material';
import { FontAwesomeIcon } from '@fortawesome/react-fontawesome';
import { useEffect, useRef, useState } from 'react';

import { useAppDispatch, useAppSelector } from 'core/redux/reduxHooks';
import { Recipient } from 'features/user/redux/user.interface';
import { Channel } from 'features/chat/redux/chat.interface';
import { addChannel, moveChannelToFront, setCurChannelId } from 'features/chat/redux/chatSlice';
import Search from 'shared/components/search/search';
import UserDrawer from '../userDrawer/userDrawer';
import styles from './newChat.module.scss';
import { isGroupChat } from 'core/utils/chat';

interface NewChatProps {
  handleClickBack: () => void;
  createNewChannel: (arg: Channel) => Promise<Channel | null>;
  broadcastChannelEvent: (arg: Channel) => void;
}

export default function NewChat({ handleClickBack, createNewChannel, broadcastChannelEvent }: NewChatProps) {
  const [drawers, setDrawers] = useState<JSX.Element[]>([]);
  const user = useAppSelector(state => state.user);
  const chat = useAppSelector(state => state.chat);
  const friends = useRef<Recipient[]>([]);
  const dispatch = useAppDispatch();

  useEffect(() => {
    friends.current = Object.values(user.recipients).filter(row => row.isFriend);
    displayFriends(friends.current);
  }, []);

  const handleClickDrawer = async (v: Recipient) => {
    // If channel exists, to route to that channel.
    // Else, create a new channel.
    let key = [v.userId, user.userId].sort().join('');
    let channel = chat.channels.find(row => row.channelId === key);

    if (!channel) {
      const timestamp = new Date().toISOString();
      channel = {
        channelId: key,
        channelName: v.displayName,
        messages: [],
        userIds: [v.userId, user.userId],
        createdAt: timestamp,
        updatedAt: timestamp,
        lastMessageId: 0,
      };
      const resp = await createNewChannel(channel);
      if (!resp) return;
      channel = resp;
      dispatch(addChannel(channel));
    } else {
      // For existing 1-on-1 chats with no messages, to set isDraft to true
      // and move channel to the front.
      if (!isGroupChat(channel) && channel.messages.length === 0) {
        dispatch(moveChannelToFront(channel.channelId));
      }
    }
    dispatch(setCurChannelId(channel.channelId));
    handleClickBack();
  };

  const displayFriends = (friends: Recipient[]) => {
    friends.sort((a, b) => a.friendName.localeCompare(b.friendName));
    setDrawers(
      friends.map(row => (
        <UserDrawer key={row.userId} data={row} title={row.friendName} text="" handleClickDrawer={handleClickDrawer} />
      ))
    );
  };

  const handleSearchResult = (v: any[]) => {
    displayFriends(v);
  };

  return (
    <>
      <div className="p-3">
        <div className={styles.headerWrapper}>
          <Tooltip title="Back" placement="bottom">
            <button className="btn-icon ms-3" onClick={() => handleClickBack()}>
              <FontAwesomeIcon size="lg" icon={['fas', 'arrow-left']} />
            </button>
          </Tooltip>
          <div className={`${styles.heading} ms-3`}>New Chat</div>
        </div>
        <div className="mb-4"></div>
        {friends.current.length > 0 && (
          <Search
            sourceData={friends.current}
            excludedFields={['userId', 'email', 'isOnline', 'isFriend', 'displayName']}
            handleSearchResult={handleSearchResult}
          />
        )}
        <div className="mb-3"></div>
      </div>
      {friends.current.length === 0 && <div className="text-center">To create a chat, add a friend.</div>}
      {friends.current.length > 0 && <div className={`${styles.drawerWrapper} p-3`}>{drawers}</div>}
    </>
  );
}
