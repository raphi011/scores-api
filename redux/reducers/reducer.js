// @flow

import * as actionNames from '../actionNames';

function normalizeMatchList(matches = []) {
  const matchesMap = {};
  const matchesIds = [];

  matches.forEach(m => {
    matchesIds.push(m.id);
    matchesMap[m.id] = m;
  });

  return { matchesIds, matchesMap };
}

function createReducer(initialState, handlers) {
  return function reducer(state = initialState, action) {
    if (handlers.hasOwnProperty(action.type)) {
      return handlers[action.type](state, action);
    }

    return state;
  };
}

function receiveMatch(state, action) {
  const matchesMap = {
    ...state.matchesMap,
    [action.id]: action.payload,
  };

  return {
    ...state,
    matchesMap,
  };
}

function receiveMatches(state, action) {
  const { matchesIds, matchesMap } = normalizeMatchList(action.payload);

  return {
    ...state,
    matchesMap,
    matchesIds,
  };
}

function receivePlayers(state, action) {
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

function loggedOut(state, action) {
  const { loginRoute } = action.payload;

  return {
    ...state,
    loginRoute,
    user: null,
  };
}

function login(state, action) {
  return {
    ...state,
    user: action.username,
  };
}

function setUserOrLoginroute(state, action) {
  const { user, loginRoute } = action.payload;
  return {
    ...state,
    loginRoute,
    user,
  };
}

function setStatus(state, action) {
  return {
    ...state,
    status: action.status,
  };
}

function clearStatus(state) {
  return {
    ...state,
    status: '',
  };
}

function removeMatch(state, action) {
  const matchesIds = state.matchesIds.filter(id => id !== action.id);
  return {
    ...state,
    matchesIds,
  };
}

function receivePlayer(state, action) {
  const playersMap = {
    ...state.playersMap,
    [action.id]: action.payload,
  };

  return {
    ...state,
    playersMap,
  };
}

function receiveStatistic(state, action) {
  const statisticsMap = {
    ...state.statisticsMap,
    [action.playerId]: action.payload,
  };

  return {
    ...state,
    statisticsMap,
  };
}

function receivePlayerMatches(state, action) {
  const { payload, playerId } = action;
  const { matchesIds, matchesMap } = normalizeMatchList(payload);

  const player =  {
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

function receiveStatistics(state, action) {
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

export const playerMatchesSelector = (state, playerId: number) => {
  const player = state.playersMap[playerId];

  if (!player || !player.matchesIds) return [];

  return player.matchesIds.map(id => state.matchesMap[id]);
}

export const statisticSelector = (state, id: number) => state.statisticsMap[id];
export const statisticsSelector = state =>
  state.statisticIds.map(id => state.statisticsMap[id]);
export const loginRouteSelector = state => state.loginRoute;
export const statusSelector = state => state.status;
export const matchesSelector = state =>
  state.matchesIds.map(id => state.matchesMap[id]);
export const matchSelector = (state, id: number) => state.matchesMap[id];
export const playerSelector = (state, id: number) => state.playersMap[id];
export const playersSelector = state => ({
  playersMap: state.playersMap,
  playerIds: state.playerIds,
});
export const userSelector = state => ({
  isLoggedIn: !!state.user,
  user: state.user,
});
