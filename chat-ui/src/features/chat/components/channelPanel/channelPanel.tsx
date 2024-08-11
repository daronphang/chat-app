import { addAppListener } from 'core/redux/listenerMiddleware';
import { useAppDispatch, useAppSelector } from 'core/redux/reduxHooks';
import { Channel, setCurChannelId } from 'features/chat/redux/chatSlice';
import { useEffect, useState } from 'react';
import Drawer from '../drawer/drawer';
import { FriendHash } from 'features/user/redux/userSlice';

export default function ChannelPanel() {
  const [drawers, setDrawers] = useState<JSX.Element[]>([]);
  const [friends, setFriends] = useState<FriendHash>({});
  const chat = useAppSelector(state => state.chat);
  const dispatch = useAppDispatch();

  useEffect(() => {
    if (chat.channels.length > 0) {
      displayChannels(chat.channels);
    }
  }, []);

  useEffect(() => {
    dispatch(
      addAppListener({
        predicate: action => {
          if (['chat/addNewMessage', 'chat/addNewChannel'].includes(action.type)) {
            return true;
          }
          return false;
        },
        effect: (action, listenerApi) => {
          const channels = listenerApi.getState().chat.channels;
          displayChannels(channels);
        },
      })
    );
    dispatch(
      addAppListener({
        predicate: action => {
          if (['user/updateOnlineFriends', 'user/updateFriendPresence'].includes(action.type)) return true;
          return false;
        },
        effect: (action, listenerApi) => {
          // Instead of having a listener in each drawer, to have one listener in parent.
          // Updated friends will be propagated down as props.
          setFriends(listenerApi.getState().user.friends);
        },
      })
    );
  }, []);

  const handleClickDrawer = (data: Channel) => {
    dispatch(setCurChannelId(data.channelId));
  };

  const displayChannels = (channels: Channel[]) => {
    // Channels are sorted in desc order.
    const drawers: JSX.Element[] = [];
    for (let i = channels.length - 1; i >= 0; i--) {
      const channel = channels[i];
      const text = channel.messages.length === 0 ? 'Draft' : channel.messages[channel.messages.length - 1].content;
      drawers.push(
        <Drawer
          key={channel.channelId}
          data={channel}
          title={channel.channelName}
          text={text}
          friends={friends}
          handleClickDrawer={handleClickDrawer}
        />
      );
    }
    setDrawers(drawers);
  };

  return <>{drawers}</>;
}
