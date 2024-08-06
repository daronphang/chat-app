import { OptionsObject } from 'notistack';

export const defaultSnackbarOptions = (): OptionsObject => {
  return {
    anchorOrigin: { horizontal: 'right', vertical: 'top' },
    autoHideDuration: 5000,
    preventDuplicate: true,
  };
};
