import React, { SyntheticEvent } from 'react';

import Avatar from '@material-ui/core/Avatar';
import Badge from '@material-ui/core/Badge';
import ListItem from '@material-ui/core/ListItem';
import ListItemText from '@material-ui/core/ListItemText';

import { Player } from '../types';

interface Props {
  player: Player;
  playerNr: number;

  onClick: (event: SyntheticEvent<{}>) => void;
}

function PlayerListItem({ player, onClick, playerNr }: Props) {
  let color: 'primary' | 'secondary' | 'default';
  let team: number;

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
      color = 'default';
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
