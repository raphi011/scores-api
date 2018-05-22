import { formatDate } from './dateFormat';

import type { Tournament } from '../types';

export function tournamentDateString(tournament: Tournament) {
  if (tournament.start === tournament.end) {
    return formatDate(tournament.start);
  }

  return `${formatDate(tournament.start)} - ${formatDate(tournament.end)}`;
}
