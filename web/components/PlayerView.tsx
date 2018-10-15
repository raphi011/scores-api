import Avatar from '@material-ui/core/Avatar';
import { createStyles, withStyles } from '@material-ui/core/styles';
import Tooltip from '@material-ui/core/Tooltip';
import Typography from '@material-ui/core/Typography';
import React from 'react';
import { Player, PlayerStatistic } from '../types';

const styles = createStyles({
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

interface Props {
  statistic: PlayerStatistic;
  player: Player;
  classes: any;
}

function PlayerView({ player, statistic, classes }: Props) {
  const avatar = statistic.player.profileImageUrl ? (
    <Avatar className={classes.avatar} src={statistic.player.profileImageUrl} />
  ) : (
    <Avatar className={classes.avatar}>{player.name.substring(0, 1)}</Avatar>
  );

  return (
    <div className={classes.profileHead}>
      {avatar}
      <Typography variant="h5">{player.name}</Typography>
      <Typography variant="subtitle1">{statistic.rank}</Typography>
      <Tooltip placement="top" id="tooltip-score" title="Played - Won">
        <Typography align="center" variant="h1">
          {statistic.gamesWon} - {statistic.gamesLost}
        </Typography>
      </Tooltip>
      <Typography align="center" variant="h2">
        {statistic.percentageWon}%
      </Typography>
    </div>
  );
}

export default withStyles(styles)(PlayerView);
