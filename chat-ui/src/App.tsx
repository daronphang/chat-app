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
} from '@fortawesome/free-solid-svg-icons';
import { library } from '@fortawesome/fontawesome-svg-core';
import { ProtectedRoutes } from 'core/guards/authGuard';
import Login from 'features/auth/components/login/login';
import { RoutePaths } from 'core/config/route.constant';
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
  faClock
);

function App() {
  useEffect(() => {
    getDeviceIdFromCookie();
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
          <Route path={RoutePaths.CHAT} element={<Chat />} />
        </Route>
        <Route path={RoutePaths.LOGIN} element={<Login />} />
        <Route path="*" element={<Navigate to={RoutePaths.CHAT} />} />
      </Routes>
    </div>
  );
}

export default App;
