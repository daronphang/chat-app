import { KeyboardEvent, useEffect, useState } from 'react';
import { FontAwesomeIcon } from '@fortawesome/react-fontawesome';

import { Channel } from 'features/chat/redux/chatSlice';
import { addAppListener } from 'shared/redux/listenerMiddleware';
import { useAppDispatch } from 'shared/redux/reduxHooks';
import styles from './channel.module.scss';
import AppLogo from 'assets/images/chatLogo.png';
import Content from '../content/content';
import { useForm } from 'react-hook-form';

interface FormInput {
  message: string;
}

export default function ChannelDialogue() {
  const [channel, setChannel] = useState<Channel | null>(null);
  const [content, setContent] = useState<JSX.Element[]>([]);
  const dispatch = useAppDispatch();
  const {
    register,
    getFieldState,
    handleSubmit,
    reset,
    formState: { touchedFields, isValid, errors, isSubmitted },
  } = useForm<FormInput>({
    defaultValues: { message: '' },
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
          const activeChannel = listenerApi.getState().chat.curChannel;
          if (activeChannel) {
            setChannel(activeChannel);
            displayContent(activeChannel);
          }
        },
      }),
      addAppListener({
        predicate: action => {
          if (action.type === 'chat/addNewMessage') {
            return true;
          }
          return false;
        },
        effect: (action, listenerApi) => {
          const activeChannel = listenerApi.getState().chat.curChannel;

          if (activeChannel) {
            displayContent(activeChannel);
          }

          setChannel(activeChannel);
        },
      })
    );
  }, []);

  const displayContent = (channel: Channel) => {
    const temp = channel.messages.map(row => <Content props={row} key={row.messageId} />);
    setContent(temp);
  };

  const onSubmit = (data: FormInput) => {
    console.log(data);
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
                {...register('message')}
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
