import { applyMiddleware, createStore } from 'redux';
import { combineReducers } from 'redux';
import { composeWithDevTools } from 'redux-devtools-extension';

import apiMiddleware from './middleware/api';
import thunkMiddleware from './middleware/thunk';

import admin, { AdminStore, initialAdminState } from './admin/reducer';
import auth, { AuthStore, initialAuthState } from './auth/reducer';
import entities, {
  EntityStore,
  initialEntitiesState,
} from './entities/reducer';
import status, { initialStatusState, StatusStore } from './status/reducer';
import tournament, {
  TournamentStore,
  initialTournamentState,
} from './tournaments/reducer';

const reducer = combineReducers({
  admin,
  auth,
  entities,
  status,
  tournament,
});

export interface Store {
  auth: AuthStore;
  admin: AdminStore;
  entities: EntityStore;
  status: StatusStore;
  tournament: TournamentStore;
}

const middleware =
  process.env.NODE_ENV === 'development'
    ? composeWithDevTools(applyMiddleware(thunkMiddleware, apiMiddleware))
    : applyMiddleware(apiMiddleware);

const initStore = (state: Store = initialState) =>
  // @ts-ignore
  createStore<Store>(reducer, state, middleware);

export const initialState = {
  admin: initialAdminState,
  auth: initialAuthState,
  entities: initialEntitiesState,
  status: initialStatusState,
  tournament: initialTournamentState,
};

export default initStore;
