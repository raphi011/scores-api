// @flow

import React from 'react';
import { withStyles } from '@material-ui/core/styles';
import {
  TableCell,
  TableRow,
} from '@material-ui/core/Table';
import type { PlayerStatistic } from '../types';

const styles = theme => ({
  denseCell: {
    paddingRight: theme.spacing.unit * 1.5,
  },
});

type Props = {
  onPlayerClick: number => void,
  statistic: PlayerStatistic,
  rank: number,
  classes: Object,
}

class StatisticListRow extends React.PureComponent<Props> {
  onPlayerClick = () => {
    const { onPlayerClick, statistic } = this.props;

    onPlayerClick(statistic.playerId);
  }

  render() {
    const { statistic, rank, classes } = this.props;

    return (
      <TableRow hover key={statistic.playerId} onClick={this.onPlayerClick} >
        <TableCell className={classes.denseCell}>{rank}</TableCell>
        <TableCell className={classes.denseCell}>{statistic.player.name}</TableCell>
        <TableCell className={classes.denseCell} numeric>{statistic.percentageWon}%</TableCell>
        <TableCell className={classes.denseCell} padding="dense">
          {statistic.gamesWon} - {statistic.gamesLost}
        </TableCell>
      </TableRow>
    );
  }
}

export default withStyles(styles)(StatisticListRow);
