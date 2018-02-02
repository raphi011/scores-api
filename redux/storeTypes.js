// @flow

import type { User, Player, Match, PlayerStatistic } from '../types';

export type Store = {
  user: ?User,
  loginRoute: ?string,
  status: ?string,
  playersMap: { [string]: Player },
  statisticsMap: { [string]: PlayerStatistic },
  matchesMap: { [string]: Match },
  matchesIds: Array<number>,
  playerIds: Array<number>,
  statisticIds: Array<number>,
};
