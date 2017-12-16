// @flow

import React from 'react';
import { withStyles } from 'material-ui/styles';
import Typography from 'material-ui/Typography';
import Tooltip from 'material-ui/Tooltip';
import Avatar from 'material-ui/Avatar';
import type { /* Match, */ Statistic, Player } from '../types';

const styles = () => ({
  profileHead: {
    display: 'flex',
    flexDirection: 'column',
    alignItems: 'center',
    justify: 'center',
    marginTop: 20,
  },
  avatar: {
    margin: 10,
    width: 128,
    height: 128,
  },
});

type Props = {
  // matches: Array<Match>,
  statistic: Statistic,
  classes: Object,
  player: Player,
};

function PlayerView({ /* matches, */ player, statistic, classes }: Props) {
  return (
    <div className={classes.profileHead}>
      <Avatar className={classes.avatar} src={statistic.profileImage} />
      <Typography type="headline">{player.name}</Typography>
      <Tooltip placement="top" id="tooltip-score" title="Played - Won">
        <Typography align="center" type="display4">
          {statistic.gamesWon} - {statistic.gamesLost}
        </Typography>
      </Tooltip>
      <Typography align="center" type="display3">
        {statistic.percentageWon}%
      </Typography>
      {/* <MatchList matches={matches} /> */}
    </div>
  );
}

export default withStyles(styles)(PlayerView);
