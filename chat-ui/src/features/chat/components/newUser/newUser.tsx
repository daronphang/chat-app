import { useAppDispatch, useAppSelector } from 'core/redux/reduxHooks';
import { useForm } from 'react-hook-form';
import { enqueueSnackbar } from 'notistack';
import { RpcError } from 'grpc-web';

import { NewFriend } from 'proto/user/user_pb';
import { Friend } from 'features/user/redux/user.interface';
import { addFriend } from 'features/user/redux/userSlice';
import { defaultSnackbarOptions } from 'core/config/snackbar.constant';
import styles from './newUser.module.scss';
import { Tooltip } from '@mui/material';
import { FontAwesomeIcon } from '@fortawesome/react-fontawesome';

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
    getFieldState,
    handleSubmit,
    reset,
    formState: { touchedFields, isValid, errors, isSubmitted },
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
    handleClickBack();
  };

  const addNewFriend = async (newFriend: Friend): Promise<boolean> => {
    try {
      const payload = new NewFriend();
      payload.setFriendemail(newFriend.email);
      payload.setDisplayname(newFriend.displayName);
      const resp = await config.api.USER_SERVICE.addFriend(payload);

      // Update required metadata.
      newFriend.userId = resp.getUserid();

      dispatch(addFriend(newFriend));
      enqueueSnackbar('New friend added', {
        ...defaultSnackbarOptions,
        variant: 'success',
      });
      return new Promise(resolve => resolve(true));
    } catch (e) {
      const err = e as RpcError;
      if (err.code === 14) {
        enqueueSnackbar(config.apiError.NETWORK_ERROR, {
          ...defaultSnackbarOptions,
          variant: 'error',
        });
      } else {
        enqueueSnackbar('Failed to add friend, please try again', {
          ...defaultSnackbarOptions,
          variant: 'error',
        });
        console.error(err.message);
      }
      return new Promise(resolve => resolve(false));
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
          {...register('email')}
          id="new-friend-email-input"
          autoComplete="on"
          placeholder="Friend Email"
          className={`base-input ${styles.inputField}`}></input>
        <input
          {...register('displayName')}
          id="new-friend-display-name-input"
          autoComplete="on"
          placeholder="Friend Display Name"
          className={`base-input mt-3 ${styles.inputField}`}></input>
      </form>
      <button className={`btn mt-4 ${styles.button}`} onClick={handleSubmit(onSubmit)}>
        Add Friend
      </button>
    </div>
  );
}
