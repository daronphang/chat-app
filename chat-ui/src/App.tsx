import { useEffect } from 'react';
import Chat from 'features/chat/components/chat/chat';
import { Routes, Route, Navigate } from 'react-router-dom';
import './App.scss';

// Font awesome declarations.
import {
  faCircleUser,
  faUsers,
  faChalkboardUser,
  faUsersRectangle,
  faGear,
  faRectangleList,
  faPaperPlane,
  faCheck,
  faCheckDouble,
  faClock,
  faUserPlus,
  faArrowLeft,
} from '@fortawesome/free-solid-svg-icons';
import { library } from '@fortawesome/fontawesome-svg-core';
import { ProtectedRoutes } from 'core/guards/authGuard';
import Login from 'features/auth/components/login/login';
import { RoutePath } from 'core/config/route.constant';
import { useAppDispatch } from 'core/redux/reduxHooks';
import { setDeviceId } from 'core/config/configSlice';
library.add(
  faCircleUser,
  faUsers,
  faChalkboardUser,
  faUsersRectangle,
  faGear,
  faRectangleList,
  faPaperPlane,
  faCheck,
  faCheckDouble,
  faClock,
  faUserPlus,
  faArrowLeft
);

function App() {
  const dispatch = useAppDispatch();
  useEffect(() => {
    dispatch(setDeviceId(getDeviceIdFromCookie()));
  }, []);

  const getDeviceIdFromCookie = (): string => {
    let deviceId: string;
    const cookie = document.cookie.match('(^|;)\\s*' + 'deviceId' + '\\s*=\\s*([^;]+)');
    if (cookie) {
      deviceId = cookie.pop() as string;
    } else {
      deviceId = crypto.randomUUID();
      document.cookie = `deviceId=${deviceId}`;
    }
    return deviceId;
  };

  return (
    <div className="App">
      <Routes>
        <Route element={<ProtectedRoutes />}>
          <Route path={RoutePath.CHAT} element={<Chat />} />
        </Route>
        <Route path={RoutePath.LOGIN} element={<Login />} />
        <Route path="*" element={<Navigate to={RoutePath.CHAT} />} />
      </Routes>
    </div>
  );
}

export default App;
