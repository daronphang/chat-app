import { KeyboardEvent, useEffect, useState } from 'react';
import { FontAwesomeIcon } from '@fortawesome/react-fontawesome';

import { Channel, Message, WebSocketEvent, addNewMessage } from 'features/chat/redux/chatSlice';
import { addAppListener } from 'core/redux/listenerMiddleware';
import { useAppDispatch, useAppSelector } from 'core/redux/reduxHooks';
import styles from './channel.module.scss';
import AppLogo from 'assets/images/chatLogo.png';
import Content from '../content/content';
import { useForm } from 'react-hook-form';
import useWebSocket from 'react-use-websocket';
import { defaultWsOptions } from 'core/config/ws.constant';
interface FormInput {
  content: string;
}

export default function ChannelDialogue() {
  const [channel, setChannel] = useState<Channel | null>(null);
  const [content, setContent] = useState<JSX.Element[]>([]);
  const user = useAppSelector(state => state.user);
  const wsUrl = useAppSelector(state => state.config.wsUrl);
  const { sendJsonMessage } = useWebSocket(wsUrl, {
    ...defaultWsOptions,
    filter: event => {
      console.log(event.data);
      const temp = JSON.parse(event.data) as WebSocketEvent;
      if (temp.event === 'message') {
        return true;
      }
      return false;
    },
    onMessage: event => {
      const temp = JSON.parse(event.data) as WebSocketEvent;
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
    // Add redux listeners.
    dispatch(
      addAppListener({
        predicate: action => {
          if (action.type === 'chat/setCurChannel') {
            return true;
          }
          return false;
        },
        effect: (action, listenerApi) => {
          const curChannel = listenerApi.getState().chat.curChannel;
          if (curChannel && curChannel.channelId !== channel?.channelId) {
            setChannel(curChannel);
            displayContent(curChannel);
          }
        },
      })
    );
  }, []);

  const displayContent = (channel: Channel) => {
    const temp = channel.messages.map(row => <Content props={row} key={row.messageId} />);
    setContent(temp);
  };

  const onSubmit = (data: FormInput) => {
    const message: Message = {
      messageId: 0,
      channelId: channel?.channelId as string,
      senderId: user.userId,
      messageType: 'string',
      content: data.content,
      createdAt: new Date().toISOString(),
      messageStatus: 'pending',
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

  const handleNewMessage = (msg: Message) => {
    if (msg.channelId === channel?.channelId) {
      setContent(v => {
        v.push(<Content props={msg} key={msg.messageId} />);
        return v;
      });
    }
    dispatch(addNewMessage(msg));
  };

  const handleOldMessages = () => {};

  return (
    <>
      {!channel && (
        <div className={styles.placeholder}>
          <img className={styles.logo} src={AppLogo}></img>
        </div>
      )}
      {channel && (
        <div className={styles.channelWrapper}>
          <div className={`${styles.header} p-3`}>
            <FontAwesomeIcon size="2x" icon={['fas', 'circle-user']} />
            <div className={`${styles.heading} ms-3`}>{channel.channelName}</div>
          </div>
          <div className={`${styles.channel} p-5`}>{content}</div>
          <div className={`${styles.footer} p-3`}>
            {/* <div className="flex-spacer"></div> */}
            <form className={styles.formWrapper}>
              <textarea
                {...register('content')}
                id="message-text-area"
                wrap="hard"
                rows={1}
                autoComplete="on"
                className="input-field"
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
