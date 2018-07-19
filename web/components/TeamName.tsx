// @flow

import React from 'react';
import { withStyles } from '@material-ui/core/styles';

import type { Team } from '../types';

const styles = theme => ({
  normal: { padding: '2px 4px' },
  highlighted: {
    backgroundColor: theme.palette.primary.light,
    padding: '2px 4px',
  },
});

type Props = {
  team: Team,
  highlightPlayerId: number,
  classes: { [string]: string },
};

const TeamName = ({ team, highlightPlayerId, classes }: Props) => {
  if (team.name) {
    const highlight =
      team.player1Id === highlightPlayerId ||
      team.player2Id === highlightPlayerId;

    return (
      <span className={highlight ? classes.highlighted : classes.normal}>
        {team.name}
      </span>
    );
  }

  const highlightPlayer1 = team.player1Id === highlightPlayerId;
  const highlightPlayer2 = team.player2Id === highlightPlayerId;

  return (
    <span>
      <span className={highlightPlayer1 ? classes.highlighted : classes.normal}>
        {team.player1.name}
      </span>
      <br />
      <span className={highlightPlayer2 ? classes.highlighted : classes.normal}>
        {team.player2.name}
      </span>
    </span>
  );
};

export default withStyles(styles)(TeamName);
