import { applyMiddleware, createStore } from 'redux';
import { composeWithDevTools } from 'redux-devtools-extension';

import apiMiddleware, { serverAction } from './apiMiddleware';
import reducer, { initialState } from './reducers';

import { IAdminStore } from './reducers/admin';
import { IAuthStore } from './reducers/auth';
import { IEntityStore } from './reducers/entities';
import { IStatusStore } from './reducers/status';

export interface IStore {
  auth: IAuthStore;
  admin: IAdminStore;
  entities: IEntityStore;
  status: IStatusStore;
}

export async function dispatchAction(dispatch, action, isServer, req, res) {
  const result = await dispatch(
    isServer ? serverAction(action, req, res) : action,
  );
  return result;
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

const middleware =
  process.env.NODE_ENV === 'development'
    ? composeWithDevTools(applyMiddleware(apiMiddleware))
    : applyMiddleware(apiMiddleware);

const initStore = (state: IStore = initialState) =>
  createStore(reducer, state, middleware);

export default initStore;
