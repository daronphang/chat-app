import { KeyboardEvent, useEffect, useRef, useState } from 'react';
import { FontAwesomeIcon } from '@fortawesome/react-fontawesome';
import { useForm } from 'react-hook-form';
import useWebSocket from 'react-use-websocket';

import AppLogo from 'assets/images/chatLogo.png';
import { addAppListener } from 'core/redux/listenerMiddleware';
import { useAppDispatch, useAppSelector } from 'core/redux/reduxHooks';
import { defaultWsOptions } from 'core/config/ws.constant';
import { FriendHash } from 'features/user/redux/user.interface';
import { addNewMessage, setCurChannelId } from 'features/chat/redux/chatSlice';
import { Channel, Message, WebSocketEvent } from 'features/chat/redux/chat.interface';
import Content from '../content/content';
import styles from './dialogue.module.scss';

interface FormInput {
  content: string;
}

export default function Dialogue() {
  const [content, setContent] = useState<JSX.Element[]>([]);
  const [isOnline, setIsOnline] = useState<boolean>(false);
  const curChannel = useRef<Channel | null>(null);
  const dialogueEndRef = useRef<null | HTMLDivElement>(null);
  const user = useAppSelector(state => state.user);
  const wsUrl = useAppSelector(state => {
    return `${state.config.chatServerWsUrl}?client=${user.userId}&device=${state.config.deviceId}`;
  });
  const { sendJsonMessage } = useWebSocket(wsUrl, {
    ...defaultWsOptions,
    filter: event => {
      const temp = JSON.parse(event.data) as WebSocketEvent;
      if (temp.event === 'event/message') {
        return true;
      }
      return false;
    },
    onMessage: event => {
      const data = JSON.parse(event.data) as WebSocketEvent;
      dispatch(addNewMessage(data.data));
    },
    onError: error => {
      console.error(error);
    },
  });
  const dispatch = useAppDispatch();
  const {
    register,
    getFieldState,
    handleSubmit,
    reset,
    formState: { touchedFields, isValid, errors, isSubmitted },
  } = useForm<FormInput>({
    defaultValues: { content: '' },
    mode: 'onTouched', // default is onSubmit for validation to trigger
  });

  useEffect(() => {
    dispatch(
      addAppListener({
        actionCreator: setCurChannelId,
        effect: (action, listenerApi) => {
          const curChannelId = listenerApi.getState().chat.curChannelId;
          if (curChannelId !== curChannel.current?.channelId) {
            const channel = listenerApi.getState().chat.channels.find(row => row.channelId === curChannelId);
            if (channel) {
              curChannel.current = channel;
              displayContent(channel);
              setTimeout(() => {
                scrollToBottom();
              }, 1);
            }
          }
        },
      })
    );
    dispatch(
      addAppListener({
        actionCreator: addNewMessage,
        effect: (action, listenerApi) => {
          const curChannelId = listenerApi.getState().chat.curChannelId;
          if (curChannelId === curChannel.current?.channelId) {
            const channel = listenerApi.getState().chat.channels.find(row => row.channelId === curChannelId);
            if (channel) {
              displayContent(channel);
              setTimeout(() => {
                scrollToBottom();
              }, 1);
            }
          }
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
          updateOnlineStatus(listenerApi.getState().user.friends);
        },
      })
    );
  }, []);

  const displayContent = (channel: Channel) => {
    const temp = channel.messages.map(row => <Content props={row} key={row.messageId || row.createdAt} />);
    setContent(temp);
  };

  const scrollToBottom = () => {
    dialogueEndRef.current?.scrollIntoView();
  };

  const updateOnlineStatus = (friends: FriendHash) => {
    if (!curChannel.current || curChannel.current.userIds.length !== 2) {
      return;
    }

    const friendId = curChannel.current.userIds.filter(row => row !== user.userId)[0];
    if (friendId in friends) {
      const friend = friends[friendId];
      setIsOnline(friend.isOnline ? true : false);
    }
  };

  const onSubmit = async (data: FormInput) => {
    const message: Message = {
      messageId: 0,
      channelId: curChannel.current?.channelId as string,
      senderId: user.userId,
      messageType: 'string',
      content: data.content,
      createdAt: new Date().toISOString(),
      messageStatus: 0,
    };
    dispatch(addNewMessage(message));
    reset();
    sendJsonMessage(message);
  };

  const handleKeyDown = (e: KeyboardEvent<HTMLTextAreaElement>) => {
    // KeyUp is too late, newline will be created on Enter.
    if (e.key === 'Enter' && !e.shiftKey) {
      e.preventDefault();
      handleSubmit(onSubmit)();
      return;
    }
  };

  const handleOldMessages = () => {};

  return (
    <>
      {!curChannel.current && (
        <div className={styles.placeholder}>
          <img className={styles.logo} src={AppLogo}></img>
        </div>
      )}
      {curChannel.current && (
        <div className={styles.dialogueWrapper}>
          <div className={`${styles.header} p-3`}>
            <FontAwesomeIcon className={styles.icon} size="3x" icon={['fas', 'circle-user']} />
            <div className="ms-3">
              <div className={`${styles.heading}`}>{curChannel.current.channelName}</div>
              {isOnline && <div>Online</div>}
            </div>
          </div>
          <div className={`${styles.dialogue} p-5`}>
            {content}
            <div ref={dialogueEndRef}></div>
          </div>
          <div className={`${styles.footer} p-3`}>
            <form className={styles.formWrapper}>
              <textarea
                {...register('content')}
                id="message-text-area"
                wrap="hard"
                rows={1}
                autoComplete="on"
                className="base-input"
                placeholder="Type a message"
                onKeyDown={handleKeyDown}></textarea>
            </form>
            <button className="btn-icon ms-3" onClick={handleSubmit(onSubmit)}>
              <FontAwesomeIcon size="lg" icon={['fas', 'paper-plane']} />
            </button>
          </div>
        </div>
      )}
    </>
  );
}
