import { Chip, Tooltip } from '@mui/material';
import { FontAwesomeIcon } from '@fortawesome/react-fontawesome';
import { useEffect, useRef, useState } from 'react';
import { enqueueSnackbar } from 'notistack';

import Search from 'shared/components/search/search';
import { Recipient } from 'features/user/redux/user.interface';
import { useAppDispatch, useAppSelector } from 'core/redux/reduxHooks';
import { defaultSnackbarOptions } from 'core/constants/snackbar.constant';
import { Channel } from 'features/chat/redux/chat.interface';
import { addChannel } from 'features/chat/redux/chatSlice';
import styles from './newGroup.module.scss';
import UserDrawer from '../userDrawer/userDrawer';
import { useForm } from 'react-hook-form';
interface NewGroupProps {
  handleClickBack: () => void;
  createNewChannel: (arg: Channel) => Promise<Channel | null>;
  broadcastChannelEvent: (arg: Channel) => void;
}

interface FormInput {
  groupName: string;
  users: Recipient[];
}

export default function NewGroup({ handleClickBack, createNewChannel, broadcastChannelEvent }: NewGroupProps) {
  const [drawers, setDrawers] = useState<JSX.Element[]>([]);
  const [users, setUsers] = useState<Recipient[]>([]);
  const friends = useRef<Recipient[]>([]);
  const user = useAppSelector(state => state.user);
  const dispatch = useAppDispatch();
  const {
    register,
    handleSubmit,
    setValue,
    reset,
    formState: { errors },
  } = useForm<FormInput>({
    defaultValues: { groupName: '', users: [] },
    mode: 'onTouched', // default is onSubmit for validation to trigger
  });

  useEffect(() => {
    friends.current = Object.values(user.recipients).filter(row => row.isFriend);
    displayFriends(friends.current);
  }, []);

  const handleSearchResult = (v: any[]) => {
    displayFriends(v);
  };

  const handleClickDrawer = (v: Recipient) => {
    setUsers(state => {
      const idx = state.findIndex(row => row.userId === v.userId);
      if (idx === -1) {
        state.push(v);
      }
      setValue('users', state, { shouldDirty: true, shouldTouch: true, shouldValidate: true });
      return [...state];
    });
  };

  const handleDeleteChip = (v: string) => {
    setUsers(state => {
      const idx = state.findIndex(row => row.userId === v);
      if (idx !== -1) {
        state.splice(idx, 1);
      }
      setValue('users', state, { shouldDirty: true, shouldTouch: true, shouldValidate: true });
      return [...state];
    });
  };

  const displayFriends = (friends: Recipient[]) => {
    friends.sort((a, b) => a.friendName.localeCompare(b.friendName));
    setDrawers(
      friends.map(row => (
        <UserDrawer key={row.userId} data={row} title={row.friendName} text="" handleClickDrawer={handleClickDrawer} />
      ))
    );
  };

  const onSubmit = async (data: FormInput) => {
    const userIds = data.users.map(row => row.userId);
    userIds.push(user.userId);
    const timestamp = new Date().toISOString();

    const newChannel: Channel = {
      channelId: '',
      channelName: data.groupName,
      userIds,
      messages: [],
      createdAt: timestamp,
      updatedAt: timestamp,
      lastMessageId: 0,
    };
    const resp = await createNewChannel(newChannel);
    if (!resp) return;

    await broadcastChannelEvent(resp);
    dispatch(addChannel(resp));
    reset();
  };

  const onError = () => {
    enqueueSnackbar('Missing or invalid fields', {
      ...defaultSnackbarOptions,
      variant: 'error',
    });
  };

  const isValidUsersInput = (): boolean => {
    if (users.length < 2) return false;
    return true;
  };

  return (
    <>
      <div className="p-3">
        <div className={styles.headerWrapper}>
          <Tooltip title="Back" placement="bottom">
            <button className="btn-icon ms-3" onClick={() => handleClickBack()}>
              <FontAwesomeIcon size="lg" icon={['fas', 'arrow-left']} />
            </button>
          </Tooltip>
          <div className={`${styles.heading} ms-3`}>New Group</div>
        </div>
        <div className="mb-4"></div>
        <form>
          <input
            {...register('groupName', { required: true })}
            id="new-group-group-name-input"
            autoComplete="on"
            placeholder="Group Name"
            className={`base-input mt-3 ${styles.inputField}`}></input>
          {errors.groupName && <span className="input-error-msg">Field is required</span>}
          <input
            {...register('users', { validate: isValidUsersInput })}
            id="new-group-users-input"
            hidden={true}></input>
          <div className={`${styles.chipWrapper} base-input mt-3`}>
            {users.map(row => (
              <Chip
                style={{ borderRadius: '0.3rem' }}
                variant="outlined"
                key={row.userId}
                id={row.userId}
                label={row.displayName}
                onDelete={() => handleDeleteChip(row.userId)}
              />
            ))}
            {users.length === 0 ? <span className={styles.placeholder}>Select users</span> : null}
          </div>
          {errors.users && <div className="input-error-msg">Minimum of 2 users</div>}
          <button className={`btn mt-4 ${styles.button}`} onClick={handleSubmit(onSubmit, onError)}>
            Create Group
          </button>
        </form>
        <div className="mb-4"></div>
        {friends.current.length > 0 && (
          <Search
            sourceData={friends.current}
            excludedFields={['userId', 'email', 'isOnline', 'isFriend', 'displayName']}
            handleSearchResult={handleSearchResult}
          />
        )}
        <div className="mb-3"></div>
      </div>
      {friends.current.length === 0 && <div className="text-center">To create a group, add friends.</div>}
      {friends.current.length > 0 && <div className={`${styles.drawerWrapper} p-3`}>{drawers}</div>}
    </>
  );
}
