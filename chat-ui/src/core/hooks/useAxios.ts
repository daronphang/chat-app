import { useEffect, useState } from 'react';
import { AxiosResponse, AxiosRequestConfig } from 'axios';

export default function useAxios(fn: (cfg: AxiosRequestConfig) => Promise<AxiosResponse<any, any>>) {
  const [data, setData] = useState<any>(null);
  const [error, setError] = useState<string>('');
  const [loading, setLoading] = useState<boolean>(false);

  useEffect(() => {
    const controller = new AbortController();
    const timeoutId = setTimeout(() => controller.abort(), 30000);
    const defaultCfg: AxiosRequestConfig = {
      signal: controller.signal,
    };

    setLoading(true);

    fn(defaultCfg)
      .then(res => {
        setData(res.data);
        setLoading(false);
      })
      .catch(error => {
        let msg = 'An unknown error has occurred while making the request';
        if (error.name === 'TimeoutError') {
          msg = 'Request was aborted due to timeout error';
        } else if (error.name === 'AbortError') {
          // If axios is aborted, do nothing to prevent memory leaks
          // current implementation is an explicit timeout abort
          // unable to tell whether the abort was caused by timeout
          console.error('Fetch request was aborted by user action or timeout error');
          return;
        } else {
          if (error.response) {
            // request was made and server responded with status code > 300
            msg = error.response.data?.message ? error.response.data.message : msg;
          } else if (error.request) {
            // request was made but no response was received
            msg = error.request;
          } else {
            // unable to trigger request
            msg = error.message;
          }
        }
        setError(msg);
        setLoading(false);
      });

    return () => {
      controller.abort();
      clearTimeout(timeoutId);
    };
  }, []);

  return { loading, data, error };
}
