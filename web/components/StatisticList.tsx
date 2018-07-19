import React from 'react';
import { withStyles } from '@material-ui/core/styles';
import Table from '@material-ui/core/Table';
import TableBody from '@material-ui/core/TableBody';
import TableCell from '@material-ui/core/TableCell';
import TableHead from '@material-ui/core/TableHead';
import TableRow from '@material-ui/core/TableRow';
import { PlayerStatistic, Classes } from '../types';

import StatisticListRow from './StatisticListRow';

const styles = theme => ({
  denseCell: {
    paddingRight: theme.spacing.unit * 1.5,
  },
});

interface Props {
  statistics: PlayerStatistic[];
  onPlayerClick: (number) => void;
  classes: Classes;
}

function StatisticList({ statistics, onPlayerClick, classes }: Props) {
  return (
    <Table>
      <TableHead>
        <TableRow>
          <TableCell className={classes.denseCell}>#</TableCell>
          <TableCell className={classes.denseCell}>Player</TableCell>
          <TableCell className={classes.denseCell} numeric>
            Won
          </TableCell>
          <TableCell className={classes.denseCell}>Games</TableCell>
        </TableRow>
      </TableHead>
      <TableBody>
        {statistics.map((s, i) => (
          <StatisticListRow
            key={s.playerId}
            onPlayerClick={onPlayerClick}
            statistic={s}
            rank={i + 1}
          />
        ))}
      </TableBody>
    </Table>
  );
}

export default withStyles(styles)(StatisticList);
