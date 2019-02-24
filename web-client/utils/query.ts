import { QueryStringMapObject } from 'next';

function isOfT<T>(value: any, options: T[]): value is T {
  return options.indexOf(value) !== -1;
}

export function multipleOf<T>(
  query: QueryStringMapObject,
  key: string,
  available: T[] = [],
): T[] {
  let values = query[key];

  if (!values) {
    return [];
  }

  if (!Array.isArray(values)) {
    values = [values];
  }

  const result: T[] = [];

  values.forEach(v => {
    if (isOfT(v, available)) {
      result.push(v);
    }
  });

  return result;
}

export function multipleOfDefault<T>(
  query: QueryStringMapObject,
  key: string,
  available: T[] = [],
  defaultValues: T[],
): T[] {
  const values = multipleOf(query, key, available);

  if (!values || (!values.length && defaultValues.length)) {
    return defaultValues;
  }

  return values;
}

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
