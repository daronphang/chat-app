import { useEffect } from 'react';
import './App.scss';
import Chat from 'features/chat/components/chat/chat';

// Font awesome declarations.
import {
  faCircleUser,
  faUsers,
  faChalkboardUser,
  faUsersRectangle,
  faGear,
  faRectangleList,
} from '@fortawesome/free-solid-svg-icons';
import { library } from '@fortawesome/fontawesome-svg-core';
library.add(faCircleUser, faUsers, faChalkboardUser, faUsersRectangle, faGear, faRectangleList);

function App() {
  useEffect(() => {
    getDeviceIdFromCookie();
  }, []);

  const getDeviceIdFromCookie = () => {
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
      <Chat />
    </div>
  );
}

export default App;
