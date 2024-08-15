import { KeyboardEvent, UIEventHandler, useEffect, useRef, useState } from 'react';
import { FontAwesomeIcon } from '@fortawesome/react-fontawesome';
import { useForm } from 'react-hook-form';
import { SendJsonMessage } from 'react-use-websocket/dist/lib/types';
import { RpcError } from 'grpc-web';

import userPb from 'proto/user/user_pb';
import messagePb from 'proto/message/message_pb';
import AppLogo from 'assets/images/chatLogo.png';
import { addAppListener } from 'core/redux/listenerMiddleware';
import { useAppDispatch, useAppSelector } from 'core/redux/reduxHooks';
import { FriendHash } from 'features/user/redux/user.interface';
import { addMessage, addOlderMessages, setCurChannelId, setLastReadMessageId } from 'features/chat/redux/chatSlice';
import { Channel, Message, MessageStatus } from 'features/chat/redux/chat.interface';
import Content from '../content/content';
import styles from './dialogue.module.scss';
import { enqueueSnackbar } from 'notistack';
import { defaultSnackbarOptions } from 'core/config/snackbar.constant';

interface FormInput {
  content: string;
}

interface DialogueProps {
  sendJsonMessage: SendJsonMessage;
}

export default function Dialogue({ sendJsonMessage }: DialogueProps) {
  const [content, setContent] = useState<JSX.Element[]>([]);
  const [isOnline, setIsOnline] = useState<boolean>(false);
  const [isLoading, setIsLoading] = useState<boolean>(false);
  const curChannel = useRef<Channel | null>(null);
  const dialogueEndRef = useRef<null | HTMLDivElement>(null);
  const scrollContainerRef = useRef<null | HTMLDivElement>(null);
  const prevScrollHeight = useRef<number>(0);

  const user = useAppSelector(state => state.user);
  const config = useAppSelector(state => state.config);
  const dispatch = useAppDispatch();
  const { register, handleSubmit, reset } = useForm<FormInput>({
    defaultValues: { content: '' },
    mode: 'onTouched', // default is onSubmit for validation to trigger
  });

  useEffect(() => {
    dispatch(
      addAppListener({
        actionCreator: setCurChannelId,
        effect: (action, listenerApi) => {
          const curChannelId = listenerApi.getState().chat.curChannelId;
          if (curChannelId === curChannel.current?.channelId) {
            return;
          }
          // Guaranteed to exist.
          const channel = listenerApi.getState().chat.channelHash[curChannelId];
          curChannel.current = channel;
          displayContent(channel);
          setTimeout(() => {
            scrollToBottom();
          }, 1);
        },
      })
    );
    dispatch(
      addAppListener({
        predicate: action => {
          if (['chat/addMessage', 'chat/addOlderMessages'].includes(action.type)) return true;
          return false;
        },
        effect: (action, listenerApi) => {
          if (!curChannel.current) {
            return;
          }
          // If the channel is updated in Redux, it will have a different object reference.
          const channel = listenerApi.getState().chat.channelHash[curChannel.current.channelId];
          if (channel === curChannel.current) {
            return;
          }
          curChannel.current = channel;
          displayContent(channel);

          if (action.type === 'chat/addMessage') {
            setTimeout(() => {
              scrollToBottom();
            }, 1);
          } else {
            setTimeout(() => {
              maintainScrollPosition();
            }, 1);
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

  useEffect(() => {
    (async () => {
      if (!curChannel.current || content.length === 0) {
        return;
      }
      const latestMessageId = curChannel.current.messages[curChannel.current.messages.length - 1].messageId;
      if (latestMessageId > curChannel.current.lastMessageId) {
        await updateLastMessageId(latestMessageId);
        dispatch(setLastReadMessageId(curChannel.current.channelId));
      }
    })();
  }, [content]);

  const displayContent = (channel: Channel) => {
    const temp = channel.messages.map(row => <Content props={row} key={row.messageId || row.createdAt} />);
    setContent(temp);
  };

  const scrollToBottom = () => {
    dialogueEndRef.current?.scrollIntoView();
  };

  const maintainScrollPosition = () => {
    if (scrollContainerRef.current) {
      const anchor = scrollContainerRef.current.scrollHeight - prevScrollHeight.current;
      scrollContainerRef.current.scrollTo(0, anchor);
    }
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
    const timestamp = new Date().toISOString();
    const message: Message = {
      messageId: 0,
      channelId: curChannel.current?.channelId as string,
      senderId: user.userId,
      messageType: 'string',
      content: data.content,
      messageStatus: MessageStatus.PENDING,
      createdAt: timestamp,
      updatedAt: timestamp,
    };
    dispatch(addMessage(message));
    sendJsonMessage(message);
    reset();
  };

  const handleKeyDown = (e: KeyboardEvent<HTMLTextAreaElement>) => {
    // KeyUp is too late, newline will be created on Enter.
    if (e.key === 'Enter' && !e.shiftKey) {
      e.preventDefault();
      handleSubmit(onSubmit)();
      return;
    }
  };

  const updateLastMessageId = async (latestMessageId: number) => {
    try {
      if (!curChannel.current) return;
      const payload = new userPb.LastReadMessage();
      payload.setUserid(user.userId);
      payload.setChannelid(curChannel.current.channelId);
      payload.setLastmessageid(latestMessageId);
      await config.api.USER_SERVICE.updateLastReadMessage(payload);
    } catch (e) {
      const err = e as RpcError;
      console.error('failed to update last messsage id', err.message);
    }
  };

  const handleScroll = async (event: React.UIEvent<HTMLDivElement>) => {
    const target = event.target as HTMLDivElement;
    if (target.scrollTop === 0) {
      prevScrollHeight.current = target.scrollHeight;
      setIsLoading(true);
      const resp = await retrieveOlderMessages();
      dispatch(addOlderMessages(resp));
      setIsLoading(false);
    }
  };

  const retrieveOlderMessages = async (): Promise<Message[]> => {
    try {
      if (!curChannel.current) {
        return new Promise(resolve => resolve([]));
      }
      const payload = new messagePb.MessageRequest();
      payload.setChannelid(curChannel.current.channelId);
      const lastMessageId = curChannel.current.messages[0].messageId;
      payload.setLastmessageid(lastMessageId);
      const resp = await config.api.MESSAGE_SERVICE.getPreviousMessages(payload);

      const messages: Message[] = resp.getMessagesList().map(row => {
        return {
          messageId: row.getMessageid(),
          channelId: row.getChannelid(),
          senderId: row.getSenderid(),
          messageType: row.getMessagetype(),
          content: row.getContent(),
          createdAt: row.getCreatedat(),
          messageStatus: row.getMessagestatus(),
          updatedAt: new Date().toISOString(),
        };
      });
      return new Promise(resolve => resolve(messages));
    } catch (e) {
      const err = e as RpcError;
      const errMsg = err.code === 14 ? config.apiError.NETWORK_ERROR : 'Failed to retrieve older messages';
      enqueueSnackbar(errMsg, {
        ...defaultSnackbarOptions,
        variant: 'error',
      });
      return new Promise(resolve => resolve([]));
    }
  };

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
          <div ref={scrollContainerRef} onScroll={handleScroll} className={`${styles.dialogue} p-5`}>
            {isLoading && <div className={`${styles.loader} p-2 mb-3`}>Loading older messages...</div>}
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
