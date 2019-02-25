import { QueryStringMapObject } from 'next';
import { SET_TOURNAMENT_FILTER } from '../actionNames';

export interface SetFilterAction {
  type: typeof SET_TOURNAMENT_FILTER;
  query: QueryStringMapObject;
}

export function setFilter(query: QueryStringMapObject): SetFilterAction {
  return {
    type: SET_TOURNAMENT_FILTER,
    query,
  };
}
