import { enqueueSnackbar } from 'notistack';
import { useForm } from 'react-hook-form';
import { RpcError } from 'grpc-web';
import { useNavigate } from 'react-router-dom';

import { useAppDispatch, useAppSelector } from 'core/redux/reduxHooks';
import { defaultSnackbarOptions } from 'core/config/snackbar.constant';
import userPb from 'proto/user/user_pb';
import { Friend, UserMetadata } from 'features/user/redux/user.interface';
import { setUser } from 'features/user/redux/userSlice';
import { RoutePath } from 'core/config/route.constant';
import styles from './signup.module.scss';

interface FormInput {
  email: string;
  displayName: string;
}

interface SignupProps {
  showLogin: () => void;
}

export default function Signup({ showLogin }: SignupProps) {
  const config = useAppSelector(state => state.config);
  const dispatch = useAppDispatch();
  const navigate = useNavigate();
  const {
    register,
    handleSubmit,
    reset,
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
    const resp = await signup(data);
    if (!resp) return;

    const user: UserMetadata = {
      userId: resp.getUserid(),
      email: resp.getEmail(),
      displayName: resp.getDisplayname(),
      friends: {},
    };

    resp.getFriendsList().forEach(row => {
      const friend: Friend = {
        userId: row.getUserid(),
        email: row.getEmail(),
        displayName: row.getDisplayname(),
        isOnline: false,
      };
      user.friends[friend.userId] = friend;
    });

    dispatch(setUser(user));
    navigate(RoutePath.CHAT);
  };

  const signup = async (data: FormInput): Promise<userPb.UserMetadata | null> => {
    try {
      const payload = new userPb.NewUser();
      payload.setEmail(data.email);
      payload.setDisplayname(data.displayName);
      const resp = await config.api.USER_SERVICE.signup(payload);
      return new Promise(resolve => resolve(resp));
    } catch (e) {
      const err = e as RpcError;
      const errMsg = err.code === 14 ? config.apiError.NETWORK_ERROR : 'Failed to register';
      enqueueSnackbar(errMsg, {
        ...defaultSnackbarOptions,
        variant: 'error',
      });
      return new Promise(resolve => resolve(null));
    }
  };

  return (
    <div className={`${styles.signupWrapper} p-4`}>
      <h2 className="mt-3">Register</h2>
      <div className="mt-5 w-100">
        <form>
          <input
            {...register('email', {
              required: true,
              pattern: /^[a-zA-Z0-9_.$!%#&*+-]+@[a-zA-Z0-9-]+\.[a-zA-Z0-9-.]+$/,
            })}
            id="signup-email-input"
            autoComplete="on"
            placeholder="Enter Email"
            className={`base-input ${styles.inputField}`}></input>
          {errors.email && <span className="input-error-msg">Email is invalid</span>}
          <input
            {...register('displayName', {
              required: true,
            })}
            id="signup-display-name-input"
            autoComplete="on"
            placeholder="Enter Display Name"
            className={`base-input ${styles.inputField} mt-3`}></input>
          {errors.displayName && <span className="input-error-msg">Field is required</span>}
        </form>
      </div>
      <button className={`btn mt-4 ${styles.button}`} onClick={handleSubmit(onSubmit, onError)}>
        Register
      </button>
      <button onClick={showLogin} className={`${styles.footer} mt-5 btn`}>
        Have an account? Login
      </button>
    </div>
  );
}
