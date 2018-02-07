// @flow

import React from 'react';
import { withStyles } from 'material-ui/styles';
import Typography from 'material-ui/Typography';
import Tooltip from 'material-ui/Tooltip';
import Avatar from 'material-ui/Avatar';
import type { PlayerStatistic, Player } from '../types';

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
    fontSize: '50px',
  },
});

type Props = {
  statistic: PlayerStatistic,
  classes: Object,
  player: Player,
};

function PlayerView({ player, statistic, classes }: Props) {
  const avatar = statistic.player.profileImageUrl ? (
    <Avatar className={classes.avatar} src={statistic.player.profileImageUrl} />
  ) : (
    <Avatar className={classes.avatar}>{player.name.substring(0, 1)}</Avatar>
  );

  return (
    <div className={classes.profileHead}>
      {avatar}
      <Typography variant="headline">{player.name}</Typography>
      <Tooltip placement="top" id="tooltip-score" title="Played - Won">
        <Typography align="center" variant="display4">
          {statistic.gamesWon} - {statistic.gamesLost}
        </Typography>
      </Tooltip>
      <Typography align="center" variant="display3">
        {statistic.percentageWon}%
      </Typography>
    </div>
  );
}

export default withStyles(styles)(PlayerView);
