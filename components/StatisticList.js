import React from 'react';
import { withStyles } from 'material-ui/styles';
import Table, {
  TableBody,
  TableCell,
  TableHead,
  TableRow,
} from 'material-ui/Table';

import StatisticListRow from './StatisticListRow';

const styles = theme => ({
  denseCell: {
    paddingRight: theme.spacing.unit * 1.5,
  },
});

function StatisticList({ statistics, onPlayerClick, classes }) {
  return (
    <Table>
      <TableHead>
        <TableRow>
          <TableCell className={classes.denseCell}>#</TableCell>
          <TableCell className={classes.denseCell}>Player</TableCell>
          <TableCell className={classes.denseCell} numeric>Won</TableCell>
          <TableCell className={classes.denseCell}>Games</TableCell>
        </TableRow>
      </TableHead>
      <TableBody>
        {statistics.map((s, i) => (
          <StatisticListRow key={s.playerId} onPlayerClick={onPlayerClick} statistic={s} rank={i+1} />
        ))}
      </TableBody>
    </Table>
  );
}

export default withStyles(styles)(StatisticList);
