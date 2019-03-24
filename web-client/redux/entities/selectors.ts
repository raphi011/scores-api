import { createSelector } from 'reselect';

import { denorm } from '../entitySchemas';

import { EntityName, User } from '../../types';
import { Store } from '../store';

const tournamentMap = (state: Store) => state.entities.tournament.values;
const userMap = (state: Store) => state.entities.user.values;
const playerMap = (state: Store) => state.entities.player.values;

export const entityMapSelector = createSelector(
  playerMap,
  tournamentMap,
  userMap,
  (player, tournament, user) => ({
    player,
    tournament,
    user,
  }),
);

export const allUsersSelector = (state: Store): User[] =>
  denorm(
    EntityName.User,
    entityMapSelector(state),
    state.entities.user.list.all,
  );

export const allPlayersSelector = (state: Store) =>
  denorm(
    EntityName.Player,
    entityMapSelector(state),
    state.entities.player.list.all,
  );

export const previousPartnersSelector = (state: Store, playerId: number) =>
  denorm(
    EntityName.Player,
    entityMapSelector(state),
    state.entities.player.by ? state.entities.player.by.partners[playerId] : [],
  );

export const playerSelector = (state: Store, playerId: number) =>
  denorm(EntityName.Player, entityMapSelector(state), playerId.toString());

export const filteredTournamentsSelector = (state: Store) =>
  denorm(
    EntityName.Tournament,
    entityMapSelector(state),
    state.entities.tournament.list.filter || [],
  );

export const tournamentSelector = (state: Store, tournamentId: string) =>
  denorm(EntityName.Tournament, entityMapSelector(state), tournamentId);

export const ladderVolleynetplayerSelector = (state: Store, gender: string) =>
  denorm(
    EntityName.Player,
    entityMapSelector(state),
    state.entities.player.by ? state.entities.player.by.ladder[gender] : [],
  );

export const searchVolleynetplayerSelector = (state: Store) =>
  denorm(
    EntityName.Player,
    entityMapSelector(state),
    state.entities.player.list.search,
  );
