import { formatDate } from './dateFormat';

import type { Tournament } from '../types';

export function tournamentDateString(tournament: Tournament) {
  if (tournament.start === tournament.end) {
    return formatDate(tournament.start);
  }

  return `${formatDate(tournament.start)} - ${formatDate(tournament.end)}`;
}

export function isSignedup(tournament: Tournament, userId: string): boolean {
  if (!tournament || !tournament.teams) {
    return false;
  }

  return tournament.some(
    t => t.player1.login === userId || t.player2.id === userId,
  );
}
