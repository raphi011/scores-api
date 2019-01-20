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
  playerId: number;
  player: Player;
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

export interface Team {
  tournamentId: number;
  player1: Player;
  player2: Player;
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
  teams: Team[];
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
