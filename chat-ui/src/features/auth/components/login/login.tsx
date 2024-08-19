import { enqueueSnackbar } from 'notistack';
import { useForm } from 'react-hook-form';
import { RpcError } from 'grpc-web';
import { useNavigate } from 'react-router-dom';

import { useAppDispatch, useAppSelector } from 'core/redux/reduxHooks';
import { defaultSnackbarOptions } from 'core/config/snackbar.constant';
import userPb from 'proto/user/user_pb';
import { Recipient, UserMetadata } from 'features/user/redux/user.interface';
import { setUser } from 'features/user/redux/userSlice';
import { RoutePath } from 'core/config/route.constant';
import styles from './login.module.scss';
import { getRandomColor } from 'core/utils/formatters';

interface FormInput {
  email: string;
}

interface LoginProps {
  showSignup: () => void;
}

export default function Login({ showSignup }: LoginProps) {
  const config = useAppSelector(state => state.config);
  const dispatch = useAppDispatch();
  const navigate = useNavigate();
  const {
    register,
    handleSubmit,
    formState: { errors },
  } = useForm<FormInput>({
    defaultValues: { email: '' },
    mode: 'onTouched', // default is onSubmit for validation to trigger
  });

  const onError = () => {
    enqueueSnackbar('Missing or invalid fields', {
      ...defaultSnackbarOptions,
      variant: 'error',
    });
  };

  const onSubmit = async (data: FormInput) => {
    const resp = await login(data);
    if (!resp) return;

    const user: UserMetadata = {
      userId: resp.getUserid(),
      email: resp.getEmail(),
      displayName: resp.getDisplayname(),
      recipients: {},
    };

    resp.getFriendsList().forEach(row => {
      const friend: Recipient = {
        userId: row.getUserid(),
        email: row.getEmail(),
        displayName: row.getDisplayname(),
        isOnline: false,
        friendName: row.getFriendname(),
        isFriend: true,
        color: getRandomColor(),
      };
      user.recipients[friend.userId] = friend;
    });

    dispatch(setUser(user));
    navigate(RoutePath.CHAT);
  };

  const login = async (data: FormInput): Promise<userPb.UserMetadata | null> => {
    try {
      const payload = new userPb.UserCredentials();
      payload.setEmail(data.email);
      const resp = await config.api.USER_SERVICE.login(payload);
      return new Promise(resolve => resolve(resp));
    } catch (e) {
      const err = e as RpcError;
      const errMsg = err.code === 14 ? config.apiError.NETWORK_ERROR : 'Invalid credentials';
      enqueueSnackbar(errMsg, {
        ...defaultSnackbarOptions,
        variant: 'error',
      });
      return new Promise(resolve => resolve(null));
    }
  };

  return (
    <div className={`${styles.loginWrapper} p-4`}>
      <h2 className="mt-3">Login</h2>
      <div className="mt-5 w-100">
        <form>
          <input
            {...register('email', {
              required: true,
              pattern: /^[a-zA-Z0-9_.$!%#&*+-]+@[a-zA-Z0-9-]+\.[a-zA-Z0-9-.]+$/,
            })}
            id="login-email-input"
            autoComplete="on"
            placeholder="Enter Email"
            className={`base-input ${styles.inputField}`}></input>
          {errors.email && <span className="input-error-msg">Email is invalid</span>}
        </form>
      </div>
      <button className={`btn mt-4 ${styles.button}`} onClick={handleSubmit(onSubmit, onError)}>
        Sign In
      </button>
      <button onClick={showSignup} className={`${styles.footer} mt-5 btn`}>
        Don't have an account? Register
      </button>
    </div>
  );
}
