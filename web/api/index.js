type Params = { [string]: string };

export function buildUrl(endpoint: string, params: Params = {}) {
  let paramUrl = '';

  const paramList = Object.keys(params)
    .filter(key => params[key])
    .map(key => `${key}=${params[key]}`);

  paramUrl = paramList.length ? `?${paramList.join('&')}` : '';

  const url = `${process.env.BACKEND_URL}/api/${endpoint}${paramUrl}`;

  return encodeURI(url);
}

export function isJson(response): boolean {
  const contentType = response.headers.get('content-type');

  return contentType && contentType.indexOf('application/json') !== -1;
}

export const { BACKEND_URL } = process.env;
