import React from 'react';

import ListItem from '@material-ui/core/ListItem';
import ListItemText from '@material-ui/core/ListItemText';

import { VolleynetPlayer } from '../../types';

interface IProps {
  player: VolleynetPlayer;
}

const LadderItem = ({ player }: IProps) => (
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
