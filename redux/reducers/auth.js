// @flow

import * as actionNames from '../actionNames';
import { createReducer } from '../reduxHelper';
import type { User } from '../../types';

export const initialAuthState = {
  user: null,
  loginRoute: null,
};

export type AuthStore = {
  user: ?User,
  loginRoute: ?string,
};

function loggedOut(state: AuthStore, action): AuthStore {
  const { loginRoute } = action.payload;

  return {
    ...state,
    loginRoute,
    user: null,
  };
}

function login(state: AuthStore, action): AuthStore {
  return {
    ...state,
    user: action.username,
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
  [actionNames.LOGIN]: login,
  [actionNames.SET_USER_OR_LOGINROUTE]: setUserOrLoginroute,
});

export default reducer;

export const loginRouteSelector = (state: AuthStore) => state.auth.loginRoute;
export const userSelector = (state: AuthStore) => ({
  isLoggedIn: !!state.auth.user,
  user: state.auth.user,
});
