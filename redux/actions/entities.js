// @flow

import * as actionNames from '../actionNames';
import type { ApiAction, Match, NewMatch, StatisticFilter } from '../../types';

export const loadMatchesAction = (after?: string): ApiAction => ({
  type: actionNames.API,
  method: 'GET',
  url: 'matches',
  params: { after },
  success: actionNames.RECEIVE_ENTITIES,
  successParams: {
    entityName: 'match',
    listName: 'all',
    mode: after ? 'append' : 'replace',
  },
});

export const loadPlayerMatchesAction = (
  playerId: number,
  after?: string,
): ApiAction => ({
  type: actionNames.API,
  method: 'GET',
  url: `playerMatches/${playerId}`,
  params: { after },
  success: actionNames.RECEIVE_ENTITIES,
  successParams: {
    entityName: 'match',
    listName: 'byPlayer',
    listKey: playerId,
    mode: after ? 'append' : 'replace',
  },
});

export const loadPlayerAction = (id: number): ApiAction => ({
  type: actionNames.API,
  method: 'GET',
  url: `players/${id}`,
  success: actionNames.RECEIVE_ENTITIES,
  successParams: {
    entityName: 'player',
    listName: 'all',
    mode: 'replace',
  },
});

export const loadPlayersAction = (): ApiAction => ({
  type: actionNames.API,
  method: 'GET',
  url: 'players',
  success: actionNames.RECEIVE_ENTITIES,
  successParams: {
    entityName: 'player',
    listName: 'all',
    mode: 'replace',
  },
});

export const loadPlayerTeamStatisticAction = (playerId: number): ApiAction => ({
  type: actionNames.API,
  method: 'GET',
  url: `playerTeamStatistics/${playerId}`,
  success: actionNames.RECEIVE_ENTITIES,
  successParams: {
    entityName: 'statistic',
    listName: 'byPlayerTeam',
    listKey: playerId,
    mode: 'replace',
    assignId: true,
  },
});

export const loadPlayerStatisticAction = (playerId: number): ApiAction => ({
  type: actionNames.API,
  method: 'GET',
  url: `statistics/${playerId}`,
  success: actionNames.RECEIVE_ENTITIES,
  successParams: {
    entityName: 'statistic',
    listName: 'byPlayer',
    listKey: playerId,
    mode: 'replace',
    assignId: true,
  },
});

export const loadStatisticsAction = (filter: StatisticFilter): ApiAction => ({
  type: actionNames.API,
  method: 'GET',
  url: 'statistics',
  params: { filter },
  success: actionNames.RECEIVE_ENTITIES,
  successParams: {
    entityName: 'statistic',
    listName: 'all',
    mode: 'replace',
    assignId: true,
  },
});

export const loadMatchAction = (id: number): ApiAction => ({
  type: actionNames.API,
  method: 'GET',
  url: `matches/${id}`,
  success: actionNames.RECEIVE_ENTITIES,
  successParams: {
    entityName: 'match',
  },
});

export const createNewMatchAction = (match: NewMatch): ApiAction => ({
  type: actionNames.API,
  method: 'POST',
  url: 'matches',
  body: JSON.stringify(match),
});

export const deleteMatchAction = (match: Match): ApiAction => ({
  type: actionNames.API,
  method: 'DELETE',
  url: `matches/${match.id}`,
  success: actionNames.DELETE_ENTITIES,
  successParams: {
    entityName: 'match',
    listNames: ['all'],
  },
  successStatus: 'Match deleted',
});
