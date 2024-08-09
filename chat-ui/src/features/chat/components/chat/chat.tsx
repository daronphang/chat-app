import { useState } from 'react';
import { Alert } from '@mui/material';

import ChannelDialogue from '../channel/channel';
import Drawer from '../drawer/drawer';
import Navbar from '../navbar/navbar';
import styles from './chat.module.scss';
import Loading from '../loading/loading';

export default function Chat() {
  const [loading, setLoading] = useState<boolean>(true);
  const [alert, setAlert] = useState<string>('');

  const handleLoading = (v: boolean) => {
    setLoading(v);
  };

  const handleAlert = (v: string) => {
    setAlert(v);
  };

  return (
    <>
      {loading && <Loading handleLoading={handleLoading} handleAlert={handleAlert} />}
      {!loading && alert && <Alert severity="error">{alert}</Alert>}
      {!loading && (
        <div className={styles.chat}>
          <Navbar />
          <Drawer />
          <ChannelDialogue />
        </div>
      )}
    </>
  );
}
