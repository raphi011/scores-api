import React from 'react';

import ListItem from '@material-ui/core/ListItem';
import ListItemText from '@material-ui/core/ListItemText';

import { VolleynetTeam } from '../../types';

interface Props {
  team: VolleynetTeam;
}

function rankOrSeed(team: VolleynetTeam) {
  if (!team.rank) {
    return `${team.seed}.`;
  }
  switch (team.rank) {
    case 1:
      return '1. 🥇';
    case 2:
      return '2. 🥈';
    case 3:
      return '3. 🥉';
    default:
      return `${team.rank}.`;
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
