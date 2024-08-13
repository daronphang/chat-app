import { useAppDispatch, useAppSelector } from 'core/redux/reduxHooks';
import { useForm } from 'react-hook-form';
import { enqueueSnackbar } from 'notistack';
import { RpcError } from 'grpc-web';
import { Tooltip } from '@mui/material';
import { FontAwesomeIcon } from '@fortawesome/react-fontawesome';

import userPb from 'proto/user/user_pb';
import { Friend } from 'features/user/redux/user.interface';
import { addFriend } from 'features/user/redux/userSlice';
import { defaultSnackbarOptions } from 'core/config/snackbar.constant';
import styles from './newUser.module.scss';

interface FormInput {
  email: string;
  displayName: string;
}

interface NewFriendProps {
  handleClickBack: () => void;
}

export default function NewUser({ handleClickBack }: NewFriendProps) {
  const dispatch = useAppDispatch();
  const config = useAppSelector(state => state.config);
  const user = useAppSelector(state => state.user);
  const {
    register,
    handleSubmit,
    reset,
    formState: { errors },
  } = useForm<FormInput>({
    defaultValues: { email: '', displayName: '' },
    mode: 'onTouched', // default is onSubmit for validation to trigger
  });

  const onSubmit = async (data: FormInput) => {
    // Check if friend exists.
    const friends = Object.values(user.friends);
    const exist = friends.find(row => row.email === data.email);
    if (exist) {
      enqueueSnackbar('Friend already exists', {
        ...defaultSnackbarOptions,
        variant: 'error',
      });
      return;
    }

    // Add friend.
    const friend: Friend = {
      userId: '',
      email: data.email,
      displayName: data.displayName,
      isOnline: false,
    };
    const resp = await addNewFriend(friend);
    if (!resp) return;

    // Update required metadata.
    friend.userId = resp.getUserid();
    dispatch(addFriend(friend));
    reset();
  };

  const onError = () => {
    enqueueSnackbar('Missing or invalid fields', {
      ...defaultSnackbarOptions,
      variant: 'error',
    });
  };

  const addNewFriend = async (newFriend: Friend): Promise<userPb.Friend | null> => {
    try {
      const payload = new userPb.NewFriend();
      payload.setUserid(user.userId);
      payload.setFriendemail(newFriend.email);
      payload.setDisplayname(newFriend.displayName);
      const resp = await config.api.USER_SERVICE.addFriend(payload);

      enqueueSnackbar('New friend added', {
        ...defaultSnackbarOptions,
        variant: 'success',
      });
      return new Promise(resolve => resolve(resp));
    } catch (e) {
      const err = e as RpcError;
      const errMsg = err.code === 14 ? config.apiError.NETWORK_ERROR : 'Failed to add friend';
      enqueueSnackbar(errMsg, {
        ...defaultSnackbarOptions,
        variant: 'error',
      });
      return new Promise(resolve => resolve(null));
    }
  };

  return (
    <div>
      <div className={styles.headerWrapper}>
        <Tooltip title="Back" placement="bottom">
          <button className="btn-icon ms-3" onClick={() => handleClickBack()}>
            <FontAwesomeIcon size="lg" icon={['fas', 'arrow-left']} />
          </button>
        </Tooltip>
        <div className={`${styles.heading} ms-3`}>New Friend</div>
      </div>
      <div className="mb-4"></div>
      <form className={styles.formWrapper}>
        <input
          {...register('email', { required: true, pattern: /^[a-zA-Z0-9_.$!%#&*+-]+@[a-zA-Z0-9-]+\.[a-zA-Z0-9-.]+$/ })}
          id="new-friend-email-input"
          autoComplete="on"
          placeholder="Friend Email"
          className={`base-input ${styles.inputField}`}></input>
        {errors.email && <span className="input-error-msg">Email is invalid</span>}
        <input
          {...register('displayName', { required: true })}
          id="new-friend-display-name-input"
          autoComplete="on"
          placeholder="Friend Name"
          className={`base-input mt-3 ${styles.inputField}`}></input>
        {errors.displayName && <span className="input-error-msg">Field is required</span>}
      </form>
      <button className={`btn mt-4 ${styles.button}`} onClick={handleSubmit(onSubmit, onError)}>
        Add Friend
      </button>
    </div>
  );
}
