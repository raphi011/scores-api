// @flow
import fetch from 'isomorphic-unfetch';
import * as actionNames from './actionNames';
import type { Action, ApiAction } from '../types';

function buildUrl(endpoint, params = {}) {
  let paramUrl = '';

  paramUrl = `?${Object.keys(params)
    .filter(key => params[key])
    .map(key => `${key}=${params[key]}`)
    .join('&')}`;

  const url = `${process.env.BACKEND_URL}/api/${endpoint}${paramUrl}`;

  return encodeURI(url);
}

export function serverAction(action, req, res) {
  return {
    ...action,
    req,
    res,
    isServer: true,
  };
}

const apiMiddleware = ({ dispatch }: Action => Promise<any>) => (next: (Action) => Promise<any>) => async (action: ApiAction) => {
  if (action.type !== actionNames.API) {
    return next(action);
  }

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
    isServer = false,
    req,
    res,
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

    const { data, message } = await response.json();

    if (response.status >= 200 && response.status < 300) {
      if (success) {
        dispatch({ type: success, payload: data, ...successParams });
      }
      if (successStatus) {
        dispatch({ type: actionNames.SET_STATUS, status: successStatus });
      }
    } else {
      dispatch({ type: actionNames.SET_STATUS, status: message });
    }
  } catch (e) {
    if (error) {
      dispatch({ type: error, error: e.message });
    } else {
      dispatch({ type: actionNames.SET_STATUS, status: e.message });
    }
  }

  return Promise.resolve();
};

export default apiMiddleware;
