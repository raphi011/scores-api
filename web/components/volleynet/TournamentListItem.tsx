import React from 'react';
import ListItem from '@material-ui/core/ListItem';
import ListItemText from '@material-ui/core/ListItemText';
import { tournamentDateString } from '../../utils/tournament';

import { Tournament } from '../../types';

interface Props {
  tournament: Tournament;
  onClick: (Tournament) => void;
}

function buildSubtitle(tournament: Tournament) {
  let st = tournamentDateString(tournament);

  if (tournament.maxTeams >= 0) {
    st += ` • ${tournament.signedupTeams} / ${tournament.maxTeams} teams`;
  }

  return st;
}

const TournamentListItem = ({ tournament, onClick }: Props) => {
  let primary: string | JSX.Element = tournament.name;

  if (tournament.status === 'canceled') {
    primary = (
      <span>
        <span style={{ textDecoration: 'line-through' }}>{primary}</span>{' '}
        (canceled)
      </span>
    );
  }

  return (
    <ListItem button onClick={() => onClick(tournament)}>
      <ListItemText primary={primary} secondary={buildSubtitle(tournament)} />
    </ListItem>
  );
};

export default TournamentListItem;
