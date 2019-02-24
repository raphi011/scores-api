import * as actionNames from '../actionNames';
import { createReducer } from '../reduxHelper';
import * as Query from '../../utils/query';
import { SetFilterAction } from './actions';

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
    year: new Date().getFullYear,
    gender: 'M',
  },
};

interface NameKey {
  name: string;
  key: string;
}

export interface TournamentStore {
  available: {
    leagues: NameKey[];
    genders: NameKey[];
    seasons: number[];
  };
  filter: {
    leagues: string[];
    genders: string[];
    season: number;
  };
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

  return {
    available,
    filter: {
      season,
      genders,
      leagues,
    },
  };
}

const reducer = createReducer(initialTournamentState, {
  [actionNames.SET_TOURNAMENT_FILTER]: setFilter,
});

export default reducer;
