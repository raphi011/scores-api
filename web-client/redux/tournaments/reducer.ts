import * as actionNames from '../actionNames';
import { createReducer } from '../reduxHelper';
import { SetFilterAction } from './actions';
import { SetUserOrLoginrouteAction } from '../auth/reducer';
import {
  SETTING_TOURNAMENT_FILTER_LEAGUE_KEY,
  SETTING_TOURNAMENT_FILTER_GENDER_KEY,
  SETTING_TOURNAMENT_FILTER_SEASON_KEY,
} from '../../types';

export const initialTournamentState = {
  options: null,
  filter: null,
};

export interface Filter {
  leagues: string[];
  genders: string[];
  seasons: string;
}

export interface FilterOptions {
  leagues: string[];
  genders: string[];
  seasons: string[];
}

export interface TournamentStore {
  options: FilterOptions | null;
  filter: Filter | null;
}

function setFilter(
  state: TournamentStore,
  action: SetFilterAction,
): TournamentStore {
  const { filter: oldFilter, options } = state;
  const { filter } = action;

  const newFilter: Filter = {
    leagues: filter.leagues.length ? filter.leagues : oldFilter.leagues,
    genders: filter.genders.length ? filter.genders : oldFilter.genders,
    seasons: filter.seasons ? filter.seasons : oldFilter.seasons,
  };

  return {
    options,
    filter: newFilter,
  };
}

function setUserFilter(
  state: TournamentStore,
  action: SetUserOrLoginrouteAction,
): TournamentStore {
  const { filter } = state;

  if (!action.payload.user) {
    return state;
  }

  const { settings } = action.payload.user;

  let leagues = filter ? filter.leagues : [];
  let genders = filter ? filter.genders : [];
  let seasons = filter ? filter.seasons : '';

  if (settings[SETTING_TOURNAMENT_FILTER_LEAGUE_KEY]) {
    leagues = settings[SETTING_TOURNAMENT_FILTER_LEAGUE_KEY];
  }

  if (settings[SETTING_TOURNAMENT_FILTER_GENDER_KEY]) {
    genders = settings[SETTING_TOURNAMENT_FILTER_GENDER_KEY];
  }

  if (settings[SETTING_TOURNAMENT_FILTER_SEASON_KEY]) {
    seasons = settings[SETTING_TOURNAMENT_FILTER_SEASON_KEY];
  }

  return {
    ...state,
    filter: {
      leagues,
      genders,
      seasons,
    },
  };
}

function loadFilterOptions(
  state: TournamentStore,
  action: { payload: FilterOptions },
): TournamentStore {
  const { payload: options } = action;

  return {
    ...state,
    options,
  };
}

const reducer = createReducer(initialTournamentState, {
  [actionNames.SET_TOURNAMENT_FILTER]: setFilter,
  [actionNames.SET_USER_OR_LOGINROUTE]: setUserFilter,
  [actionNames.LOAD_TOURNAMENT_FILTER_OPTIONS]: loadFilterOptions,
});

export default reducer;
