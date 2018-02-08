import { createStore, applyMiddleware } from 'redux';
import { composeWithDevTools } from 'redux-devtools-extension';

import reducer, { initialState } from './reducers';
import apiMiddleware, { serverAction } from './apiMiddleware';

import type { AuthStore } from './reducers/auth';
import type { EntitiesStore } from './reducers/entities';
import type { StatusStore } from './reducers/status';

export type Store = {
  auth: AuthStore,
  entities: EntitiesStore,
  status: StatusStore,
};

export async function dispatchAction(dispatch, isServer, req, res, action) {
  const result = await dispatch(
    isServer ? serverAction(action, req, res) : action,
  );
  return result;
}

export async function dispatchActions(
  dispatch,
  isServer,
  req,
  res,
  actions = [],
) {
  for (let i = 0; i < actions.length; i += 1) {
    const action = isServer ? serverAction(actions[i], req, res) : actions[i];

    await dispatch(action);
  }
}

const initStore = (state: Store = initialState) =>
  createStore(
    reducer,
    state,
    composeWithDevTools(applyMiddleware(apiMiddleware)),
  );

export default initStore;
