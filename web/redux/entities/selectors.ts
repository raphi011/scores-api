import { createSelector } from 'reselect';

import { denorm } from '../entitySchemas';

import { IStore } from '../store';

const groupMap = state => state.entities.group.values;
const playerMap = state => state.entities.player.values;
const teamMap = state => state.entities.team.values;
const matchMap = state => state.entities.match.values;
const statisticMap = state => state.entities.statistic.values;
const tournamentMap = state => state.entities.tournament.values;
const userMap = state => state.entities.user.values;
const volleynetplayerMap = state => state.entities.volleynetplayer.values;

export const entityMapSelector = createSelector(
  groupMap,
  playerMap,
  teamMap,
  matchMap,
  statisticMap,
  tournamentMap,
  userMap,
  volleynetplayerMap,
  (group, player, team, match, statistic, tournament, volleynetplayer) => ({
    group,
    match,
    player,
    statistic,
    team,
    tournament,
    volleynetplayer,
  }),
);

export const allUsersSelector = (state: IStore) =>
  state.entities.user.all.length
    ? denorm('user', entityMapSelector(state), state.entities.user.all)
    : [];

export const allPlayersSelector = (state: IStore) =>
  state.entities.player.all.length
    ? denorm('player', entityMapSelector(state), state.entities.player.all)
    : [];

export const groupSelector = (state: IStore, groupId: number) =>
  denorm('group', entityMapSelector(state), groupId);

export const groupPlayersSelector = (state: IStore, groupId: number) =>
  (state.entities.player.byGroup[groupId] || []).length
    ? denorm(
        'player',
        entityMapSelector(state),
        state.entities.player.byGroup[groupId],
      )
    : [];

export const matchSelector = (state: IStore, id: number) =>
  denorm('match', entityMapSelector(state), id);

const allMatchIdsSelector = state =>
  state.entities.match.all.length ? state.entities.match.all : [];

export const allMatchesSelector = createSelector(
  allMatchIdsSelector,
  entityMapSelector,
  (ids, entities) => denorm('match', entities, ids),
);

export const playerSelector = (state: IStore, playerId: number) =>
  denorm('player', entityMapSelector(state), playerId);

export const matchesByGroupSelector = (state: IStore, groupId: number) =>
  (state.entities.match.byGroup[groupId] || []).length
    ? denorm(
        'match',
        entityMapSelector(state),
        state.entities.match.byGroup[groupId],
      )
    : [];

export const matchesByPlayerSelector = (state: IStore, playerId: number) =>
  (state.entities.match.byPlayer[playerId] || []).length
    ? denorm(
        'match',
        entityMapSelector(state),
        state.entities.match.byPlayer[playerId],
      )
    : [];

export const allStatisticSelector = (state: IStore) =>
  state.entities.statistic.all.length
    ? denorm(
        'statistic',
        entityMapSelector(state),
        state.entities.statistic.all,
      )
    : [];

export const statisticByPlayerTeamSelector = (
  state: IStore,
  playerId: number,
) =>
  (state.entities.statistic.byPlayerTeam[playerId] || []).length
    ? denorm(
        'statistic',
        entityMapSelector(state),
        state.entities.statistic.byPlayerTeam[playerId],
      )
    : [];

export const statisticByGroupSelector = (state: IStore, groupId: number) =>
  (state.entities.statistic.byGroup[groupId] || []).length
    ? denorm(
        'statistic',
        entityMapSelector(state),
        state.entities.statistic.byGroup[groupId],
      )
    : [];

export const statisticByPlayerSelector = (state: IStore, playerId: number) =>
  (state.entities.statistic.byPlayer[playerId] || []).length
    ? denorm(
        'statistic',
        entityMapSelector(state),
        state.entities.statistic.byPlayer[playerId][0],
      )
    : null;

export const tournamentsByLeagueSelector = (
  state: IStore,
  leagues: string[],
) => {
  let tournaments = [];

  leagues.forEach(league => {
    tournaments = [
      ...tournaments,
      ...((state.entities.tournament.byLeague[league] || []).length
        ? denorm(
            'tournament',
            entityMapSelector(state),
            state.entities.tournament.byLeague[league],
          )
        : []),
    ];
  });

  return tournaments;
};

export const tournamentSelector = (state: IStore, tournamentId: number) =>
  denorm('tournament', entityMapSelector(state), tournamentId);

export const ladderVolleynetplayerSelector = (state: IStore, gender: string) =>
  (state.entities.volleynetplayer.ladder[gender] || []).length
    ? denorm(
        'volleynetplayer',
        entityMapSelector(state),
        state.entities.volleynetplayer.ladder[gender],
      )
    : [];

export const searchVolleynetplayerSelector = (state: IStore) =>
  (state.entities.volleynetplayer.search || []).length
    ? denorm(
        'volleynetplayer',
        entityMapSelector(state),
        state.entities.volleynetplayer.search,
      )
    : [];
