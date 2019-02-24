import * as http from 'http';
import { Dispatch } from 'redux';

import { serverAction } from './apiMiddleware';

export interface Action {
  type: string;
}

export async function dispatchActions(
  dispatch: Dispatch,
  actions = [],
  req?: http.IncomingMessage,
  res?: http.OutgoingMessage,
) {
  for (const a of actions) {
    const action = req && res ? serverAction(a, req, res) : a;

    await dispatch(action);
  }
}

export async function dispatchAction(
  dispatch: Dispatch,
  action: any,
  req?: http.IncomingMessage,
  res?: http.ServerResponse,
) {
  const result = await dispatch(
    req && res ? serverAction(action, req, res) : action,
  );
  return result;
}
