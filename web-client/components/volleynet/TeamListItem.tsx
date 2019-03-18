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

  return `${team.result}.`;
}

function medal(team: Team) {
  switch (team.result) {
    case 1:
      return '🥇';
    case 2:
      return '🥈';
    case 3:
      return '🥉';
    default:
      return '';
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
          ${team.player2.firstName} ${team.player2.lastName} ${medal(team)}`}
        </span>
      }
      secondary={`${team.wonPoints || team.totalPoints} points`}
    />
  </ListItem>
);

export default TeamListItem;
