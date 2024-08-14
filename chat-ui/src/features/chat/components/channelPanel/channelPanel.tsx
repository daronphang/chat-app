import { useEffect, useState } from 'react';
import moment from 'moment';
import { useAppDispatch, useAppSelector } from 'core/redux/reduxHooks';
import { addAppListener } from 'core/redux/listenerMiddleware';
import { setCurChannelId } from 'features/chat/redux/chatSlice';
import { FriendHash } from 'features/user/redux/user.interface';
import { Channel } from 'features/chat/redux/chat.interface';
import Drawer from '../drawer/drawer';
import styles from './channelPanel.module.scss';

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
          if (['chat/addMessage', 'chat/addChannel', 'chat/updateDisplayName'].includes(action.type)) {
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

  const getSubtitle = (timestamp: string) => {
    const today = new Date();
    if (new Date(timestamp).setHours(0, 0, 0, 0) == today.setHours(0, 0, 0, 0)) {
      return moment(timestamp).format('HH:mm');
    }
    return moment(timestamp).format('DD/MM');
  };

  const displayChannels = (channels: Channel[]) => {
    // Channels are sorted in desc order.
    const drawers: JSX.Element[] = [];
    for (let i = channels.length - 1; i >= 0; i--) {
      const channel = channels[i];

      // For 1-on-1 chats without messages, to skip display if it is not a draft.
      if (channel.messages.length === 0 && channel.userIds.length === 2 && !channel.isDraft) {
        continue;
      }

      let text = '';
      let subtitle = '';

      if (channel.messages.length === 0) {
        text = 'Draft';
        subtitle = getSubtitle(channel.createdAt);
      } else {
        text = channel.messages[channel.messages.length - 1].content;
        const createdAt = channel.messages[channel.messages.length - 1].createdAt;
        subtitle = getSubtitle(createdAt);
      }

      drawers.push(
        <Drawer
          key={channel.channelId}
          data={channel}
          title={channel.channelName}
          subtitle={subtitle}
          text={text}
          friends={friends}
          handleClickDrawer={handleClickDrawer}
        />
      );
    }
    setDrawers(drawers);
  };

  return <div className={`${styles.drawerWrapper} p-3 pt-0`}>{drawers}</div>;
}
