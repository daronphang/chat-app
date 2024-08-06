import React from 'react';
import ReactDOM from 'react-dom/client';
import { Provider } from 'react-redux';
import { BrowserRouter } from 'react-router-dom';
import { store } from 'core/redux/store';
import App from './App';
import reportWebVitals from './reportWebVitals';

// Styles.
import 'bootstrap/dist/css/bootstrap.min.css';
import './index.scss';
import { SnackbarProvider, closeSnackbar } from 'notistack';

const root = ReactDOM.createRoot(document.getElementById('root') as HTMLElement);

root.render(
  <React.StrictMode>
    <Provider store={store}>
      <BrowserRouter>
        <SnackbarProvider action={snackbarId => <button onClick={() => closeSnackbar(snackbarId)}>Dismiss</button>}>
          <App />
        </SnackbarProvider>
      </BrowserRouter>
    </Provider>
  </React.StrictMode>
);

// If you want to start measuring performance in your app, pass a function
// to log results (for example: reportWebVitals(console.log))
// or send to an analytics endpoint. Learn more: https://bit.ly/CRA-vitals
reportWebVitals();
