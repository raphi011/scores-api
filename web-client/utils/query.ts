import { QueryStringMapObject } from 'next';

export type Query = Record<string, string | string[]>;

export function str(query: QueryStringMapObject, key: string): string {
  const value = query[key];

  if (!value) {
    return '';
  }

  if (Array.isArray(value)) {
    return value[0];
  }

  return value;
}

export function oneOf<T>(
  query: QueryStringMapObject,
  key: string,
  available: T[] = [],
): T | '' {
  const value = str(query, key);

  if (!value) {
    return '';
  }

  const tValue = (value as unknown) as T;

  if (available.includes(tValue)) {
    return tValue;
  }

  return '';
}

export function oneOfDefault<T>(
  query: QueryStringMapObject,
  key: string,
  available: T[] = [],
  defaultValue: T,
): T {
  const value = oneOf(query, key, available);

  return value || defaultValue;
}
