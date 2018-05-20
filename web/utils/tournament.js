import { formatDate } from './dateFormat';

import type { Tournament } from '../types';

export function tournamentDateString(tournament: Tournament) {
  if (tournament.startDate === tournament.endDate) {
    return formatDate(tournament.startDate);
  }

  return `${formatDate(tournament.startDate)} - ${formatDate(
    tournament.endDate,
  )}`;
}
