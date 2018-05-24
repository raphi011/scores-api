// @flow

export type EntityName =
  | 'player'
  | 'team'
  | 'match'
  | 'group'
  | 'tournament'
  | 'volleynetplayer';

export type Player = {
  id: number,
  name: string,
  userId: ?number,
  profileImageUrl: ?string,
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
  groupId: number,
  scoreTeam1: number,
  scoreTeam2: number,
  createdAt: string,
  team1: Team,
  team2: Team,
};

export type NewMatch = {
  groupId: number,
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
  playerId: number,
  volleynetUserId: number,
  volleynetLogin: string,
};

export type Statistic = {
  played: number,
  gamesWon: number,
  gamesLost: number,
  pointsWon: number,
  pointsLost: number,
  percentageWon: number,
  rank: string,
};

export type TeamStatistic = {
  player2Id: number,
  player2Id: number,
  team: Team,
} & Statistic;

export type StatisticFilter = 'today' | 'month' | 'thisyear' | 'all';

export type PlayerStatistic = {
  playerId: number,
  player: Player,
} & Statistic;

export type Group = {
  id: number,
  name: string,
  imageUrl: string,
  players: Player[],
  matches: Match[],
};

export type VolleynetSearchPlayer = {
  firstName: string,
  lastName: string,
  id: number,
  login: string,
  birthday: string,
};

export type VolleynetPlayer = {
  firstName: string,
  lastName: string,
  id: number,
  birthday: string,
  totalPoints: string,
  countryUnion: string,
  license: string,
  login: string,
};

export type VolleynetTeam = {
  player1: VolleynetPlayer,
  player2: VolleynetPlayer,
  totalPoints: string,
  seedOrRank: string,
  wonPoints: string,
  prizeMoney: string,
};

export type Gender = 'M' | 'W';

export type Tournament = {
  id: number,
  gender: Gender,
  registrationOpen: boolean,
  start: string,
  end: string,
  name: string,
  league: string,
  link: string,
  entryLink: string,
  teams: Array<VolleynetTeam>,
  status: string,
  location: string,
  htmlNotes: string,
  mode: string,
  maxTeams: number,
  minTeams: string,
  maxPoints: string,
  endRegistration: string,
  organiser: string,
  phone: string,
  email: string,
  web: string,
  currentPoints: string,
  livescoringLink: string,
  latitude: number,
  longitude: number,
};

export type GenericStatistic = PlayerStatistic | TeamStatistic;

export type EntityType = Group | Player | Team | Match | GenericStatistic;

export type Action = {
  type: string,
};

export type Classes = { [string]: string };

export type ApiAction =
  | {
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
    }
  | Action;

export type ApiActions =
  | {
      type: 'API_MULTI',
      actions: Array<ApiAction>,
      req?: Object, // todo
      res?: Object, // todo
      isServer?: boolean,
    }
  | Action;
