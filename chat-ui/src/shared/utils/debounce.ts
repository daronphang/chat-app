export const debounce = (duration: number, fn: (data: any) => {}) => {
  let timeoutId: NodeJS.Timeout | null = null;
  return (data: any) => {
    if (timeoutId) {
      clearTimeout(timeoutId);
    }
    timeoutId = setTimeout(() => {
      fn(data);
    }, duration);
  };
};
