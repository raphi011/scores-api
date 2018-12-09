import { serverAction } from './apiMiddleware';

export interface Action {
  type: string;
}

export async function dispatchActions(
  dispatch,
  actions = [],
  isServer,
  req?: object,
  res?: object,
) {
  for (const a of actions) {
    const action = isServer ? serverAction(a, req, res) : a;

    await dispatch(action);
  }
}

export async function dispatchAction(dispatch, action, isServer, req, res) {
  const result = await dispatch(
    isServer ? serverAction(action, req, res) : action,
  );
  return result;
}
