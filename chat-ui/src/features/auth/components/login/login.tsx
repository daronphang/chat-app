import { useEffect, useState } from 'react';
import { useNavigate } from 'react-router-dom';
import { useAppDispatch, useAppSelector } from 'core/redux/reduxHooks';
import { RpcError } from 'grpc-web';

import { UserCredentials } from 'proto/user/user_pb';
import { setUser } from 'features/user/redux/userSlice';
import { RoutePaths } from 'core/config/route.constant';
import styles from './login.module.scss';
import { useSnackbar } from 'notistack';
import { defaultSnackbarOptions } from 'shared/utils/snackbar';
import { ApiErrors } from 'core/config/api.constant';

export default function Login() {
  const [loading, isLoading] = useState<boolean>(false);
  const { enqueueSnackbar } = useSnackbar();
  const navigate = useNavigate();
  const api = useAppSelector(state => state.config.api);
  const dispatch = useAppDispatch();

  useEffect(() => {
    login();
  }, []);

  const login = async () => {
    try {
      const payload = new UserCredentials();
      payload.setEmail('daronphang@gmail.com');
      const resp = await api.USER_SERVICE.login(payload);
      dispatch(
        setUser({
          userId: resp.getUserid(),
          email: resp.getEmail(),
          displayName: resp.getDisplayname(),
          contacts: [],
        })
      );
      navigate(RoutePaths.CHAT);
    } catch (e) {
      const err = e as RpcError;

      if (err.code === 14) {
        enqueueSnackbar(ApiErrors.NETWORK_ERROR, {
          ...defaultSnackbarOptions(),
          variant: 'error',
        });
      } else {
        enqueueSnackbar('Invalid credentials', {
          ...defaultSnackbarOptions(),
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
