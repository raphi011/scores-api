import { ApiAction } from '../../redux/api/actions';
import * as actionNames from '../actionNames';

export const userOrLoginRouteAction = (): ApiAction => ({
  method: 'GET',
  success: actionNames.SET_USER_OR_LOGINROUTE,
  type: actionNames.API,
  url: 'user-or-login',
});

export const loginWithPasswordAction = (credentials): ApiAction => ({
  body: JSON.stringify(credentials),
  method: 'POST',
  success: actionNames.SET_USER_OR_LOGINROUTE,
  successStatus: 'Logged in',
  type: actionNames.API,
  url: 'pw-auth',
});

export const logoutAction = (): ApiAction => ({
  method: 'POST',
  success: actionNames.LOGGEDOUT,
  successStatus: 'Logged out',
  type: actionNames.API,
  url: 'logout',
});

export const loggedOutAction = () => ({
  type: actionNames.LOGGEDOUT,
});
