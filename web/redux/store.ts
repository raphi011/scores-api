import { applyMiddleware, createStore } from 'redux';
import { combineReducers } from 'redux';
import { composeWithDevTools } from 'redux-devtools-extension';

import apiMiddleware from './apiMiddleware';

import admin, { AdminStore, initialAdminState } from './admin/reducer';
import auth, { AuthStore, initialAuthState } from './auth/reducer';
import entities, {
  EntityStore,
  initialEntitiesState,
} from './entities/reducer';
import status, { initialStatusState, StatusStore } from './status/reducer';

const reducer = combineReducers({
  admin,
  auth,
  entities,
  status,
});

export interface Store {
  auth: AuthStore;
  admin: AdminStore;
  entities: EntityStore;
  status: StatusStore;
}

const middleware =
  process.env.NODE_ENV === 'development'
    ? composeWithDevTools(applyMiddleware(apiMiddleware))
    : applyMiddleware(apiMiddleware);

const initStore = (state: Store = initialState) =>
  createStore<Store>(reducer, state, middleware);

export const initialState = {
  admin: initialAdminState,
  auth: initialAuthState,
  entities: initialEntitiesState,
  status: initialStatusState,
};

export default initStore;
