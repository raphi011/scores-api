import ListItem from '@material-ui/core/ListItem';
import ListItemText from '@material-ui/core/ListItemText';
import React from 'react';

import { Tournament } from '../../types';

interface IProps {
  tournament: Tournament;
  onClick: (Tournament) => void;
}

function buildSubtitle(tournament: Tournament) {
  let st = tournament.signedupTeams.toString();

  if (tournament.maxTeams >= 0) {
    st += ` / ${tournament.maxTeams} teams`;
  }

  st += ` • ${tournament.league}`;

  return st;
}

const TournamentListItem = ({ tournament, onClick }: IProps) => {
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
