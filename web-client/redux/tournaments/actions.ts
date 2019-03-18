import {
  API,
  SET_TOURNAMENT_FILTER,
  LOAD_TOURNAMENT_FILTER_OPTIONS,
} from '../actionNames';
import { Filter } from './reducer';
import { ApiAction } from '../api/actions';

export interface SetFilterAction {
  type: typeof SET_TOURNAMENT_FILTER;
  filter: Filter;
}

export const setFilterAction = (filter: Filter): SetFilterAction => ({
  type: SET_TOURNAMENT_FILTER,
  filter,
});

export const loadFilterOptionsAction = (): ApiAction => ({
  method: 'GET',
  success: LOAD_TOURNAMENT_FILTER_OPTIONS,
  type: API,
  url: 'filters',
});
