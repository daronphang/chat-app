import { useState } from 'react';
import Login from '../login/login';
import Signup from '../signup/signup';
import styles from './auth.module.scss';

export default function Auth() {
  const [isNewAccount, setIsNewAccount] = useState<boolean>(false);

  const showLogin = () => {
    setIsNewAccount(false);
  };

  const showSignup = () => {
    setIsNewAccount(true);
  };

  return (
    <div className={styles.wrapper}>
      {!isNewAccount && <Login showSignup={showSignup} />}
      {isNewAccount && <Signup showLogin={showLogin} />}
    </div>
  );
}
