// @flow

import fetch from 'isomorphic-unfetch';
import * as actionNames from './actionNames';
import type { Action, ApiAction } from '../types';

type Params = { [string]: string };

function buildUrl(endpoint: string, params: Params = {}) {
  let paramUrl = '';
  const backendUrl = process.env.BACKEND_URL || '';

  paramUrl = `?${Object.keys(params)
    .filter(key => params[key])
    .map(key => `${key}=${params[key]}`)
    .join('&')}`;

  const url = `${backendUrl}/api/${endpoint}${paramUrl}`;

  return encodeURI(url);
}

function isJson(response): boolean {
  const contentType = response.headers.get('content-type');

  return contentType && contentType.indexOf('application/json') !== -1;
}

export function serverAction(action, req, res) {
  return {
    ...action,
    req,
    res,
    isServer: true,
  };
}

function fetchApiAction(action) {
  let { headers = {} } = action;
  const { url, body, params, method = 'GET', isServer = false, req } = action;

  if (isServer && req.headers.cookie) {
    // set the client's cookie for serverside request
    headers = { ...headers, cookie: req.headers.cookie };
  }

  const endpoint = buildUrl(url, params);

  return fetch(endpoint, {
    method,
    headers,
    body,
    credentials: 'same-origin',
  });
}

async function handleApiResponse(dispatch, action, response) {
  const { isServer, res, args, success, successStatus } = action;

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
      dispatch({ type: success, payload, ...args });
    }
    if (successStatus) {
      dispatch({ type: actionNames.SET_STATUS, status: successStatus });
    }

    if (Array.isArray(payload)) {
      const empty = !payload.length;
      return Promise.resolve({ empty });
    }

    return Promise.resolve({ empty: false });
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
}

const apiMiddleware = ({ dispatch }: Action => Promise<any>) => (
  next: Action => Promise<any>,
) => async (action: ApiAction) => {
  if (
    action.type !== actionNames.API ||
    action.type !== actionNames.API_MULTI
  ) {
    return next(action);
  }

  let actions;

  if (action.type === actionNames.API_MULTI) {
    // PARALLEL API REQUESTS
    actions = action.actions;
  } else {
    actions = [action];
  }

  const requests = actions.map(fetchApiAction);
  const responses = [];

  for (let i = 0; i < requests.length; i += 1) {
    try {
      responses.push(await requests[i]);
    } catch (e) {
      if (error) {
        dispatch({ type: error, error: e.message });
      } else {
        dispatch({ type: actionNames.SET_STATUS, status: e.message });
      }
    }
  }
};

export default apiMiddleware;
