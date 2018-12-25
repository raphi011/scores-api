import React from 'react';

import ListItem from '@material-ui/core/ListItem';
import ListItemText from '@material-ui/core/ListItemText';

import { VolleynetPlayer } from '../../types';

interface Props {
  player: VolleynetPlayer;
}

const LadderItem = ({ player }: Props) => (
  <ListItem>
    <ListItemText
      primary={`${player.rank}. ${player.firstName} ${player.lastName}`}
      secondary={`${player.totalPoints} points • ${player.club} • ${
        player.countryUnion
      }`}
    />
  </ListItem>
);

export default LadderItem;
