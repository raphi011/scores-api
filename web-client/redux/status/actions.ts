import * as actionNames from '../actionNames';

export type StatusActionTypes = SetStatusAction | ClearStatusAction;

export interface SetStatusAction {
  type: typeof actionNames.SET_STATUS;
  status: string;
}

export const setStatusAction = (status: string) => ({
  status,
  type: actionNames.SET_STATUS,
});

export interface ClearStatusAction {
  type: typeof actionNames.CLEAR_STATUS;
}

export const clearStatusAction = () => ({
  type: actionNames.CLEAR_STATUS,
});
