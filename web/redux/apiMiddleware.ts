import * as http from 'http';
import fetch from 'isomorphic-unfetch';
import { BACKEND_URL, buildUrl, isJson } from '../api';
import { ApiAction, ApiActions } from '../redux/api/actions';
import * as actionNames from './actionNames';
import { userSelector } from './auth/selectors';

function getHost(req?: http.IncomingMessage): string {
  if (req) {
    return BACKEND_URL;
  }

  return `${window.location.origin}/api`;
}

export function serverAction(action, req, res) {
  return {
    ...action,
    isServer: true,
    req,
    res,
  };
}

async function doAction(
  store,
  action: ApiAction,
  isServer = false,
  req?: http.IncomingMessage,
  res?: http.ServerResponse,
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
  const { dispatch, getState } = store;

  if (isServer && req.headers.cookie) {
    const cookie = Array.isArray(req.headers.cookie)
      ? req.headers.cookie[0]
      : req.headers.cookie;

    // set the client's cookie for serverside request
    headers = { ...headers, cookie };
  }

  const endpoint = buildUrl(getHost(req), url, params);
  let responseCode = -1;

  try {
    const response = await fetch(endpoint, {
      body,
      credentials: 'same-origin',
      headers,
      method,
    });

    responseCode = response.status;

    if (responseCode === 401 && userSelector(getState()).isLoggedIn) {
      dispatch({ type: actionNames.LOGGEDOUT });
      dispatch({
        status: 'You have to be logged in for this action',
        type: actionNames.SET_STATUS,
      });
      return Promise.reject();
    }

    if (responseCode === 504) {
      dispatch({
        status: 'Cannot connect to server, please try again',
        type: actionNames.SET_STATUS,
      });
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
      statusMessage = message || statusMessage;
    }

    if (responseCode >= 200 && responseCode < 300) {
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

      return Promise.resolve({
        empty: false,
        response: payload,
        responseCode,
      });
    }

    dispatch({ type: actionNames.SET_STATUS, status: statusMessage });
  } catch (e) {
    if (error) {
      dispatch({ type: error, error: e.message });
    } else {
      dispatch({ type: actionNames.SET_STATUS, status: e.message });
    }
  }

  return Promise.reject({ responseCode });
}

const apiMiddleware = store => next => async action => {
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
      actions.map(a => doAction(store, a, isServer, req, res)),
    );
  } else {
    result = await doAction(store, apiAction, isServer, req, res);
  }

  return result;
};

export default apiMiddleware;
