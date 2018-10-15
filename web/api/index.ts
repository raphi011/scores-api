export interface IParams {
  [key: string]: string;
}

export function buildUrl(host: string, endpoint: string, params: IParams = {}) {
  let paramUrl = '';

  const paramList = Object.keys(params)
    .filter(key => params[key])
    .map(key => `${key}=${params[key]}`);

  paramUrl = paramList.length ? `?${paramList.join('&')}` : '';

  const url = `${host}/${endpoint}${paramUrl}`;

  return encodeURI(url);
}

export function isJson(response: {
  headers: { get: (key: string) => string };
}): boolean {
  const contentType = response.headers.get('content-type');

  return !!contentType && contentType.indexOf('application/json') !== -1;
}

export const { BACKEND_URL } = process.env;
