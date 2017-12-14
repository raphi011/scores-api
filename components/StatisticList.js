// @flow

import React from 'react';
import { withStyles } from 'material-ui/styles';
import Table, {
  TableBody,
  TableCell,
  TableHead,
  TableRow,
} from 'material-ui/Table';
import type { Statistic } from '../types';

import StatisticListRow from './StatisticListRow';

const styles = theme => ({
  denseCell: {
    paddingRight: theme.spacing.unit * 1.5,
  },
});

type Props = {
  statistics: Array<Statistic>,
  onPlayerClick: number => void,
  classes: Object
}

function StatisticList({ statistics, onPlayerClick, classes }: Props) {
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
