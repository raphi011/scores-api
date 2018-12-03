import { combineReducers } from 'redux';

import admin, { initialAdminState } from './admin';
import auth, { initialAuthState } from './auth';
import entities, { initialEntitiesState } from './entities';
import status, { initialStatusState } from './status';

export const initialState = {
  admin: initialAdminState,
  auth: initialAuthState,
  entities: initialEntitiesState,
  status: initialStatusState,
};

export default combineReducers({
  admin,
  auth,
  entities,
  status,
});
