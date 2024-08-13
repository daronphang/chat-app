const wait = (ms: number) => new Promise(res => setTimeout(res, ms));

export const exponentialBackoff = async (fn: () => any, retries = 0, maxRetries = 3): Promise<any> => {
  try {
    return await fn();
  } catch (e) {
    if (retries > maxRetries) {
      throw e;
    }
    await wait(retries * 2 * 1000);

    return exponentialBackoff(fn, retries + 1, maxRetries);
  }
};
