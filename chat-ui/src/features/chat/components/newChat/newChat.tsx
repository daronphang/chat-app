import { Tooltip } from '@mui/material';
import { FontAwesomeIcon } from '@fortawesome/react-fontawesome';
import { useEffect, useState } from 'react';
import { enqueueSnackbar } from 'notistack';
import { RpcError } from 'grpc-web';

import { NewChannel } from 'proto/user/user_pb';
import { useAppDispatch, useAppSelector } from 'core/redux/reduxHooks';
import { defaultSnackbarOptions } from 'core/config/snackbar.constant';
import { Friend } from 'features/user/redux/user.interface';
import { Channel } from 'features/chat/redux/chat.interface';
import { addNewChannel, setCurChannelId, updateChannel } from 'features/chat/redux/chatSlice';
import Search from 'shared/components/search/search';
import Drawer from '../drawer/drawer';
import styles from './newChat.module.scss';

interface NewChatProps {
  handleClickBack: () => void;
}

export default function NewChat({ handleClickBack }: NewChatProps) {
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
    let key = [v.userId, user.userId].sort().join('');
    const idx = chat.channels.findIndex(row => row.channelId === key);

    if (idx === -1) {
      const newChannel: Channel = {
        channelId: key,
        channelName: v.displayName,
        createdAt: new Date().toISOString(),
        messages: [],
        userIds: [v.userId, user.userId],
        isDraft: true,
      };
      const resp = await createNewChannel(newChannel);
      if (!resp) return;
    } else {
      const channel = chat.channels[idx];
      key = channel.channelId;

      if (channel.messages.length === 0 && channel.userIds.length == 2) {
        const updatedChannel: Channel = {
          channelId: channel.channelId,
          channelName: channel.channelName,
          createdAt: new Date().toISOString(),
          messages: channel.messages,
          userIds: channel.userIds,
          isDraft: true,
        };
        dispatch(updateChannel(updatedChannel));
      }
    }
    dispatch(setCurChannelId(key));
    handleClickBack();
  };

  const createNewChannel = async (newChannel: Channel): Promise<boolean> => {
    try {
      const payload = new NewChannel();
      payload.setChannelid(newChannel.channelId);
      payload.setChannelname(user.displayName); // For benefit of recipient.
      payload.setUseridsList(newChannel.userIds);
      const resp = await config.api.USER_SERVICE.createChannel(payload);

      // Update required metadata.
      newChannel.channelId = resp.getChannelid();
      newChannel.createdAt = resp.getCreatedat();
      dispatch(addNewChannel(newChannel));
      enqueueSnackbar('New chat created', {
        ...defaultSnackbarOptions,
        variant: 'success',
      });
      return new Promise(resolve => resolve(true));
    } catch (e) {
      const err = e as RpcError;
      if (err.code === 14) {
        enqueueSnackbar(config.apiError.NETWORK_ERROR, {
          ...defaultSnackbarOptions,
          variant: 'error',
        });
      } else {
        enqueueSnackbar('Failed to create new chat', {
          ...defaultSnackbarOptions,
          variant: 'error',
        });
        console.error(err.message);
      }
      return new Promise(resolve => resolve(false));
    }
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
