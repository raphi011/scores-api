import { applyMiddleware, createStore } from 'redux';
import { composeWithDevTools } from 'redux-devtools-extension';

import apiMiddleware, { serverAction } from './apiMiddleware';
import reducer, { initialState } from './reducers';

import { AuthStore } from './reducers/auth';
import { EntityStore } from './reducers/entities';
import { StatusStore } from './reducers/status';

export interface Store {
  auth: AuthStore;
  entities: EntityStore;
  status: StatusStore;
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
  for (let i = 0; i < actions.length; i += 1) {
    const action = isServer ? serverAction(actions[i], req, res) : actions[i];

    await dispatch(action);
  }
}

const middleware = process.env.NODE_ENV === 'development'
 ? composeWithDevTools(applyMiddleware(apiMiddleware))
 : applyMiddleware(apiMiddleware);

const initStore = (state: Store = initialState) =>
  createStore(
    reducer,
    state,
    middleware,
  );

export default initStore;
