import React from 'react';

import ListItem from '@material-ui/core/ListItem';
import ListItemText from '@material-ui/core/ListItemText';

import { Team } from '../../types';

interface Props {
  team: Team;
}

function rankOrSeed(team: Team) {
  if (!team.result) {
    return `${team.seed}.`;
  }
  switch (team.result) {
    case 1:
      return '1. 🥇';
    case 2:
      return '2. 🥈';
    case 3:
      return '3. 🥉';
    default:
      return `${team.result}.`;
  }
}

const TeamListItem = ({ team }: Props) => (
  <ListItem button>
    <ListItemText
      primary={
        <span>
          {`${rankOrSeed(team)} ${team.player1.firstName} ${
            team.player1.lastName
          } / 
          ${team.player2.firstName} ${team.player2.lastName}`}
        </span>
      }
      secondary={`${team.wonPoints || team.totalPoints} points`}
    />
  </ListItem>
);

export default TeamListItem;
