import { FontAwesomeIcon } from '@fortawesome/react-fontawesome';
import { useForm } from 'react-hook-form';
import { KeyboardEvent } from 'react';
import { SendJsonMessage } from 'react-use-websocket/dist/lib/types';

import { Channel, Message, MessageStatus } from 'features/chat/redux/chat.interface';
import { useAppDispatch, useAppSelector } from 'core/redux/reduxHooks';
import { addMessage } from 'features/chat/redux/chatSlice';
import styles from './dialogueFooter.module.scss';

interface FormInput {
  content: string;
}

interface DialogueFooterProps {
  channel: Channel;
  sendJsonMessage: SendJsonMessage;
  handleScrollDown: () => void;
}

export default function DialogueFooter({ channel, sendJsonMessage, handleScrollDown }: DialogueFooterProps) {
  const dispatch = useAppDispatch();
  const user = useAppSelector(state => state.user);
  const { register, handleSubmit, reset } = useForm<FormInput>({
    defaultValues: { content: '' },
    mode: 'onTouched', // default is onSubmit for validation to trigger
  });

  const onSubmit = async (data: FormInput) => {
    const timestamp = new Date().toISOString();
    const message: Message = {
      messageId: 0,
      channelId: channel.channelId,
      senderId: user.userId,
      messageType: 'string',
      content: data.content,
      messageStatus: MessageStatus.PENDING,
      createdAt: timestamp,
      updatedAt: timestamp,
    };
    dispatch(addMessage(message));
    sendJsonMessage(message);
    handleScrollDown();
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
    <div className={`${styles.footerWrapper} p-3`}>
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
  );
}
