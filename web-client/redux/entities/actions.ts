import { ApiAction } from '../../redux/api/actions';

import * as actionNames from '../actionNames';

export const searchVolleynetPlayersAction = (filters: {
  fname: string;
  lname: string;
  bday: string;
}): ApiAction => ({
  method: 'GET',
  params: filters,
  success: actionNames.RECEIVE_ENTITIES,
  successParams: {
    entityName: 'player',
    listOptions: {
      volleynetplayer: {
        mode: 'replace',
        name: 'search',
      },
    },
  },
  type: actionNames.API,
  url: 'players/search',
});

export const loadTournamentAction = (tournamentId: string): ApiAction => ({
  method: 'GET',
  success: actionNames.RECEIVE_ENTITIES,
  successParams: {
    entityName: 'tournament',
  },
  type: actionNames.API,
  url: `tournaments/${tournamentId}`,
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
  url: 'tournaments',
});

export const loadLadderAction = (gender: 'M' | 'W'): ApiAction => ({
  method: 'GET',
  params: { gender },
  success: actionNames.RECEIVE_ENTITIES,
  successParams: {
    entityName: 'player',
    listOptions: {
      volleynetplayer: {
        key: gender,
        mode: 'replace',
        name: 'ladder',
      },
    },
  },
  type: actionNames.API,
  url: 'ladder',
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
  url: 'signup',
});
