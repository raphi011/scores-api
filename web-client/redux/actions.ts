import * as http from 'http';
import { Dispatch } from 'redux';

import { serverAction } from './apiMiddleware';

export interface Action {
  type: string;
}

export async function dispatchActions(
  dispatch: Dispatch,
  actions = [],
  isServer: boolean,
  req?: http.IncomingMessage,
  res?: http.OutgoingMessage,
) {
  for (const a of actions) {
    const action = isServer ? serverAction(a, req, res) : a;

    await dispatch(action);
  }
}

export async function dispatchAction(
  dispatch: Dispatch,
  action,
  isServer: boolean,
  req?: http.IncomingMessage,
  res?: http.ServerResponse,
) {
  const result = await dispatch(
    isServer ? serverAction(action, req, res) : action,
  );
  return result;
}
