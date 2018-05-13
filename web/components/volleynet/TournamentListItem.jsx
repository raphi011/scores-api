// @flow
import React from 'react';
// import { withStyles } from 'material-ui/styles';
import {
  ListItem,
  ListItemText,
  ListItemSecondaryAction,
} from 'material-ui/List';
import RegisterIcon from 'material-ui-icons/Create';
import IconButton from 'material-ui/IconButton';

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
      <IconButton aria-label="Register">
        <RegisterIcon />
      </IconButton>
    </ListItemSecondaryAction>
  </ListItem>
);

export default TournamentListItem;
