import { User } from '../../types';
import * as actionNames from '../actionNames';
import { createReducer } from '../reduxHelper';

export const initialAuthState = {
  loginRoute: '',
  user: null,
};

export interface AuthStore {
  user: User | null;
  loginRoute: string;
}

interface LoggedOutAction {
  payload: {
    loginRoute: string;
  };
}

function loggedOut(state: AuthStore, action: LoggedOutAction): AuthStore {
  const { loginRoute = '' } = action.payload;

  return {
    ...state,
    loginRoute,
    user: null,
  };
}

export interface SetUserOrLoginrouteAction {
  payload: {
    user?: User;
    loginRoute?: string;
  };
}

function setUserOrLoginroute(
  state: AuthStore,
  action: SetUserOrLoginrouteAction,
): AuthStore {
  const { user = null, loginRoute = '' } = action.payload;

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
