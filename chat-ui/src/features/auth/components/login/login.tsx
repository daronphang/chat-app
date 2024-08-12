import { useEffect, useState } from 'react';
import { useNavigate } from 'react-router-dom';
import { useAppDispatch, useAppSelector } from 'core/redux/reduxHooks';
import { RpcError } from 'grpc-web';
import { useSnackbar } from 'notistack';

import { UserCredentials } from 'proto/user/user_pb';
import { setUser } from 'features/user/redux/userSlice';
import { Friend, FriendHash } from 'features/user/redux/user.interface';
import { RoutePath } from 'core/config/route.constant';
import { defaultSnackbarOptions } from 'core/config/snackbar.constant';
import styles from './login.module.scss';

export default function Login() {
  const [loading, isLoading] = useState<boolean>(false);
  const { enqueueSnackbar } = useSnackbar();
  const navigate = useNavigate();
  const config = useAppSelector(state => state.config);
  const dispatch = useAppDispatch();

  useEffect(() => {
    login();
  }, []);

  const login = async () => {
    try {
      const payload = new UserCredentials();
      payload.setEmail('daronphang@gmail.com');
      const resp = await config.api.USER_SERVICE.login(payload);
      const friends: FriendHash = {};
      resp.getFriendsList().forEach(row => {
        const friend: Friend = {
          userId: row.getUserid(),
          email: row.getEmail(),
          displayName: row.getDisplayname(),
          isOnline: false,
        };
        friends[friend.userId] = friend;
      });

      dispatch(
        setUser({
          userId: resp.getUserid(),
          email: resp.getEmail(),
          displayName: resp.getDisplayname(),
          friends: friends,
        })
      );
      navigate(RoutePath.CHAT);
    } catch (e) {
      const err = e as RpcError;

      if (err.code === 14) {
        enqueueSnackbar(config.apiError.NETWORK_ERROR, {
          ...defaultSnackbarOptions,
          variant: 'error',
        });
      } else {
        enqueueSnackbar('Invalid credentials', {
          ...defaultSnackbarOptions,
          variant: 'error',
        });
      }
    }
  };

  return (
    <>
      <div>Login page</div>
    </>
  );
}
