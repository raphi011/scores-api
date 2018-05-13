// @flow
import React from 'react';

import { ListItem, ListItemText } from 'material-ui/List';

import type { VolleynetTeam } from '../../types';

type Props = {
  team: VolleynetTeam,
};

const TeamListItem = ({ team }: Props) => (
  <ListItem button>
    <ListItemText
      primary={
        <span>
          {`${team.player1.firstName} ${team.player1.lastName} / 
          ${team.player2.firstName} ${team.player2.lastName}`}
        </span>
      }
      secondary={team.totalPoints}
    />
  </ListItem>
);

export default TeamListItem;
