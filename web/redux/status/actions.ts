import * as actionNames from '../actionNames';

export const setStatusAction = (status: string) => ({
  status,
  type: actionNames.SET_STATUS,
});

export const clearStatusAction = () => ({
  type: actionNames.CLEAR_STATUS,
});
