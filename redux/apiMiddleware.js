import fetch from "isomorphic-unfetch";
import * as actionNames from "./actionNames";

function buildUrl(endpoint, params = {}) {
  let paramUrl = "";

  paramUrl = `?${Object.keys(params)
    .filter(key => params[key])
    .map(key => `${key}=${params[key]}`)
    .join("&")}`;

  const url = `${process.env.BACKEND_URL}/api/${endpoint}${paramUrl}`;

  return encodeURI(url);
}

const apiMiddleware = ({ getState, dispatch }) => next => async action => {
  if (action.type !== actionNames.API) {
    return next(action);
  }

  const {
    success,
    successParams = {},
    successStatus,
    error,
    headers,
    url,
    body,
    params,
    method = "GET"
  } = action;

  const endpoint = buildUrl(url, params);

  try {
    const response = await fetch(endpoint, { method, headers, body });
    if (response.status === 401) {
      dispatch({ type: action.LOGOUT });
      return;
    }

    if (response.status >= 200 && response.status < 300) {
      // result OK
      const contentType = response.headers.get("content-type");

      let payload = {};

      if (contentType && contentType.indexOf("application/json") !== -1) {
        payload = await response.json();
      }

      if (success) {
        dispatch({ type: success, payload, ...successParams });
      }
      if (successStatus) {
        dispatch({ type: actionNames.SET_STATUS, status: successStatus });
      }
    } else {
      // TODO: get error message from response
      dispatch({ type: error, error: "An error occured" });
    }
  } catch (e) {
    dispatch({ type: error, error: e.message });
  }
};

export default apiMiddleware;
