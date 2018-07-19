import * as actionNames from '../actionNames';
import { createReducer } from '../reduxHelper';
import { User } from '../../types';
import { Store } from '../store';

export const initialAuthState = {
  user: null,
  loginRoute: null,
};

export interface AuthStore {
  user?: User;
  loginRoute?: string;
}

function loggedOut(state: AuthStore, action): AuthStore {
  // todo
  let loginRoute = '';

  if (action.payload) {
    loginRoute = action.payload.loginRoute || '';
  }

  return {
    ...state,
    loginRoute,
    user: null,
  };
}

function setUserOrLoginroute(state: AuthStore, action): AuthStore {
  const { user, loginRoute } = action.payload;
  return {
    ...state,
    loginRoute,
    user,
  };
}

const reducer = createReducer(initialAuthState, {
  [actionNames.LOGGEDOUT]: loggedOut,
  [actionNames.SET_USER_OR_LOGINROUTE]: setUserOrLoginroute,
});

export default reducer;

export const loginRouteSelector = (state: Store) => state.auth.loginRoute;
export const userSelector = (state: Store) => ({
  isLoggedIn: !!state.auth.user,
  user: state.auth.user,
});
