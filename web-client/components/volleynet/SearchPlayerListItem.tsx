import React from 'react';

import ListItem from '@material-ui/core/ListItem';
import ListItemText from '@material-ui/core/ListItemText';

import { Player } from '../../types';

interface Props {
  player: Player;
  onClick: (player: Player) => void;
}

const SearchPlayerListItem = ({ player, onClick }: Props) => (
  <ListItem button onClick={() => onClick(player)}>
    <ListItemText
      primary={`${player.firstName} ${player.lastName} (#${player.ladderRank})`}
      secondary={player.birthday ? new Date(player.birthday).getFullYear() : ''}
    />
  </ListItem>
);

export default SearchPlayerListItem;
