// @flow

import * as actionNames from '../actionNames';

export const setStatusAction = (status: string) => ({
  type: actionNames.SET_STATUS,
  status,
});

export const clearStatusAction = () => ({
  type: actionNames.CLEAR_STATUS,
});
