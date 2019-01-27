export enum EntityName {
  Player = 'player',
  Team = 'team',
  Tournament = 'tournament',
  User = 'user',
}

export interface User {
  id: number;
  email: string;
  role: string;
  profileImageUrl?: string;
  volleynetUserId: number;
  volleynetLogin: string;
}

export interface SearchPlayer {
  firstName: string;
  lastName: string;
  id: number;
  login: string;
  birthday: string;
}

export type Gender = 'M' | 'W';

export interface Player {
  id: number;

  createdAt: string;
  updatedAt: string;

  birthday?: string;
  club: string;
  countryUnion: string;
  firstName: string;
  gender: Gender;
  ladderRank: string;
  lastName: string;
  license: string;
  totalPoints: string;
}

export interface Team {
  tournamentId: number;
  player1: Player;
  player2: Player;

  createdAt: string;
  updatedAt: string;

  deregistered: boolean;
  prizeMoney: number;
  result: number;
  seed: number;
  totalPoints: number;
  wonPoints: number;
}

export enum TournamentStatus {
  Upcoming = 'upcoming',
  Done = 'done',
  Canceled = 'canceled',
}

export interface Tournament {
  id: number;

  createdAt: string;
  updatedAt: string;

  currentPoints: string;
  email: string;
  end: string;
  endRegistration?: string;
  entryLink: string;
  gender: Gender;
  htmlNotes: string;
  latitude: number;
  league: string;
  link: string;
  livescoringLink: string;
  location: string;
  longitude: number;
  maxPoints: number;
  maxTeams: number;
  minTeams: number;
  mode: string;
  name: string;
  organiser: string;
  phone: string;
  registrationOpen: boolean;
  signedupTeams: number;
  start: string;
  status: TournamentStatus;
  teams: Team[];
  website: string;
}

export interface ScrapeJob {
  start: string;
  end: string;
  sleep: string;
  lastDuration: number;
  job: {
    maxRuns: number;
    name: string;
    maxFailures: number;
    interval: string;
  };
  // errors:       []error TODO
  runs: number;
  state: number;
}

export type EntityType =
  | User
  | Player
  | Team
  | Tournament;

export interface Classes {
  [key: string]: string;
}
