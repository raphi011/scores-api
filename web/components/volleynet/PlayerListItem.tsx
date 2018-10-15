import React from 'react';

import ListItem from '@material-ui/core/ListItem';
import ListItemText from '@material-ui/core/ListItemText';

import { VolleynetSearchPlayer } from '../../types';

interface IProps {
  player: VolleynetSearchPlayer;
  onClick: (VolleynetSearchPlayer) => void;
}

const PlayerListItem = ({ player, onClick }: IProps) => (
  <ListItem button onClick={() => onClick(player)}>
    <ListItemText
      primary={`${player.firstName} ${player.lastName}`}
      secondary={player.birthday}
    />
  </ListItem>
);

export default PlayerListItem;
