import { Chip, Tooltip } from '@mui/material';
import { FontAwesomeIcon } from '@fortawesome/react-fontawesome';
import { useEffect, useState } from 'react';
import { enqueueSnackbar } from 'notistack';

import Search from 'shared/components/search/search';
import { Friend } from 'features/user/redux/user.interface';
import { useAppDispatch, useAppSelector } from 'core/redux/reduxHooks';
import { defaultSnackbarOptions } from 'core/config/snackbar.constant';
import { Channel } from 'features/chat/redux/chat.interface';
import { addChannel } from 'features/chat/redux/chatSlice';
import styles from './newGroup.module.scss';
import Drawer from '../drawer/drawer';
import { useForm } from 'react-hook-form';
interface NewGroupProps {
  handleClickBack: () => void;
  createNewChannel: (arg: Channel) => Promise<Channel | null>;
  broadcastChannelEvent: (arg: Channel) => void;
}

interface FormInput {
  groupName: string;
  users: Friend[];
}

export default function NewGroup({ handleClickBack, createNewChannel, broadcastChannelEvent }: NewGroupProps) {
  const [drawers, setDrawers] = useState<JSX.Element[]>([]);
  const [users, setUsers] = useState<Friend[]>([]);
  const config = useAppSelector(state => state.config);
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
    const friends: Friend[] = Object.values(user.friends);
    displayFriends(friends);
  }, []);

  const handleSearchResult = (v: any[]) => {
    displayFriends(v);
  };

  const handleClickDrawer = (v: Friend) => {
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

  const displayFriends = (friends: Friend[]) => {
    friends.sort((a, b) => a.displayName.localeCompare(b.displayName));
    setDrawers(
      friends.map(row => (
        <Drawer key={row.userId} data={row} title={row.displayName} text="" handleClickDrawer={handleClickDrawer} />
      ))
    );
  };

  const onSubmit = async (data: FormInput) => {
    const userIds = data.users.map(row => row.userId);
    userIds.push(user.userId);

    const newChannel: Channel = {
      channelId: '',
      channelName: data.groupName,
      userIds,
      messages: [],
      createdAt: new Date().toISOString(),
      updatedAt: new Date().toISOString(),
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
    <div>
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
        <input {...register('users', { validate: isValidUsersInput })} id="new-group-users-input" hidden={true}></input>
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
      <Search
        sourceData={Object.values(user.friends)}
        excludedFields={['userId', 'email', 'isOnline']}
        handleSearchResult={handleSearchResult}
      />
      <div className="mb-3"></div>
      {drawers}
    </div>
  );
}
