import { Tooltip } from '@mui/material';
import { FontAwesomeIcon } from '@fortawesome/react-fontawesome';
import { useEffect, useState } from 'react';

import { useAppDispatch, useAppSelector } from 'core/redux/reduxHooks';
import Drawer from '../drawer/drawer';
import styles from './newChat.module.scss';
import { Friend } from 'features/user/redux/userSlice';
import { addNewChannel, Channel, setCurChannelId } from 'features/chat/redux/chatSlice';
import { NewChannel } from 'proto/user/user_pb';
import { RpcError } from 'grpc-web';
import { enqueueSnackbar } from 'notistack';
import { defaultSnackbarOptions } from 'core/config/snackbar.constant';

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
    friends.sort((a, b) => a.displayName.localeCompare(b.displayName));
    setDrawers(
      friends.map(row => (
        <Drawer key={row.userId} data={row} title={row.displayName} text="" handleClickDrawer={handleClickDrawer} />
      ))
    );
  }, []);

  const handleClickDrawer = async (v: Friend) => {
    const key = [v.userId, user.userId].sort().join('');
    const idx = chat.channels.findIndex(row => row.channelId === key);

    if (idx === -1) {
      const newChannel: Channel = {
        channelId: key,
        channelName: v.displayName,
        createdAt: new Date().toISOString(),
        messages: [],
        userIds: [v.userId, user.userId],
      };
      await createNewChannel(newChannel);
    }
    dispatch(setCurChannelId(key));
    handleClickBack();
  };

  const createNewChannel = async (newChannel: Channel) => {
    try {
      const payload = new NewChannel();
      payload.setChannelid(newChannel.channelId);
      payload.setChannelname(newChannel.channelName);
      payload.setUseridsList(newChannel.userIds);
      const resp = await config.api.USER_SERVICE.createChannel(payload);

      // Update required metadata.
      newChannel.channelId = resp.getChannelid();
      newChannel.createdAt = resp.getCreatedat();

      dispatch(addNewChannel(newChannel));
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
    }
  };

  return (
    <div>
      <div className={styles.headerWrapper}>
        <Tooltip title="New Friend" placement="bottom">
          <button className="btn-icon ms-3" onClick={() => handleClickBack()}>
            <FontAwesomeIcon size="lg" icon={['fas', 'arrow-left']} />
          </button>
        </Tooltip>
        <div className={`${styles.heading} ms-3`}>New Chat</div>
      </div>
      <div className="mb-4"></div>
      {drawers}
    </div>
  );
}
