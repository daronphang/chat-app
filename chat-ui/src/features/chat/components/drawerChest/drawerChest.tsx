import { useState } from 'react';
import { enqueueSnackbar } from 'notistack';
import { RpcError } from 'grpc-web';

import userPb from 'proto/user/user_pb';
import commonPb from 'proto/common/common_pb';
import { useAppSelector } from 'core/redux/reduxHooks';
import { defaultSnackbarOptions } from 'core/config/snackbar.constant';
import { Channel } from 'features/chat/redux/chat.interface';
import Toolbar, { Display } from '../toolbar/toolbar';
import ChannelPanel from '../channelPanel/channelPanel';
import styles from './drawerChest.module.scss';
import NewChat from '../newChat/newChat';
import NewUser from '../newUser/newUser';
import NewGroup from '../newGroup/newGroup';

export default function DrawerChest() {
  const config = useAppSelector(state => state.config);
  const user = useAppSelector(state => state.user);
  const [display, setDisplay] = useState<Display>('channel');

  const handleClickToolbarButton = (v: Display) => {
    setDisplay(v);
  };

  const handleClickBack = () => {
    setDisplay('channel');
  };

  const createNewChannel = async (arg: Channel): Promise<Channel | null> => {
    try {
      const payload = new userPb.NewChannel();
      payload.setChannelname(arg.channelName);
      payload.setUseridsList(arg.userIds);
      const resp = await config.api.USER_SERVICE.createChannel(payload);

      enqueueSnackbar('New chat created', {
        ...defaultSnackbarOptions,
        variant: 'success',
      });

      // Update metadata.
      arg.channelId = resp.getChannelid();
      arg.createdAt = resp.getCreatedat();
      arg.updatedAt = resp.getCreatedat();
      return new Promise(resolve => resolve(arg));
    } catch (e) {
      const err = e as RpcError;
      const errMsg = err.code === 14 ? config.apiError.NETWORK_ERROR : 'Failed to create new chat';
      enqueueSnackbar(errMsg, {
        ...defaultSnackbarOptions,
        variant: 'error',
      });
      return new Promise(resolve => resolve(null));
    }
  };

  const broadcastChannelEvent = async (arg: Channel) => {
    const userIds = [...arg.userIds];
    try {
      // Remove current user.
      const idx = userIds.findIndex(row => row === user.userId);
      if (idx !== -1) {
        userIds.splice(idx, 1);
      }

      const payload = new commonPb.Channel();
      payload.setChannelid(arg.channelId);
      payload.setChannelname(arg.channelName);
      payload.setCreatedat(arg.createdAt);
      payload.setUseridsList(userIds);
      await config.api.SESSION_SERVICE.broadcastChannelEvent(payload);
    } catch (e) {
      const err = e as RpcError;
      console.error('failed to broadcast channel event', err.message);
    }
  };

  return (
    <div className={`${styles.drawer}`}>
      {display === 'channel' && (
        <div>
          <Toolbar handleClickToolbarButton={handleClickToolbarButton} />
          <div className="mb-4"></div>
          <ChannelPanel />
        </div>
      )}
      {display === 'newChat' && (
        <NewChat
          handleClickBack={handleClickBack}
          createNewChannel={createNewChannel}
          broadcastChannelEvent={broadcastChannelEvent}
        />
      )}
      {display === 'newUser' && <NewUser handleClickBack={handleClickBack} />}
      {display === 'newGroup' && (
        <NewGroup
          handleClickBack={handleClickBack}
          createNewChannel={createNewChannel}
          broadcastChannelEvent={broadcastChannelEvent}
        />
      )}
    </div>
  );
}
