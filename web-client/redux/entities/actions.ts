import { ApiAction } from '../../redux/api/actions';

import { EntityName, EntityType } from '../../types';
import * as actionNames from '../actionNames';

export type EntityActionTypes = ReceiveEntitiesAction;

export interface ReceiveEntitiesAction {
  payload?: EntityType[];
  entityName: EntityName;
  assignId?: boolean;
  listOptions?: {
    [key in EntityName]?: {
      name: string;
      key?: string;
      mode?: 'replace' | 'append';
    }
  };
}

export const previousPartnersAction = (playerId: number): ApiAction => ({
  method: 'GET',
  success: actionNames.RECEIVE_ENTITIES,
  successParams: {
    entityName: EntityName.Player,
    listOptions: {
      [EntityName.Player]: {
        mode: 'replace',
        key: playerId.toString(),
        name: 'partners',
      },
    },
  },
  type: actionNames.API,
  url: `players/partners/${playerId}`,
});

export const searchPlayersAction = (filters: {
  fname: string;
  lname: string;
  gender: string;
}): ApiAction => ({
  method: 'GET',
  params: filters,
  success: actionNames.RECEIVE_ENTITIES,
  successParams: {
    entityName: EntityName.Player,
    listOptions: {
      [EntityName.Player]: {
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
    entityName: EntityName.Tournament,
  },
  type: actionNames.API,
  url: `tournaments/${tournamentId}`,
});

export const loadTournamentsAction = (filters: {
  genders: string[];
  leagues: string[];
  seasons: string;
}): ApiAction => ({
  method: 'GET',
  params: filters,
  success: actionNames.RECEIVE_ENTITIES,
  successParams: {
    entityName: EntityName.Tournament,
    listOptions: {
      [EntityName.Tournament]: {
        mode: 'replace',
        name: 'filter',
      },
    },
  },
  type: actionNames.API,
  url: 'tournaments',
});

export const loadLadderAction = (gender: string): ApiAction => ({
  method: 'GET',
  params: { gender },
  success: actionNames.RECEIVE_ENTITIES,
  successParams: {
    entityName: EntityName.Player,
    listOptions: {
      [EntityName.Player]: {
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
