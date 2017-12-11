import React from 'react';
import { withStyles } from 'material-ui/styles';
import {
  TableCell,
  TableRow,
} from 'material-ui/Table';

const styles = theme => ({
  denseCell: {
    paddingRight: theme.spacing.unit * 1.5,
  },
});

class StatisticListRow extends React.PureComponent {
  onPlayerClick = () => {
    const { onPlayerClick, statistic } = this.props;

    onPlayerClick(statistic.playerId);
  }

  render() {
    const { statistic, rank, classes } = this.props;

    return (
      <TableRow key={statistic.playerId} onClick={this.onPlayerClick} >
        <TableCell className={classes.denseCell}>{rank}</TableCell>
        <TableCell className={classes.denseCell}>{statistic.name}</TableCell>
        <TableCell className={classes.denseCell} numeric>{statistic.percentageWon}%</TableCell>
        <TableCell className={classes.denseCell} padding="dense">
          {statistic.gamesWon} - {statistic.gamesLost}
        </TableCell>
      </TableRow>
    );
  }
}

export default withStyles(styles)(StatisticListRow);
