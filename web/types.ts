import { IParams } from './api';
export type EntityName =
  | 'player'
  | 'team'
  | 'match'
  | 'group'
  | 'tournament'
  | 'statistic'
  | 'volleynetplayer';

export interface Player {
  id: number;
  name: string;
  userId?: number;
  profileImageUrl?: string;
  groups: Group[];
}

export interface Team {
  name: string;
  id: number;
  player1: Player;
  player1Id: number;
  player2: Player;
  player2Id: number;
}

export interface Match {
  id: number;
  groupId: number;
  scoreTeam1: number;
  scoreTeam2: number;
  createdAt: string;
  team1: Team;
  team2: Team;
}

export interface NewMatch {
  groupId: number;
  player1Id: number;
  player2Id: number;
  player3Id: number;
  player4Id: number;
  scoreTeam1: number;
  scoreTeam2: number;
  targetScore: number;
}

export interface User {
  id: number;
  email: string;
  profileImageUrl?: string;
  playerId: number;
  volleynetUserId: number;
  volleynetLogin: string;
}

export interface Statistic {
  played: number;
  gamesWon: number;
  gamesLost: number;
  pointsWon: number;
  pointsLost: number;
  percentageWon: number;
  rank: string;
}

export interface TeamStatistic extends Statistic {
  player1Id: number;
  player2Id: number;
  team: Team;
}

export type StatisticFilter = 'today' | 'month' | 'thisyear' | 'all';

export interface PlayerStatistic extends Statistic {
  playerId: number;
  player: Player;
}

export interface Group {
  id: number;
  name: string;
  imageUrl: string;
  players: Player[];
  matches: Match[];
}

export interface VolleynetSearchPlayer {
  firstName: string;
  lastName: string;
  id: number;
  login: string;
  birthday: string;
}

export type Gender = 'M' | 'W';

export interface VolleynetPlayer {
  id: number;
  firstName: string;
  lastName: string;
  login: string;
  birthday: string;
  gender: Gender;
  totalPoints: string;
  rank: string;
  club: string;
  countryUnion: string;
  license: string;
}

export interface VolleynetTeam {
  tournamentId: number;
  player1: VolleynetPlayer;
  player2: VolleynetPlayer;
  totalPoints: string;
  seed: number;
  rank: number;
  wonPoints: string;
  prizeMoney: string;
  deregistered: boolean;
}

export interface Tournament {
  id: number;
  updatedAt: string;
  gender: Gender;
  registrationOpen: boolean;
  start: string;
  end: string;
  name: string;
  league: string;
  link: string;
  entryLink: string;
  teams: VolleynetTeam[];
  status: string;
  location: string;
  htmlNotes: string;
  mode: string;
  signedupTeams: number;
  maxTeams: number;
  minTeams: string;
  maxPoints: string;
  endRegistration: string;
  organiser: string;
  phone: string;
  email: string;
  web: string;
  currentPoints: string;
  livescoringLink: string;
  latitude: number;
  longitude: number;
}

export type GenericStatistic = PlayerStatistic | TeamStatistic;

export type EntityType = Group | Player | Team | Match | GenericStatistic;

export interface Action {
  type: string;
}

export interface Classes {
  [key: string]: string;
}

export interface ApiAction extends Action {
  type: 'API';
  method: string;
  url: string;
  success?: string;
  isServer?: boolean;
  params?: IParams;
  req?: Object; // todo
  res?: Object; // todo
  headers?: { [key: string]: string };

  error?: string;
  body?: string;
  successStatus?: string;
  successParams?: Object;
}

export interface ApiActions extends Action {
  type: 'API_MULTI';
  actions: Array<ApiAction>;
  req?: Object; // todo
  res?: Object; // todo
  isServer?: boolean;
}
