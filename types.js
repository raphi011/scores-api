// @flow

export type Player = {
  id: number,
  name: string,
};

export type Team = {
  name: string,
  id: number,
  player1: Player,
  player2: Player,
};

export type Match = {
  id: number,
  scoreTeam1: number,
  scoreTeam2: number,
  createdAt: string,
  team1: Team,
  team2: Team,
};

export type User = {
  id: number,
  email: string,
  profileImageUrl: string,
};

export type Statistic = {
  playerId: number,
  name: string,
  profileImage: string,
  gamesWon: number,
  gamesLost: number,
  percentageWon: number,
};

export type Action = {
  type: string,
};

export type ApiAction = {
  type: 'API',
  method: string,
  url: string,
  success: string,
  params: Object,
  isServer: boolean,
  req?: Object, // todo
  res?: Object, // todo
  headers: Object,

  error?: string,
  body?: string,
  successStatus?: string,
  successParams?: Object,
};
