import React from 'react';

import ListItem from '@material-ui/core/ListItem';
import ListItemText from '@material-ui/core/ListItemText';

import { SearchPlayer } from '../../types';

interface Props {
  player: SearchPlayer;
  onClick: (player: SearchPlayer) => void;
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
