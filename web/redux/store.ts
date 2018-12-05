import { applyMiddleware, createStore } from 'redux';
import { combineReducers } from 'redux';
import { composeWithDevTools } from 'redux-devtools-extension';

import apiMiddleware, { serverAction } from './apiMiddleware';

import admin, { IAdminStore, initialAdminState } from './admin/reducer';
import auth, { IAuthStore, initialAuthState } from './auth/reducer';
import entities, {
  IEntityStore,
  initialEntitiesState,
} from './entities/reducer';
import status, { initialStatusState, IStatusStore } from './status/reducer';

const reducer = combineReducers({
  admin,
  auth,
  entities,
  status,
});

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

export const initialState = {
  admin: initialAdminState,
  auth: initialAuthState,
  entities: initialEntitiesState,
  status: initialStatusState,
};

export default initStore;
