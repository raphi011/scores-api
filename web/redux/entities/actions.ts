import { ApiAction } from '../../redux/api/actions';
import { NewMatch, StatisticFilter } from '../../types';

import * as actionNames from '../actionNames';

export const loadMatchesAction = (
  groupId: number,
  after?: string,
): ApiAction => ({
  method: 'GET',
  params: { after },
  success: actionNames.RECEIVE_ENTITIES,
  successParams: {
    entityName: 'match',
    listOptions: {
      match: {
        key: groupId,
        mode: after ? 'append' : 'replace',
        name: 'group',
      },
    },
  },
  type: actionNames.API,
  url: `groups/${groupId}/matches`,
});

export const loadPlayerMatchesAction = (
  playerId: number,
  after?: string,
): ApiAction => ({
  method: 'GET',
  params: { after },
  success: actionNames.RECEIVE_ENTITIES,
  successParams: {
    entityName: 'match',
    listOptions: {
      match: {
        key: playerId,
        mode: after ? 'append' : 'replace',
        name: 'player',
      },
    },
  },
  type: actionNames.API,
  url: `players/${playerId}/matches`,
});

export const loadPlayerAction = (id: number): ApiAction => ({
  method: 'GET',
  success: actionNames.RECEIVE_ENTITIES,
  successParams: {
    entityName: 'player',
    listOptions: {
      player: {
        mode: 'replace',
        name: 'all',
      },
    },
  },
  type: actionNames.API,
  url: `players/${id}`,
});

export const loadPlayerTeamStatisticAction = (playerId: number): ApiAction => ({
  method: 'GET',
  success: actionNames.RECEIVE_ENTITIES,
  successParams: {
    assignId: true,
    entityName: 'statistic',
    listOptions: {
      statistic: {
        key: playerId,
        mode: 'replace',
        name: 'playerTeam',
      },
    },
  },
  type: actionNames.API,
  url: `players/${playerId}/teamStatistics`,
});

export const loadPlayerStatisticAction = (playerId: number): ApiAction => ({
  method: 'GET',
  success: actionNames.RECEIVE_ENTITIES,
  successParams: {
    assignId: true,
    entityName: 'statistic',
    listOptions: {
      statistic: {
        key: playerId,
        mode: 'replace',
        name: 'player',
      },
    },
  },
  type: actionNames.API,
  url: `players/${playerId}/playerStatistics`,
});

export const loadGroupAction = (groupId: number): ApiAction => ({
  method: 'GET',
  success: actionNames.RECEIVE_ENTITIES,
  successParams: {
    entityName: 'group',
    listOptions: {
      group: {
        key: groupId,
        name: 'all',
      },
      match: {
        key: groupId,
        name: 'group',
      },
      player: {
        key: groupId,
        name: 'group',
      },
    },
  },
  type: actionNames.API,
  url: `groups/${groupId}`,
});

export const loadGroupStatisticsAction = (
  groupId: number,
  filter: StatisticFilter,
): ApiAction => ({
  method: 'GET',
  params: { filter },
  success: actionNames.RECEIVE_ENTITIES,
  successParams: {
    assignId: true,
    entityName: 'statistic',
    listOptions: {
      statistic: {
        key: groupId,
        mode: 'replace',
        name: 'group',
      },
    },
  },
  type: actionNames.API,
  url: `groups/${groupId}/playerStatistics`,
});

export const loadMatchAction = (id: number): ApiAction => ({
  method: 'GET',
  success: actionNames.RECEIVE_ENTITIES,
  successParams: {
    entityName: 'match',
  },
  type: actionNames.API,
  url: `matches/${id}`,
});

export const createNewMatchAction = (match: NewMatch): ApiAction => ({
  body: JSON.stringify(match),
  method: 'POST',
  type: actionNames.API,
  url: `groups/${match.groupId}/matches`,
});

// export const deleteMatchAction = (match: Match): ApiAction => ({
//   method: 'DELETE',
//   success: actionNames.DELETE_ENTITIES,
//   successParams: {
//     entityName: 'match',
//     listOptions: {
//       match: {
//         names: ['all'],
//       },
//     },
//   },
//   successStatus: 'Match deleted',
//   type: actionNames.API,
//   url: `matches/${match.id}`,
// });

export const searchVolleynetPlayersAction = (filters: {
  fname: string;
  lname: string;
  bday: string;
}): ApiAction => ({
  method: 'GET',
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
  type: actionNames.API,
  url: 'volleynet/players/search',
});

export const loadTournamentAction = (tournamentId: string): ApiAction => ({
  method: 'GET',
  success: actionNames.RECEIVE_ENTITIES,
  successParams: {
    entityName: 'tournament',
  },
  type: actionNames.API,
  url: `volleynet/tournaments/${tournamentId}`,
});

export const loadTournamentsAction = (filters: {
  gender: string;
  league: string;
  season: string;
}): ApiAction => ({
  method: 'GET',
  params: filters,
  success: actionNames.RECEIVE_ENTITIES,
  successParams: {
    entityName: 'tournament',
    listOptions: {
      tournament: {
        key: filters.league,
        mode: 'replace',
        name: 'league',
      },
    },
  },
  type: actionNames.API,
  url: 'volleynet/tournaments',
});

export const loadLadderAction = (gender: 'M' | 'W'): ApiAction => ({
  method: 'GET',
  params: { gender },
  success: actionNames.RECEIVE_ENTITIES,
  successParams: {
    entityName: 'volleynetplayer',
    listOptions: {
      volleynetplayer: {
        key: gender,
        mode: 'replace',
        name: 'ladder',
      },
    },
  },
  type: actionNames.API,
  url: 'volleynet/ladder',
});

export const tournamentSignupAction = (form: {
  username: string;
  password: string;
  partnerId: number;
  tournamentId: number;
  partnerName: string;
  rememberMe: boolean;
}): ApiAction => ({
  body: JSON.stringify(form),
  method: 'POST',
  successStatus: '🎉 Successfully signed up',
  type: actionNames.API,
  url: 'volleynet/signup',
});
