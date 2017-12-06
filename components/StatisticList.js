import React from 'react';
import { withStyles } from 'material-ui/styles';
import Avatar from 'material-ui/Avatar';
import Table, {
  TableBody,
  TableCell,
  TableHead,
  TableRow,
} from 'material-ui/Table';
// import List, { ListItem, ListItemText } from 'material-ui/List';

const styles = theme => ({
  root: {
    width: '100%',
    background: theme.palette.background.paper,
  },
  listContainer: {
    display: 'flex',
    flexDirection: 'row',
    justifyContent: 'space-between',
  },
  rank: {
    fontSize: '20px',
    fontWeight: 'lighter',
  },
  rankContainer: {
    flexGrow: 0,
    flexShrink: 0,
  },
  denseCell: {
    paddingRight: theme.spacing.unit * 1.5,
  },
});

function LetterAvatar({ imageUrl, name }) {
  if (imageUrl) return <Avatar src={imageUrl} />;

  if (!name) return null;

  return <Avatar>{name[0]}</Avatar>;
}

function StatisticList({ statistics, classes }) {
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
          <TableRow key={s.playerId} >
            <TableCell className={classes.denseCell}>{i + 1}</TableCell>
            <TableCell className={classes.denseCell}>{s.name}</TableCell>
            <TableCell className={classes.denseCell} numeric>{s.percentageWon}%</TableCell>
            <TableCell className={classes.denseCell} padding="dense">
              {s.gamesWon} - {s.gamesLost}
            </TableCell>
          </TableRow>
        ))}
      </TableBody>
    </Table>
  );
}

export default withStyles(styles)(StatisticList);
