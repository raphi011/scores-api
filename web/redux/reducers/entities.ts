import { createSelector } from 'reselect';

import * as actionNames from '../actionNames';
import { denorm, norm } from '../entitySchemas';
import { createReducer } from '../reduxHelper';

import { EntityName, EntityType } from '../../types';
import { IStore } from '../store';

export interface IEntityStore {
  [key: string]: { values: { [key: number]: EntityType } };
}

export interface IEntityMap {
  [key: string]: { [key: number]: EntityType };
}

export interface IReceiveEntityAction {
  payload: any[];
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

export interface IDeleteEntityAction {
  payload: any;
  entityName: EntityName;
  listNames: string[];
}

export const initialEntitiesState = {
  group: { values: {}, all: [] },
  match: { values: {}, all: [], byPlayer: {}, byGroup: {} },
  player: { values: {}, all: [], byGroup: {} },
  statistic: {
    all: [],
    byGroup: {},
    byPlayer: {},
    byPlayerTeam: {},
    values: {},
  },
  team: { values: {} },
  tournament: { values: {}, all: [], byLeague: {} },
  volleynetplayer: {
    all: [],
    ladder: {},
    search: [],
    values: {},
  },
};

function deleteEntities(state, action: IDeleteEntityAction) {
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

function receiveEntities(state: IEntityStore, action: IReceiveEntityAction) {
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

        if (previousList) {
          list = previousList;
        }
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

function receiveUser(state: IEntityStore, action) {
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
    match,
    player,
    statistic,
    team,
    tournament,
    volleynetplayer,
  }),
);

export const allPlayersSelector = (state: IStore) =>
  state.entities.player.all.length
    ? denorm('player', entityMapSelector(state), state.entities.player.all)
    : [];

export const groupSelector = (state: IStore, groupId: number) =>
  denorm('group', entityMapSelector(state), groupId);

export const groupPlayersSelector = (state: IStore, groupId: number) =>
  (state.entities.player.byGroup[groupId] || []).length
    ? denorm(
        'player',
        entityMapSelector(state),
        state.entities.player.byGroup[groupId],
      )
    : [];

export const matchSelector = (state: IStore, id: number) =>
  denorm('match', entityMapSelector(state), id);

const allMatchIdsSelector = state =>
  state.entities.match.all.length ? state.entities.match.all : [];

export const allMatchesSelector = createSelector(
  allMatchIdsSelector,
  entityMapSelector,
  (ids, entities) => denorm('match', entities, ids),
);

export const playerSelector = (state: IStore, playerId: number) =>
  denorm('player', entityMapSelector(state), playerId);

export const matchesByGroupSelector = (state: IStore, groupId: number) =>
  (state.entities.match.byGroup[groupId] || []).length
    ? denorm(
        'match',
        entityMapSelector(state),
        state.entities.match.byGroup[groupId],
      )
    : [];

export const matchesByPlayerSelector = (state: IStore, playerId: number) =>
  (state.entities.match.byPlayer[playerId] || []).length
    ? denorm(
        'match',
        entityMapSelector(state),
        state.entities.match.byPlayer[playerId],
      )
    : [];

export const allStatisticSelector = (state: IStore) =>
  state.entities.statistic.all.length
    ? denorm(
        'statistic',
        entityMapSelector(state),
        state.entities.statistic.all,
      )
    : [];

export const statisticByPlayerTeamSelector = (
  state: IStore,
  playerId: number,
) =>
  (state.entities.statistic.byPlayerTeam[playerId] || []).length
    ? denorm(
        'statistic',
        entityMapSelector(state),
        state.entities.statistic.byPlayerTeam[playerId],
      )
    : [];

export const statisticByGroupSelector = (state: IStore, groupId: number) =>
  (state.entities.statistic.byGroup[groupId] || []).length
    ? denorm(
        'statistic',
        entityMapSelector(state),
        state.entities.statistic.byGroup[groupId],
      )
    : [];

export const statisticByPlayerSelector = (state: IStore, playerId: number) =>
  (state.entities.statistic.byPlayer[playerId] || []).length
    ? denorm(
        'statistic',
        entityMapSelector(state),
        state.entities.statistic.byPlayer[playerId][0],
      )
    : null;

export const tournamentsByLeagueSelector = (state: IStore, league: string) =>
  (state.entities.tournament.byLeague[league] || []).length
    ? denorm(
        'tournament',
        entityMapSelector(state),
        state.entities.tournament.byLeague[league],
      )
    : null;

export const tournamentSelector = (state: IStore, tournamentId: number) =>
  denorm('tournament', entityMapSelector(state), tournamentId);

export const ladderVolleynetplayerSelector = (state: IStore, gender: string) =>
  (state.entities.volleynetplayer.ladder[gender] || []).length
    ? denorm(
        'volleynetplayer',
        entityMapSelector(state),
        state.entities.volleynetplayer.ladder[gender],
      )
    : [];

export const searchVolleynetplayerSelector = (state: IStore) =>
  (state.entities.volleynetplayer.search || []).length
    ? denorm(
        'volleynetplayer',
        entityMapSelector(state),
        state.entities.volleynetplayer.search,
      )
    : [];
