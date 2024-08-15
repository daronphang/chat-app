import { FontAwesomeIcon } from '@fortawesome/react-fontawesome';
import { useEffect, useRef, useState } from 'react';
import { Badge } from '@mui/material';
import moment from 'moment';

import { Message, MessageStatus } from 'features/chat/redux/chat.interface';
import { useAppDispatch, useAppSelector } from 'core/redux/reduxHooks';
import { addUnreadChannel, setChannelAsRead, setCurChannelId } from 'features/chat/redux/chatSlice';
import styles from './channelDrawer.module.scss';

interface DrawerProps {
  channelId: string;
}

export default function ChannelDrawer({ channelId }: DrawerProps) {
  const [isOnline, setIsOnline] = useState<boolean>(false);
  const [title, setTitle] = useState<string>('');
  const [subtitle, setSubtitle] = useState<string>('');
  const [picture, setPicture] = useState<string>('');
  const [text, setText] = useState<string>('');
  const [badgeContent, setBadgeContent] = useState<number>(0);
  // For 1-on-1 chat, to store the friendId.
  const chatType = useRef<string>('');

  const user = useAppSelector(state => state.user);
  const channel = useAppSelector(state => state.chat.channelHash[channelId]);
  const curChannelId = useAppSelector(state => state.chat.curChannelId);
  const dispatch = useAppDispatch();

  useEffect(() => {
    if (!chatType.current) {
      setChatType();
    }

    setTitle(getTitle());
    if (channel.messages.length === 0) {
      setText('Draft');
      setSubtitle(getSubtitle(channel.createdAt));
    } else {
      setText(channel.messages[channel.messages.length - 1].content);
      setSubtitle(getSubtitle(channel.messages[channel.messages.length - 1].createdAt));

      if (curChannelId !== channel.channelId) {
        const count = getUnreadMessages();
        setBadgeContent(count);
        if (count > 0) {
          dispatch(addUnreadChannel(channel.channelId));
        }
      }
    }
  }, [channel]);

  useEffect(() => {
    if (!chatType.current || chatType.current === 'group') {
      return;
    }

    if (chatType.current in user.friends) {
      setIsOnline(user.friends[chatType.current].isOnline);
    }
  }, [user.friends]);

  const setChatType = () => {
    // Group channel.
    if (channel.channelId.length <= 36) {
      chatType.current = 'group';
      return;
    }

    const friendId = channel.userIds.find(row => row !== user.userId);
    if (!friendId) {
      return;
    }
    chatType.current = friendId;
  };

  const getSubtitle = (timestamp: string) => {
    const today = new Date();
    if (new Date(timestamp).setHours(0, 0, 0, 0) == today.setHours(0, 0, 0, 0)) {
      return moment(timestamp).format('HH:mm');
    }
    return moment(timestamp).format('DD/MM');
  };

  const getTitle = (): string => {
    if (chatType.current === 'group') {
      return channel.channelName;
    }

    // Check if user is friend.
    if (chatType.current in user.friends) {
      return user.friends[chatType.current].displayName;
    }
    // User is not friend. To fetch user's name instead.
    return '';
  };

  const fetchUser = async () => {};

  const getUnreadMessages = (): number => {
    let count = 0;
    let cur: Message;
    for (let i = channel.messages.length - 1; i >= 0; i--) {
      cur = channel.messages[i];

      if (cur.senderId === user.userId) {
        continue;
      } else if (channel.lastMessageId < cur.messageId) {
        count += 1;
      } else {
        break;
      }
    }
    return count;
  };

  const handleClickDrawer = () => {
    if (badgeContent > 0) {
      setBadgeContent(0);
      dispatch(setChannelAsRead(channel.channelId));
    }
    dispatch(setCurChannelId(channel.channelId));
  };

  return (
    <div onClick={handleClickDrawer} className={`${styles.channelDrawer} gap-3`}>
      <div className={styles.iconWrapper}>
        {isOnline && <FontAwesomeIcon className={styles.status} size="2xs" icon={['fas', 'circle']} />}
        <FontAwesomeIcon size="3x" icon={['fas', 'circle-user']} />
      </div>
      <div className={`${styles.bodyWrapper}`}>
        <div className={styles.headingWrapper}>
          <div className={`${styles.title} truncated`}>{title}</div>
          <div className="flex-spacer"></div>
          {subtitle && <div className={`${styles.subtitle} me-3`}>{subtitle}</div>}
        </div>
        <div className={styles.headingWrapper}>
          <div className="truncated">{text}</div>
          <div className="flex-spacer"></div>
          <Badge className="me-4" badgeContent={badgeContent} color="primary" max={99}></Badge>
        </div>
      </div>
    </div>
  );
}
