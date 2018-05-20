// @flow

import React from 'react';
import ListItem from '@material-ui/core/ListItem';
import ListItemText from '@material-ui/core/ListItemText';
import { formatDate } from '../../utils/dateFormat';

import type { Tournament } from '../../types';

type Props = {
  tournament: Tournament,
  onClick: Tournament => void,
};

function dateString(tournament) {
  if (tournament.startDate === tournament.endDate) {
    return formatDate(tournament.startDate);
  }

  return `${formatDate(tournament.startDate)} - ${formatDate(
    tournament.endDate,
  )}`;
}

const TournamentListItem = ({ tournament, onClick }: Props) => (
  <ListItem button onClick={() => onClick(tournament)}>
    <ListItemText
      primary={tournament.name}
      secondary={dateString(tournament)}
    />
  </ListItem>
);

export default TournamentListItem;
