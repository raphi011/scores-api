export interface Params {
  [key: string]: ParamValue;
}

type ParamValue = number | string | string[]

export function buildUrl(host: string, endpoint: string, params: Params = {}) {
  let paramUrl = '';

  const paramList = Object.keys(params)
    .filter(key => params[key])
    .map(key => keyValue(key, params[key]));

  paramUrl = paramList.length ? `?${paramList.join('&')}` : '';

  const url = `${host}/${endpoint}${paramUrl}`;

  return encodeURI(url);
}

function keyValue(key: string, value: ParamValue): string {
  return Array.isArray(value)
    ? value.map(v => `${key}=${v}`).join('&')
    : `${key}=${value}`
}

export function isJson(response: {
  headers: { get: (key: string) => string };
}): boolean {
  const contentType = response.headers.get('content-type');

  return !!contentType && contentType.indexOf('application/json') !== -1;
}

export const { BACKEND_URL } = process.env;
