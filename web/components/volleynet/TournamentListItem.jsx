// @flow
import React from 'react';
// import { withStyles } from '@material-ui/core/styles';
import {
  ListItem,
  ListItemText,
  ListItemSecondaryAction,
} from '@material-ui/core/List';
import GotoIcon from '@material-ui/icons/ArrowForward';

import type { Tournament } from '../../types';

type Props = {
  tournament: Tournament,
  onClick: Tournament => void,
};

const TournamentListItem = ({ tournament, onClick }: Props) => (
  <ListItem button onClick={() => onClick(tournament)}>
    <ListItemText
      primary={tournament.name}
      secondary={`${tournament.startDate} - ${tournament.league}`}
    />
    <ListItemSecondaryAction>
      <GotoIcon />
    </ListItemSecondaryAction>
  </ListItem>
);

export default TournamentListItem;
