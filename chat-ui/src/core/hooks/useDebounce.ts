import { useEffect, useState } from 'react';

interface DebounceProps {
  debounce: number;
  callback: (data: any) => void;
}

export default function useDebounce({ callback, debounce }: DebounceProps) {
  const [data, setData] = useState<any>(null);

  useEffect(() => {
    let timeoutId: NodeJS.Timeout | null = null;
    if (timeoutId) {
      clearTimeout(timeoutId);
    }
    timeoutId = setTimeout(() => {
      callback(data);
    }, debounce);

    return () => {
      if (timeoutId) {
        clearTimeout(timeoutId);
      }
    };
  }, [data]);

  return { setData };
}
