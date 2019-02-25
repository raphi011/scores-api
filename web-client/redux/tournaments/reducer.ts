import * as actionNames from '../actionNames';
import { createReducer } from '../reduxHelper';
import * as Query from '../../utils/query';
import { SetFilterAction } from './actions';

const thisYear = new Date().getFullYear();

export const initialTournamentState = {
  available: {
    leagues: [
      { name: 'Junior Tour', key: 'junior-tour' },
      { name: 'Amateur Tour', key: 'amateur-tour' },
      { name: 'Pro Tour', key: 'pro-tour' },
    ],
    seasons: [2018, 2019],
    genders: [{ name: 'Female', key: 'W' }, { name: 'Male', key: 'M' }],
  },
  filter: {
    season: thisYear,
    genders: ['M'],
    leagues: ['amateur-tour'],
  },
};

interface NameKey {
  name: string;
  key: string;
}

export interface Filter {
  leagues: string[];
  genders: string[];
  season: number;
}

export interface TournamentStore {
  available: {
    leagues: NameKey[];
    genders: NameKey[];
    seasons: number[];
  };
  filter: Filter;
}

function setFilter(
  state: TournamentStore,
  action: SetFilterAction,
): TournamentStore {
  const { available, filter } = state;
  const { query } = action;

  const season = Query.oneOfDefault(
    query,
    'season',
    available.seasons,
    filter.season,
  );

  const genders = Query.multipleOfDefault(
    query,
    'gender',
    available.genders.map(g => g.key),
    filter.genders,
  );

  const leagues = Query.multipleOfDefault(
    query,
    'league',
    available.leagues.map(g => g.key),
    filter.leagues,
  );

  const newFilter: Filter = {
    season,
    genders,
    leagues,
  };

  if (!hasFilterChanged(filter, newFilter)) {
    return state;
  }

  return {
    available,
    filter: newFilter,
  };
}

function hasFilterChanged(oldFilter: Filter, newFilter: Filter): boolean {
  if (oldFilter.season !== newFilter.season) {
    return true;
  }

  if (oldFilter.leagues.some(l => newFilter.leagues.indexOf(l) !== -1)) {
    return true;
  }

  if (oldFilter.genders.some(l => newFilter.genders.indexOf(l) !== -1)) {
    return true;
  }

  return false;
}

const reducer = createReducer(initialTournamentState, {
  [actionNames.SET_TOURNAMENT_FILTER]: setFilter,
});

export default reducer;
