import { createSelector } from 'reselect';

import * as actionNames from '../actionNames';
import { createReducer } from '../reduxHelper';
import { denorm, norm } from '../entitySchemas';

import { EntityName, EntityType } from '../../types';
import { Store } from '../store';

export interface EntityStore {
  [key: string]: { values: { [key: number]: EntityType } };
}

export interface EntityMap {
  [key: string]: { [key: number]: EntityType };
}

export interface ReceiveEntityAction {
  payload: Object[];
  entityName: EntityName;
  assignId?: boolean;
  listOptions?: {
    [key: string]: {
      name: string;
      key?: number;
      mode?: 'replace' | 'append';
    };
  };
}

export interface DeleteEntityAction {
  payload: Object;
  entityName: EntityName;
  listNames: string[];
}

export const initialEntitiesState = {
  group: { values: {}, all: [] },
  player: { values: {}, all: [], byGroup: {} },
  team: { values: {} },
  match: { values: {}, all: [], byPlayer: {}, byGroup: {} },
  tournament: { values: {}, all: [], byLeague: {} },
  volleynetplayer: { values: {}, all: [], search: [] },
  statistic: {
    values: {},
    all: [],
    byPlayer: {},
    byPlayerTeam: {},
    byGroup: {},
  },
};

function deleteEntities(state, action: DeleteEntityAction) {
  const { entityName, listNames, payload } = action;

  const deletedId = payload.id;
  // TODO: delete from values
  // TODO: delete from [listName][listKey] lists
  const newLists = listNames.reduce((listObj, name) => {
    listObj[name] = state[entityName][name].filter(id => id !== deletedId);
    return listObj;
  }, {});

  return {
    ...state,
    [entityName]: {
      ...state[entityName],
      ...newLists,
    },
  };
}

function receiveEntities(state: EntityStore, action: ReceiveEntityAction) {
  // STEP 1: normalize entites
  const { entityName, payload, assignId = false, listOptions = {} } = action;

  const { entities } = norm(entityName, payload, assignId);

  const newState = { ...state };

  // STEP 2: add entities to entity map(s)
  Object.keys(entities).forEach((entityKey: EntityName) => {
    const statePart = { ...state[entityKey] };

    const newIds = Object.keys(entities[entityKey]).map(n =>
      Number.parseInt(n, 10),
    );

    statePart.values = {
      ...state[entityKey].values,
      ...entities[entityKey],
    };

    const options = listOptions[entityKey];

    // STEP 3: append or replace ids
    if (options) {
      let list = [];

      if (options.mode === 'append') {
        const previousList = options.key
          ? (state[entityKey][options.name] || {})[options.key]
          : state[entityKey][options.name];

        if (previousList) list = previousList;
      }

      list = [...list, ...newIds];

      statePart[options.name] = options.key
        ? { ...(state[entityKey][options.name] || {}), [options.key]: list }
        : list;
    }

    newState[entityKey] = statePart;
  });

  return newState;
}

function receiveUser(state: EntityStore, action) {
  const { user } = action.payload;

  if (!user || !user.player) {
    return state;
  }

  return receiveEntities(state, {
    entityName: 'player',
    payload: user.player,
  });
}

const reducer = createReducer(initialEntitiesState, {
  [actionNames.RECEIVE_ENTITIES]: receiveEntities,
  [actionNames.DELETE_ENTITIES]: deleteEntities,
  [actionNames.SET_USER_OR_LOGINROUTE]: receiveUser,
});

export default reducer;

const groupMap = state => state.entities.group.values;
const playerMap = state => state.entities.player.values;
const teamMap = state => state.entities.team.values;
const matchMap = state => state.entities.match.values;
const statisticMap = state => state.entities.statistic.values;
const tournamentMap = state => state.entities.tournament.values;
const volleynetplayerMap = state => state.entities.volleynetplayer.values;

export const entityMapSelector = createSelector(
  groupMap,
  playerMap,
  teamMap,
  matchMap,
  statisticMap,
  tournamentMap,
  volleynetplayerMap,
  (group, player, team, match, statistic, tournament, volleynetplayer) => ({
    group,
    player,
    team,
    match,
    statistic,
    tournament,
    volleynetplayer,
  }),
);

export const allPlayersSelector = (state: Store) =>
  state.entities.player.all.length
    ? denorm('player', entityMapSelector(state), state.entities.player.all)
    : [];

export const groupPlayersSelector = (state: Store, groupId: number) =>
  (state.entities.player.byGroup[groupId] || []).length
    ? denorm(
        'player',
        entityMapSelector(state),
        state.entities.player.byGroup[groupId],
      )
    : [];

export const matchSelector = (state: Store, id: number) =>
  denorm('match', entityMapSelector(state), id);

const allMatchIdsSelector = state =>
  state.entities.match.all.length ? state.entities.match.all : [];

export const allMatchesSelector = createSelector(
  allMatchIdsSelector,
  entityMapSelector,
  (ids, entities) => denorm('match', entities, ids),
);

export const playerSelector = (state: Store, playerId: number) =>
  denorm('player', entityMapSelector(state), playerId);

export const matchesByGroupSelector = (state: Store, groupId: number) =>
  (state.entities.match.byGroup[groupId] || []).length
    ? denorm(
        'match',
        entityMapSelector(state),
        state.entities.match.byGroup[groupId],
      )
    : [];

export const matchesByPlayerSelector = (state: Store, playerId: number) =>
  (state.entities.match.byPlayer[playerId] || []).length
    ? denorm(
        'match',
        entityMapSelector(state),
        state.entities.match.byPlayer[playerId],
      )
    : [];

export const allStatisticSelector = (state: Store) =>
  state.entities.statistic.all.length
    ? denorm(
        'statistic',
        entityMapSelector(state),
        state.entities.statistic.all,
      )
    : [];

export const statisticByPlayerTeamSelector = (state: Store, playerId: number) =>
  (state.entities.statistic.byPlayerTeam[playerId] || []).length
    ? denorm(
        'statistic',
        entityMapSelector(state),
        state.entities.statistic.byPlayerTeam[playerId],
      )
    : [];

export const statisticByGroupSelector = (state: Store, groupId: number) =>
  (state.entities.statistic.byGroup[groupId] || []).length
    ? denorm(
        'statistic',
        entityMapSelector(state),
        state.entities.statistic.byGroup[groupId],
      )
    : [];

export const statisticByPlayerSelector = (state: Store, playerId: number) =>
  (state.entities.statistic.byPlayer[playerId] || []).length
    ? denorm(
        'statistic',
        entityMapSelector(state),
        state.entities.statistic.byPlayer[playerId][0],
      )
    : null;

export const tournamentsByLeagueSelector = (state: Store, league: string) =>
  (state.entities.tournament.byLeague[league] || []).length
    ? denorm(
        'tournament',
        entityMapSelector(state),
        state.entities.tournament.byLeague[league],
      )
    : null;

export const tournamentSelector = (state: Store, tournamentId: number) =>
  denorm('tournament', entityMapSelector(state), tournamentId);

export const searchVolleynetplayerSelector = (state: Store) =>
  (state.entities.volleynetplayer.search || []).length
    ? denorm(
        'volleynetplayer',
        entityMapSelector(state),
        state.entities.volleynetplayer.search,
      )
    : [];
