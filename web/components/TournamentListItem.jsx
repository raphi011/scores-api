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

// import Typography from 'material-ui/Typography';

import type { Tournament } from '../types';

// const itemStyles = () => ({
//   listContainer: {
//     display: 'flex',
//     flexDirection: 'row',
//     alignItems: 'center',
//     width: '100%',
//   },
//   team: { flex: '1 1 0' },
//   points: { fontWeight: 'lighter', flex: '2 2 0' },
// });

type Props = {
  //   match: Match,
  //   onMatchClick: Match => void,
  //   highlightPlayerId: number,
  //   classes: Object,
  tournament: Tournament,
};

const TournamentListItem = ({ tournament }: Props) => (
  <ListItem button>
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
