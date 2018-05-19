// @flow

import fetch from 'isomorphic-unfetch';
import * as actionNames from './actionNames';
import type { Action, ApiAction, ApiActions } from '../types';
import { buildUrl, isJson } from '../api';

export function serverAction(action, req, res) {
  return {
    ...action,
    req,
    res,
    isServer: true,
  };
}

async function doAction(
  dispatch,
  action: ApiAction,
  isServer = false,
  req,
  res,
) {
  let { headers = {} } = action;
  const {
    success,
    successParams = {},
    successStatus,
    error,
    url,
    body,
    params,
    method = 'GET',
  } = action;

  if (isServer && req.headers.cookie) {
    // set the client's cookie for serverside request
    headers = { ...headers, cookie: req.headers.cookie };
  }

  const endpoint = buildUrl(url, params);

  try {
    const response = await fetch(endpoint, {
      method,
      headers,
      body,
      credentials: 'same-origin',
    });
    if (response.status === 401) {
      dispatch({ type: actionNames.LOGGEDOUT });
      dispatch({
        type: actionNames.SET_STATUS,
        status: 'You have to be logged in for this action',
      });
      return Promise.reject();
    }

    if (isServer) {
      const setCookie = response.headers.get('Set-Cookie');

      if (setCookie) {
        res.setHeader('Set-Cookie', setCookie);
      }
    }

    let payload;
    let statusMessage = 'Unknown error';

    if (isJson(response)) {
      const { data, message } = await response.json();
      payload = data;
      statusMessage = message;
    }

    if (response.status >= 200 && response.status < 300) {
      if (success) {
        dispatch({ type: success, payload, ...successParams });
      }
      if (successStatus) {
        dispatch({ type: actionNames.SET_STATUS, status: successStatus });
      }

      if (Array.isArray(payload)) {
        const empty = !payload.length;
        return Promise.resolve({ empty });
      }

      return Promise.resolve({ empty: false, response: payload });
    }

    // error ...
    if (response.status === 504) {
      dispatch({
        type: actionNames.SET_STATUS,
        status: 'Cannot connect to server, please try again',
      });
    } else {
      dispatch({ type: actionNames.SET_STATUS, status: statusMessage });
    }
  } catch (e) {
    if (error) {
      dispatch({ type: error, error: e.message });
    } else {
      dispatch({ type: actionNames.SET_STATUS, status: e.message });
    }
  }

  return Promise.reject();
}

const apiMiddleware = ({ dispatch }: Action => Promise<any>) => (
  next: Action => Promise<any>,
) => async (action: Action) => {
  if (
    action.type !== actionNames.API &&
    action.type !== actionNames.API_MULTI
  ) {
    return next(action);
  }

  const apiAction: ApiAction | ApiActions = action;

  const { req, res, isServer } = apiAction;

  let result;

  if (apiAction.type === actionNames.API_MULTI) {
    const { actions } = apiAction;

    result = await Promise.all(
      actions.map(a => doAction(dispatch, a, isServer, req, res)),
    );
  } else {
    result = await doAction(dispatch, apiAction, isServer, req, res);
  }

  return result;
};

export default apiMiddleware;
