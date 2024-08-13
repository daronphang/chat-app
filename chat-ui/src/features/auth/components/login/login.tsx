import { useEffect, useState } from 'react';
import { useNavigate } from 'react-router-dom';
import { useAppDispatch, useAppSelector } from 'core/redux/reduxHooks';
import { RpcError } from 'grpc-web';
import { useSnackbar } from 'notistack';

import userPb from 'proto/user/user_pb';
import { setUser } from 'features/user/redux/userSlice';
import { Friend, FriendHash, UserMetadata } from 'features/user/redux/user.interface';
import { RoutePath } from 'core/config/route.constant';
import { defaultSnackbarOptions } from 'core/config/snackbar.constant';
import styles from './login.module.scss';

interface FormInput {
  email: string;
}

export default function Login() {
  const [loading, isLoading] = useState<boolean>(false);
  const { enqueueSnackbar } = useSnackbar();
  const navigate = useNavigate();
  const config = useAppSelector(state => state.config);
  const dispatch = useAppDispatch();

  useEffect(() => {
    onSubmit({ email: 'daronphang@gmail.com' });
  }, []);

  const onSubmit = async (data: FormInput) => {
    const resp = await login(data);
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
    <>
      <div>Login page</div>
    </>
  );
}
