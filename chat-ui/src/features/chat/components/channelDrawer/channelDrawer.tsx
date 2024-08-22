import { FontAwesomeIcon } from '@fortawesome/react-fontawesome';
import { useEffect, useRef, useState } from 'react';
import { Badge } from '@mui/material';
import moment from 'moment';
import { RpcError } from 'grpc-web';

import wrappers from 'google-protobuf/google/protobuf/wrappers_pb';
import { Channel, Message } from 'features/chat/redux/chat.interface';
import { useAppDispatch, useAppSelector } from 'core/redux/reduxHooks';
import { addChannel, addUnreadChannel, setChannelAsRead, setCurChannelId } from 'features/chat/redux/chatSlice';
import styles from './channelDrawer.module.scss';
import { fetchUnknownUsers, isGroupChat } from 'core/utils/chat';
import { Recipient } from 'features/user/redux/user.interface';
import { addRecipient, addRecipients } from 'features/user/redux/userSlice';

interface DrawerProps {
  channel: Channel;
}

export default function ChannelDrawer({ channel }: DrawerProps) {
  const [isOnline, setIsOnline] = useState<boolean>(false);
  const [title, setTitle] = useState<string>('');
  const [subtitle, setSubtitle] = useState<string>('');
  const [picture, setPicture] = useState<string>('');
  const [text, setText] = useState<string>('');
  const [badgeContent, setBadgeContent] = useState<number>(0);
  // For 1-on-1 chat, to store the friendId.
  const chatType = useRef<string>('group');

  const user = useAppSelector(state => state.user);
  const config = useAppSelector(state => state.config);
  const curChannelId = useAppSelector(state => state.chat.curChannelId);
  const dispatch = useAppDispatch();

  useEffect(() => {
    setChatType();
  }, []);

  useEffect(() => {
    handleNewMessage();
  }, [channel.messages]);

  useEffect(() => {
    handleRecipientStatusChange();
  }, [user.recipients]);

  useEffect(() => {
    setTitle(channel.channelName);
  }, [channel.channelName]);

  const setChatType = () => {
    if (isGroupChat(channel)) {
      return;
    }
    const friendId = channel.userIds.find(row => row !== user.userId);
    if (!friendId) {
      chatType.current = 'unknown';
    } else {
      chatType.current = friendId;
    }
  };

  const handleNewMessage = async () => {
    if (channel.messages.length === 0) {
      setText('Draft');
      setSubtitle(getSubtitle(channel.createdAt));
      return;
    }

    const latestMsg = channel.messages[channel.messages.length - 1];
    if (isGroupChat(channel)) {
      const name = await resolveRecipientIdToName(latestMsg.senderId);
      setText(`${name}: ${latestMsg.content}`);
    } else {
      setText(latestMsg.content);
    }
    setSubtitle(getSubtitle(latestMsg.createdAt));

    if (curChannelId !== channel.channelId) {
      const count = countUnreadMessages();
      setBadgeContent(count);
      if (count > 0) {
        dispatch(addUnreadChannel(channel.channelId));
      }
    }
  };

  const handleRecipientStatusChange = async () => {
    if (chatType.current === 'group') {
      return;
    }

    if (chatType.current in user.recipients && user.recipients[chatType.current].isFriend) {
      setIsOnline(user.recipients[chatType.current].isOnline);
    }

    const newChannelName = await resolveRecipientIdToName(chatType.current);
    if (channel.channelName !== newChannelName) {
      const payload: Channel = {
        ...channel,
        channelName: newChannelName,
        updatedAt: new Date().toISOString(),
      };
      dispatch(addChannel(payload));
    }
  };

  const getSubtitle = (timestamp: string) => {
    const today = new Date();
    if (new Date(timestamp).setHours(0, 0, 0, 0) == today.setHours(0, 0, 0, 0)) {
      return moment(timestamp).format('HH:mm');
    }
    return moment(timestamp).format('DD/MM');
  };

  // Applicable only for group chats.
  const resolveRecipientIdToName = async (recipientId: string): Promise<string> => {
    if (user.userId === recipientId) {
      return new Promise(resolve => resolve('You'));
    } else if (recipientId in user.recipients) {
      const recipient = user.recipients[recipientId];
      if (recipient.isFriend) return new Promise(resolve => resolve(recipient.friendName));
      return new Promise(resolve => resolve(recipient.email));
    }

    const unknownUser = await getUnknownUser(recipientId);
    if (!unknownUser) {
      return new Promise(resolve => resolve('Unknown User'));
    }
    dispatch(addRecipient(unknownUser));
    return new Promise(resolve => resolve(unknownUser.email));
  };

  const getUnknownUser = async (userId: string): Promise<Recipient | null> => {
    try {
      const payload = new wrappers.StringValue();
      payload.setValue(userId);
      const resp = await config.api.USER_SERVICE.getUser(payload);
      const rv: Recipient = {
        userId: resp.getUserid(),
        email: resp.getEmail(),
        displayName: resp.getDisplayname(),
        isFriend: false,
        friendName: '',
        isOnline: false,
        color: '#000000',
      };
      return new Promise(resolve => resolve(rv));
    } catch (e) {
      const err = e as RpcError;
      console.error('unable to get unknown user', err.message);
      return new Promise(resolve => resolve(null));
    }
  };

  const countUnreadMessages = (): number => {
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

  const handleClickDrawer = async () => {
    // All users in a group are fetched when the chat is clicked.
    if (chatType.current === 'group') {
      const userIds = channel.userIds.filter(row => row !== user.userId && !(row in user.recipients));
      const resp = await fetchUnknownUsers(config, userIds);
      if (resp && resp.length > 0) {
        dispatch(addRecipients(resp));
      }
    }
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
