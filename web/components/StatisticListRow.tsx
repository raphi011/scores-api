import { createStyles, Theme, withStyles } from '@material-ui/core/styles';
import TableCell from '@material-ui/core/TableCell';
import TableRow from '@material-ui/core/TableRow';
import React from 'react';

import { PlayerStatistic } from '../types';

const styles = (theme: Theme) =>
  createStyles({
    denseCell: {
      paddingRight: theme.spacing.unit * 1.5,
    },
  });

interface IProps {
  onPlayerClick: (playerId: number) => void;
  statistic: PlayerStatistic;
  rank: number;
  classes: any;
}

class StatisticListRow extends React.PureComponent<IProps> {
  onPlayerClick = () => {
    const { onPlayerClick, statistic } = this.props;

    onPlayerClick(statistic.playerId);
  };

  render() {
    const { statistic, rank, classes } = this.props;

    return (
      <TableRow hover key={statistic.playerId} onClick={this.onPlayerClick}>
        <TableCell className={classes.denseCell}>{rank}</TableCell>
        <TableCell className={classes.denseCell}>
          {statistic.player.name}
        </TableCell>
        <TableCell className={classes.denseCell} numeric>
          {statistic.percentageWon}%
        </TableCell>
        <TableCell className={classes.denseCell} padding="dense">
          {statistic.gamesWon} - {statistic.gamesLost}
        </TableCell>
      </TableRow>
    );
  }
}

export default withStyles(styles)(StatisticListRow);
