import React from 'react';
import ListItem from '@material-ui/core/ListItem';
import ListItemText from '@material-ui/core/ListItemText';
import { tournamentDateString } from '../../utils/tournament';

import { Tournament } from '../../types';

interface Props {
  tournament: Tournament;
  onClick: (Tournament) => void;
}

const TournamentListItem = ({ tournament, onClick }: Props) => (
  <ListItem button onClick={() => onClick(tournament)}>
    <ListItemText
      primary={tournament.name}
      secondary={`${tournamentDateString(tournament)} • ${
        tournament.signedupTeams
      } / ${tournament.maxTeams} teams`}
    />
  </ListItem>
);

export default TournamentListItem;
