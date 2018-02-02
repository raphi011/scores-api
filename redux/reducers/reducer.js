// @flow

import * as actionNames from '../actionNames';
import { createReducer } from '../reduxHelper';
import type { Match } from '../../types';
import type { Store } from '../storeTypes';

function normalizeMatchList(matches: Array<Match> = []) {
  const matchesMap = {};
  const matchesIds = [];

  matches.forEach(m => {
    matchesIds.push(m.id);
    matchesMap[m.id] = m;
  });

  return { matchesIds, matchesMap };
}

function receiveMatch(state: Store, action) {
  const matchesMap = {
    ...state.matchesMap,
    [action.id]: action.payload,
  };

  return {
    ...state,
    matchesMap,
  };
}

function receiveMatches(state: Store, action): Store {
  const { matchesIds, matchesMap } = normalizeMatchList(action.payload);

  return {
    ...state,
    matchesMap,
    matchesIds,
  };
}

function receivePlayers(state: Store, action): Store {
  const playersMap = {};
  const playerIds = [];
  action.payload.forEach(p => {
    playerIds.push(p.id);
    playersMap[p.id] = p;
  });

  return {
    ...state,
    playersMap,
    playerIds,
  };
}

function loggedOut(state: Store, action): Store {
  const { loginRoute } = action.payload;

  return {
    ...state,
    loginRoute,
    user: null,
  };
}

function login(state: Store, action): Store {
  return {
    ...state,
    user: action.username,
  };
}

function setUserOrLoginroute(state: Store, action): Store {
  const { user, loginRoute } = action.payload;
  return {
    ...state,
    loginRoute,
    user,
  };
}

function setStatus(state: Store, action): Store {
  return {
    ...state,
    status: action.status,
  };
}

function clearStatus(state: Store): Store {
  return {
    ...state,
    status: '',
  };
}

function removeMatch(state: Store, action): Store {
  const matchesIds = state.matchesIds.filter(id => id !== action.id);
  return {
    ...state,
    matchesIds,
  };
}

function receivePlayer(state: Store, action): Store {
  const playersMap = {
    ...state.playersMap,
    [action.id]: action.payload,
  };

  return {
    ...state,
    playersMap,
  };
}

function receiveStatistic(state: Store, action): Store {
  const statisticsMap = {
    ...state.statisticsMap,
    [action.playerId]: action.payload,
  };

  return {
    ...state,
    statisticsMap,
  };
}

function receivePlayerMatches(state: Store, action): Store {
  const { payload, playerId } = action;
  const { matchesIds, matchesMap } = normalizeMatchList(payload);

  const player = {
    ...state.playersMap[playerId],
    matchesIds,
  };

  return {
    ...state,
    matchesMap: {
      ...state.matchesMap,
      ...matchesMap,
    },
    playersMap: {
      ...state.playersMap,
      [playerId]: player,
    },
  };
}

function receiveStatistics(state: Store, action): Store {
  const statisticsMap = {};
  const statisticIds = [];
  action.payload.forEach(p => {
    statisticIds.push(p.playerId);
    statisticsMap[p.playerId] = p;
  });

  return {
    ...state,
    statisticsMap,
    statisticIds,
  };
}

const reducer = createReducer(
  {},
  {
    [actionNames.RECEIVE_PLAYER_MATCHES]: receivePlayerMatches,
    [actionNames.RECEIVE_STATISTIC]: receiveStatistic,
    [actionNames.RECEIVE_STATISTICS]: receiveStatistics,
    [actionNames.RECEIVE_MATCH]: receiveMatch,
    [actionNames.RECEIVE_PLAYER]: receivePlayer,
    [actionNames.RECEIVE_MATCHES]: receiveMatches,
    [actionNames.RECEIVE_PLAYERS]: receivePlayers,
    [actionNames.LOGGEDOUT]: loggedOut,
    [actionNames.LOGIN]: login,
    [actionNames.SET_USER_OR_LOGINROUTE]: setUserOrLoginroute,
    [actionNames.SET_STATUS]: setStatus,
    [actionNames.CLEAR_STATUS]: clearStatus,
    [actionNames.REMOVE_MATCH]: removeMatch,
  },
);

export default reducer;

export const playerMatchesSelector = (state: Store, playerId: number) => {
  const player = state.playersMap[playerId];

  if (!player || !player.matchesIds) return [];

  return player.matchesIds.map(id => state.matchesMap[id]);
};

export const statisticSelector = (state: Store, id: number) =>
  state.statisticsMap[id];
export const statisticsSelector = (state: Store) =>
  state.statisticIds.map(id => state.statisticsMap[id]);
export const loginRouteSelector = (state: Store) => state.loginRoute;
export const statusSelector = (state: Store) => state.status;
export const matchesSelector = (state: Store) =>
  state.matchesIds.map(id => state.matchesMap[id]);
export const matchSelector = (state: Store, id: number): Match =>
  state.matchesMap[id];
export const playerSelector = (state: Store, id: number) =>
  state.playersMap[id];
export const playersSelector = (state: Store) => ({
  playersMap: state.playersMap,
  playerIds: state.playerIds,
});
export const userSelector = (state: Store) => ({
  isLoggedIn: !!state.user,
  user: state.user,
});
