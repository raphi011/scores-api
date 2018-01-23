import * as actionNames from '../actionNames';
import type { ApiAction, Match, StatisticFilter } from '../../types';

export const loadMatchesAction = (): ApiAction => ({
  type: actionNames.API,
  method: 'GET',
  url: 'matches',
  success: actionNames.RECEIVE_MATCHES,
});

export const loadPlayerMatchesAction = (playerId: number): ApiAction => ({
  type: actionNames.API,
  method: 'GET',
  url: `playerMatches/${playerId}`,
  success: actionNames.RECEIVE_PLAYER_MATCHES,
  successParams: { playerId },
});

export const loadPlayerAction = (id: number): ApiAction => ({
  type: actionNames.API,
  method: 'GET',
  url: `players/${id}`,
  success: actionNames.RECEIVE_PLAYER,
  successParams: { id },
});

export const loadStatisticAction = (playerId: number): ApiAction => ({
  type: actionNames.API,
  method: 'GET',
  url: `statistics/${playerId}`,
  success: actionNames.RECEIVE_STATISTIC,
  successParams: { playerId },
});

export const loadStatisticsAction = (filter: StatisticFilter): ApiAction => ({
  type: actionNames.API,
  method: 'GET',
  url: 'statistics',
  params: { filter },
  success: actionNames.RECEIVE_STATISTICS,
});

export const loadPlayersAction = (): ApiAction => ({
  type: actionNames.API,
  method: 'GET',
  url: 'players',
  success: actionNames.RECEIVE_PLAYERS,
});

export const loadMatchAction = (id: number): ApiAction => ({
  type: actionNames.API,
  method: 'GET',
  url: `matches/${id}`,
  success: actionNames.RECEIVE_MATCH,
  successParams: { id },
});

export const createNewMatchAction = (match: Match): ApiAction => ({
  type: actionNames.API,
  method: 'POST',
  url: 'matches',
  body: JSON.stringify(match),
  successStatus: 'New Match created',
});

export const deleteMatchAction = (match: Match): ApiAction => ({
  type: actionNames.API,
  method: 'DELETE',
  url: `matches/${match.id}`,
  success: actionNames.REMOVE_MATCH,
  successParams: { id: match.id },
  successStatus: 'Match deleted',
});

export const userOrLoginRouteAction = (): ApiAction => ({
  type: actionNames.API,
  method: 'GET',
  url: 'userOrLoginRoute',
  success: actionNames.SET_USER_OR_LOGINROUTE,
});

export const loggedInAction = username => ({
  type: actionNames.LOGIN,
  username,
});

export const logoutAction = (): ApiAction => ({
  type: actionNames.API,
  method: 'POST',
  url: 'logout',
  success: actionNames.LOGGEDOUT,
  successStatus: 'Logged out',
});

export const loggedOutAction = () => ({
  type: actionNames.LOGGEDOUT,
});

export const setStatusAction = status => ({
  type: actionNames.SET_STATUS,
  status,
});

export const clearStatusAction = () => ({
  type: actionNames.CLEAR_STATUS,
});
