import { combineReducers } from 'redux';

import auth, { initialAuthState } from './auth';
import entities, { initialEntitiesState } from './entities';
import status, { initialStatusState } from './status';

export const initialState = {
  auth: initialAuthState,
  entities: initialEntitiesState,
  status: initialStatusState,
};

export default combineReducers({
  auth,
  entities,
  status,
});
