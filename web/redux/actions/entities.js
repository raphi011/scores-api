// @flow

import * as actionNames from '../actionNames';
import type { ApiAction, Match, NewMatch, StatisticFilter } from '../../types';

export const loadMatchesAction = (
  groupId: number,
  after?: string,
): ApiAction => ({
  type: actionNames.API,
  method: 'GET',
  url: `groups/${groupId}/matches`,
  params: { after },
  success: actionNames.RECEIVE_ENTITIES,
  successParams: {
    entityName: 'match',
    listName: 'byGroup',
    listKey: groupId,
    mode: after ? 'append' : 'replace',
  },
});

export const loadPlayerMatchesAction = (
  playerId: number,
  after?: string,
): ApiAction => ({
  type: actionNames.API,
  method: 'GET',
  url: `players/${playerId}/matches`,
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

// export const loadPlayersAction = (groupId: number): ApiAction => ({
//   type: actionNames.API,
//   method: 'GET',
//   url: `groups/${groupId}/players`,
//   success: actionNames.RECEIVE_ENTITIES,
//   successParams: {
//     entityName: 'player',
//     listName: 'group',
//     listKey: groupId,
//     mode: 'replace',
//   },
// });

export const loadPlayerTeamStatisticAction = (playerId: number): ApiAction => ({
  type: actionNames.API,
  method: 'GET',
  url: `players/${playerId}/teamStatistics`,
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
  url: `players/${playerId}/playerStatistics`,
  success: actionNames.RECEIVE_ENTITIES,
  successParams: {
    entityName: 'statistic',
    listName: 'byPlayer',
    listKey: playerId,
    mode: 'replace',
    assignId: true,
  },
});

export const loadGroupAction = (groupId: number): ApiAction => ({
  type: actionNames.API,
  method: 'GET',
  url: `groups/${groupId}`,
  successParams: {
    entityName: 'group',
    listName: 'all',
    listKey: groupId,
  },
});

export const loadGroupStatisticsAction = (
  groupId: number,
  filter: StatisticFilter,
): ApiAction => ({
  type: actionNames.API,
  method: 'GET',
  url: `groups/${groupId}/playerStatistics`,
  params: { filter },
  success: actionNames.RECEIVE_ENTITIES,
  successParams: {
    entityName: 'statistic',
    listName: 'byGroup',
    listKey: groupId,
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
  url: `/groups/${match.groupId}/matches`,
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
