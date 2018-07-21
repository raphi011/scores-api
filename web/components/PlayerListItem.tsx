import React from 'react';

import Avatar from '@material-ui/core/Avatar';
import ListItem from '@material-ui/core/ListItem';
import ListItemText from '@material-ui/core/ListItemText';
import Badge from '@material-ui/core/Badge';

import { Player } from '../types';

interface Props {
  onClick: (Event) => void;
  player: Player;
  playerNr: number;
}

function PlayerListItem({ player, onClick, playerNr }: Props) {
  let color;
  let team;

  switch (playerNr) {
    case 1:
    case 2:
      color = 'primary';
      team = 1;
      break;
    case 3:
    case 4:
      color = 'secondary';
      team = 2;
      break;
    default:
      color = '';
      team = null;
  }

  let avatar = player.profileImageUrl ? (
    <Avatar src={player.profileImageUrl} />
  ) : (
    <Avatar>{player.name.substring(0, 1)}</Avatar>
  );

  if (playerNr) {
    avatar = (
      <Badge badgeContent={team} color={color}>
        {avatar}
      </Badge>
    );
  }

  return (
    <ListItem onClick={onClick} button>
      {avatar}
      <ListItemText inset primary={player.name} />
    </ListItem>
  );
}

export default PlayerListItem;
