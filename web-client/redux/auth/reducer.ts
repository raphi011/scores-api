import { User } from '../../types';
import * as actionNames from '../actionNames';
import { createReducer } from '../reduxHelper';

export const initialAuthState = {
  loginRoute: '',
  user: null,
};

export type AuthStore = {
  user: User;
  loginRoute: string;
};

function loggedOut(state: AuthStore, action): AuthStore {
  const { loginRoute = '' } = action.payload;

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
