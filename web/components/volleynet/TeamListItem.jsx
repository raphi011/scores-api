// @flow
import React from 'react';

import ListItem from '@material-ui/core/ListItem';
import ListItemText from '@material-ui/core/ListItemText';

import type { VolleynetTeam } from '../../types';

type Props = {
  team: VolleynetTeam,
};

const TeamListItem = ({ team }: Props) => (
  <ListItem button>
    <ListItemText
      primary={
        <span>
          {`${team.seedOrRank}. ${team.player1.firstName} ${
            team.player1.lastName
          } / 
          ${team.player2.firstName} ${team.player2.lastName}`}
        </span>
      }
      secondary={`${team.totalPoints || team.wonPoints} points`}
    />
  </ListItem>
);

export default TeamListItem;
