import { createSelector } from 'reselect';

import { denorm } from '../entitySchemas';

import { EntityName } from '../../types';
import { Store } from '../store';

const teamMap = (state: Store) => state.entities.team.values;
const tournamentMap = (state: Store) => state.entities.tournament.values;
const userMap = (state: Store) => state.entities.user.values;
const playerMap = (state: Store) => state.entities.player.values;

export const entityMapSelector = createSelector(
  playerMap,
  teamMap,
  tournamentMap,
  userMap,
  (
    player,
    team,
    tournament,
    user,
  ) => ({
    player,
    team,
    tournament,
    user,
  }),
);

export const allUsersSelector = (state: Store) =>
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

export const playerSelector = (state: Store, playerId: number) =>
  denorm(EntityName.Player, entityMapSelector(state), playerId);

export const tournamentsByLeagueSelector = (
  state: Store,
  leagues: string[],
) => {
  let tournaments = [];
  const entityMap = entityMapSelector(state);

  leagues.forEach(league => {
    tournaments = [
      ...tournaments,
      ...denorm(
        EntityName.Tournament,
        entityMap,
        state.entities.tournament.by.league[league] || [],
      ),
    ];
  });

  return tournaments;
};

export const tournamentSelector = (state: Store, tournamentId: number) =>
  denorm(EntityName.Tournament, entityMapSelector(state), tournamentId);

export const ladderVolleynetplayerSelector = (state: Store, gender: string) =>
  denorm(
    EntityName.Player,
    entityMapSelector(state),
    state.entities.player.by.ladder[gender],
  );

export const searchVolleynetplayerSelector = (state: Store) =>
  denorm(
    EntityName.Player,
    entityMapSelector(state),
    state.entities.player.list.search,
  );
