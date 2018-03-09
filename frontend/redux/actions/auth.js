// @flow

import * as actionNames from '../actionNames';
import type { ApiAction } from '../../types';

export const userOrLoginRouteAction = (): ApiAction => ({
  type: actionNames.API,
  method: 'GET',
  url: 'userOrLoginRoute',
  success: actionNames.SET_USER_OR_LOGINROUTE,
});

export const loginWithPasswordAction = (credentials): ApiAction => ({
  type: actionNames.API,
  method: 'POST',
  url: 'pwAuth',
  body: JSON.stringify(credentials),
  successStatus: 'Logged in',
  success: actionNames.SET_USER_OR_LOGINROUTE,
});

export const logoutAction = (): ApiAction => ({
  type: actionNames.API,
  method: 'POST',
  url: 'logout',
  success: actionNames.LOGGEDOUT,
  successStatus: 'Logged out',
});

export const loggedOutAction = () => ({
  type: actionNames.LOGGEDOUT,
});
