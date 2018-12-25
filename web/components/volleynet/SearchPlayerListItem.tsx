import React from 'react';

import ListItem from '@material-ui/core/ListItem';
import ListItemText from '@material-ui/core/ListItemText';

import { VolleynetSearchPlayer } from '../../types';

interface Props {
  player: VolleynetSearchPlayer;
  onClick: (VolleynetSearchPlayer) => void;
}

const SearchPlayerListItem = ({ player, onClick }: Props) => (
  <ListItem button onClick={() => onClick(player)}>
    <ListItemText
      primary={`${player.firstName} ${player.lastName}`}
      secondary={player.birthday}
    />
  </ListItem>
);

export default SearchPlayerListItem;
