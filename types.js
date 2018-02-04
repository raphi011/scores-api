// @flow

export type Player = {
  id: number,
  name: string,
  userId: ?number,
  profileImageUrl: ?string,
  matchesIds: Array<number>,
};

export type Team = {
  name: string,
  id: number,
  player1: Player,
  player1Id: number,
  player2: Player,
  player2Id: number,
};

export type Match = {
  id: number,
  scoreTeam1: number,
  scoreTeam2: number,
  createdAt: string,
  team1: Team,
  team2: Team,
};

export type NewMatch = {
  player1Id: number,
  player2Id: number,
  player3Id: number,
  player4Id: number,
  scoreTeam1: number,
  scoreTeam2: number,
  targetScore: number,
};

export type User = {
  id: number,
  email: string,
  profileImageUrl: ?string,
};

export type Statistic = {
  played: number,
  gamesWon: number,
  gamesLost: number,
  pointsWon: number,
  pointsLost: number,
  percentageWon: number,
};

export type TeamStatistic = {
  player2Id: number,
  player2Id: number,
  team: Team,
} & Statistic;

export type StatisticFilter = 'week' | 'month' | 'quarter' | 'year';

export type PlayerStatistic = {
  playerId: number,
  player: Player,
} & Statistic;

export type Action = {
  type: string,
};

export type Classes = { [string]: string };

export type ApiAction = {
  type: 'API',
  method: string,
  url: string,
  success?: string,
  isServer?: boolean,
  params?: Object,
  req?: Object, // todo
  res?: Object, // todo
  headers?: Object,

  error?: string,
  body?: string,
  successStatus?: string,
  successParams?: Object,
};
