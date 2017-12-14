// @flow

export type Player = {
  ID: number,
  Name: string,
};

export type Team = {
  Name: string,
  ID: number,
  Player1: Player,
  Player2: Player,
};

export type Match = {
  ID: number,
  ScoreTeam1: number,
  ScoreTeam2: number,
  CreatedAt: string,
  Team1: Team,
  Team2: Team,
};

export type User = {
  ID: number,
  Name: string,
  ProfileImageURL: string,
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
