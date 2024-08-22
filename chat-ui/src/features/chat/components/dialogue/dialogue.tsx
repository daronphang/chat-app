import { useEffect, useRef, useState } from 'react';
import { SendJsonMessage } from 'react-use-websocket/dist/lib/types';
import { RpcError } from 'grpc-web';
import { enqueueSnackbar } from 'notistack';

import userPb from 'proto/user/user_pb';
import messagePb from 'proto/message/message_pb';
import AppLogo from 'assets/images/chatLogo2.png';
import { useAppDispatch, useAppSelector } from 'core/redux/reduxHooks';
import { addOlderMessages, setLastReadMessageInChannel } from 'features/chat/redux/chatSlice';
import { Channel, Message } from 'features/chat/redux/chat.interface';
import Content from '../content/content';
import styles from './dialogue.module.scss';
import { defaultSnackbarOptions } from 'core/config/snackbar.constant';
import DialogueHeader from '../dialogueHeader/dialogueHeader';
import DialogueFooter from '../dialogueFooter/dialogueFooter';
import { isGroupChat } from 'core/utils/chat';

interface DialogueProps {
  sendJsonMessage: SendJsonMessage;
}

export default function Dialogue({ sendJsonMessage }: DialogueProps) {
  const [content, setContent] = useState<JSX.Element[]>([]);
  const [isLoading, setIsLoading] = useState<boolean>(false);
  const [curChannel, setCurChannel] = useState<Channel | null>(null);
  const dialogueEndRef = useRef<null | HTMLDivElement>(null);
  const scrollContainerRef = useRef<null | HTMLDivElement>(null);
  const curScrollHeight = useRef<number>(-1);

  const user = useAppSelector(state => state.user);
  const config = useAppSelector(state => state.config);
  const curChannelId = useAppSelector(state => state.chat.curChannelId);
  const channels = useAppSelector(state => state.chat.channels);
  const dispatch = useAppDispatch();

  useEffect(() => {
    if (!curChannelId) {
      setCurChannel(null);
      return;
    }

    const latestChannel = channels.find(row => row.channelId === curChannelId);
    if (!latestChannel) {
      setCurChannel(null);
      return;
    } else if (latestChannel === curChannel) {
      // If the channel is updated in Redux, it will have a different object reference.
      return;
    }

    // If channel has changed, to reset scrollHeight.
    if (curChannel?.channelId !== curChannelId) {
      curScrollHeight.current = -1;
    }

    handleMessageUpdates(latestChannel);
    setCurChannel(latestChannel);
  }, [curChannelId, channels]);

  const handleMessageUpdates = async (latestChannel: Channel) => {
    displayContent(latestChannel, isGroupChat(latestChannel));

    // Update lastMessageId if required.
    if (latestChannel.messages.length > 0) {
      const latestMessageId = latestChannel.messages[latestChannel.messages.length - 1].messageId;
      if (latestMessageId > latestChannel.lastMessageId) {
        await updateLastMessageId(latestChannel.channelId, latestMessageId);
        dispatch(setLastReadMessageInChannel(latestChannel.channelId));
      }
    }

    // Manage scroll position.
    if (curScrollHeight.current === -1) {
      setTimeout(() => {
        scrollToBottom();
      }, 1);
    } else if (scrollContainerRef.current?.scrollHeight === 0) {
      setTimeout(() => {
        maintainScrollPosition();
      }, 1);
    }
  };

  const displayContent = (channel: Channel, isGroup: boolean) => {
    const temp = channel.messages.map(row => (
      <Content isGroup={isGroup} message={row} key={row.messageId || new Date(row.createdAt).getTime()} />
    ));
    setContent(temp);
  };

  const scrollToBottom = () => {
    dialogueEndRef.current?.scrollIntoView();
  };

  const handleScrollDown = () => {
    curScrollHeight.current = -1;
  };

  const maintainScrollPosition = () => {
    if (scrollContainerRef.current) {
      const anchor = scrollContainerRef.current.scrollHeight - curScrollHeight.current;
      scrollContainerRef.current.scrollTo(0, anchor);
    }
  };

  const updateLastMessageId = async (channelId: string, latestMessageId: number) => {
    try {
      const payload = new userPb.LastReadMessage();
      payload.setUserid(user.userId);
      payload.setChannelid(channelId);
      payload.setLastmessageid(latestMessageId);
      await config.api.USER_SERVICE.updateLastReadMessage(payload);
    } catch (e) {
      const err = e as RpcError;
      console.error('failed to update last messsage id', err.message);
    }
  };

  const handleScroll = async (event: React.UIEvent<HTMLDivElement>) => {
    const target = event.target as HTMLDivElement;
    const isScrolledToBottom = Math.abs(target.scrollHeight - target.scrollTop - target.clientHeight) < 1;
    if (isScrolledToBottom) {
      curScrollHeight.current = -1;
    } else {
      curScrollHeight.current = target.scrollHeight;
    }
    if (target.scrollTop === 0) {
      setIsLoading(true);
      const resp = await retrieveOlderMessages();
      if (resp.length > 0) {
        dispatch(addOlderMessages(resp));
      }
      setIsLoading(false);
    }
  };

  const retrieveOlderMessages = async (): Promise<Message[]> => {
    try {
      if (!curChannel || curChannel.messages.length === 0) {
        return new Promise(resolve => resolve([]));
      }
      const payload = new messagePb.MessageRequest();
      payload.setChannelid(curChannel.channelId);
      const lastMessageId = curChannel.messages[0].messageId;
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
      {!curChannel && (
        <div className={styles.placeholder}>
          <img className={styles.logo} src={AppLogo}></img>
        </div>
      )}
      {curChannel && (
        <div className={styles.dialogueWrapper}>
          <DialogueHeader channel={curChannel} />
          <div ref={scrollContainerRef} onScroll={handleScroll} className={`${styles.dialogue} p-5`}>
            {isLoading && <div className={`${styles.loader} p-2 mb-3`}>Loading older messages...</div>}
            {content}
            <div ref={dialogueEndRef}></div>
          </div>
          <DialogueFooter channel={curChannel} sendJsonMessage={sendJsonMessage} handleScrollDown={handleScrollDown} />
        </div>
      )}
    </>
  );
}
