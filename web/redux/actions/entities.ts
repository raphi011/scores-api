import * as actionNames from '../actionNames';
import { ApiAction, Match, NewMatch, StatisticFilter } from '../../types';

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
    listOptions: {
      match: {
        name: 'byGroup',
        key: groupId,
        mode: after ? 'append' : 'replace',
      },
    },
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
    listOptions: {
      match: {
        mode: after ? 'append' : 'replace',
        name: 'byPlayer',
        key: playerId,
      },
    },
  },
});

export const loadPlayerAction = (id: number): ApiAction => ({
  type: actionNames.API,
  method: 'GET',
  url: `players/${id}`,
  success: actionNames.RECEIVE_ENTITIES,
  successParams: {
    entityName: 'player',
    listOptions: {
      player: {
        name: 'all',
        mode: 'replace',
      },
    },
  },
});

export const loadPlayerTeamStatisticAction = (playerId: number): ApiAction => ({
  type: actionNames.API,
  method: 'GET',
  url: `players/${playerId}/teamStatistics`,
  success: actionNames.RECEIVE_ENTITIES,
  successParams: {
    entityName: 'statistic',
    assignId: true,
    listOptions: {
      statistic: {
        mode: 'replace',
        name: 'byPlayerTeam',
        key: playerId,
      },
    },
  },
});

export const loadPlayerStatisticAction = (playerId: number): ApiAction => ({
  type: actionNames.API,
  method: 'GET',
  url: `players/${playerId}/playerStatistics`,
  success: actionNames.RECEIVE_ENTITIES,
  successParams: {
    entityName: 'statistic',
    assignId: true,
    listOptions: {
      statistic: {
        name: 'byPlayer',
        key: playerId,
        mode: 'replace',
      },
    },
  },
});

export const loadGroupAction = (groupId: number): ApiAction => ({
  type: actionNames.API,
  method: 'GET',
  url: `groups/${groupId}`,
  success: actionNames.RECEIVE_ENTITIES,
  successParams: {
    entityName: 'group',
    listOptions: {
      group: {
        name: 'all',
        key: groupId,
      },
      player: {
        name: 'byGroup',
        key: groupId,
      },
      match: {
        name: 'byGroup',
        key: groupId,
      },
    },
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
    assignId: true,
    listOptions: {
      statistic: {
        name: 'byGroup',
        key: groupId,
        mode: 'replace',
      },
    },
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
  url: `groups/${match.groupId}/matches`,
  body: JSON.stringify(match),
});

export const deleteMatchAction = (match: Match): ApiAction => ({
  type: actionNames.API,
  method: 'DELETE',
  url: `matches/${match.id}`,
  success: actionNames.DELETE_ENTITIES,
  successParams: {
    entityName: 'match',
    listOptions: {
      match: {
        names: ['all'],
      },
    },
  },
  successStatus: 'Match deleted',
});

export const searchVolleynetPlayersAction = (filters: {
  fname: string;
  lname: string;
  bday: string;
}): ApiAction => ({
  type: actionNames.API,
  method: 'GET',
  url: 'volleynet/players/search',
  params: filters,
  success: actionNames.RECEIVE_ENTITIES,
  successParams: {
    entityName: 'volleynetplayer',
    listOptions: {
      volleynetplayer: {
        mode: 'replace',
        name: 'search',
      },
    },
  },
});

export const loadTournamentAction = (tournamentId: string): ApiAction => ({
  type: actionNames.API,
  method: 'GET',
  url: `volleynet/tournaments/${tournamentId}`,
  success: actionNames.RECEIVE_ENTITIES,
  successParams: {
    entityName: 'tournament',
  },
});

export const loadTournamentsAction = (filters: {
  gender: string;
  league: string;
  season: string;
}): ApiAction => ({
  type: actionNames.API,
  method: 'GET',
  url: 'volleynet/tournaments',
  params: filters,
  success: actionNames.RECEIVE_ENTITIES,
  successParams: {
    entityName: 'tournament',
    listOptions: {
      tournament: {
        name: 'byLeague',
        key: filters.league,
        mode: 'replace',
      },
    },
  },
});

export const tournamentSignupAction = (form: {
  username: string;
  password: string;
  partnerId: number;
  tournamentId: number;
  partnerName: string;
  rememberMe: boolean;
}): ApiAction => ({
  type: actionNames.API,
  method: 'POST',
  url: 'volleynet/signup',
  body: JSON.stringify(form),
  successStatus: '🎉 Successfully signed up',
});
