import { User } from '../../types';
import * as actionNames from '../actionNames';
import { createReducer } from '../reduxHelper';
import { IStore } from '../store';

export const initialAuthState = {
  loginRoute: null,
  user: null,
};

export interface IAuthStore {
  user?: User;
  loginRoute?: string;
}

function loggedOut(state: IAuthStore, action): IAuthStore {
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

function setUserOrLoginroute(state: IAuthStore, action): IAuthStore {
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

export const loginRouteSelector = (state: IStore) => state.auth.loginRoute;
export const userSelector = (state: IStore) => ({
  isLoggedIn: !!state.auth.user,
  user: state.auth.user,
});
