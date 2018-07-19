import React from 'react';

import ListItem from '@material-ui/core/ListItem';
import ListItemText from '@material-ui/core/ListItemText';

import { VolleynetPlayer } from '../../types';

interface Props {
  player: VolleynetPlayer;
  onClick: (VolleynetPlayer) => void;
}

const PlayerListItem = ({ player, onClick }: Props) => (
  <ListItem button onClick={() => onClick(player)}>
    <ListItemText
      primary={`${player.firstName} ${player.lastName}`}
      secondary={player.birthday}
    />
  </ListItem>
);

export default PlayerListItem;
