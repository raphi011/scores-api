import React from 'react';

import Card from '@material-ui/core/Card';
import List from '@material-ui/core/List';
import { createStyles, WithStyles, withStyles } from '@material-ui/core/styles';

import { Tournament } from '../../../types';
import TournamentListItem from './TournamentListItem';

const styles = createStyles({
  container: {
    marginBottom: '15px',
    padding: 0,
  },
  list: {
    padding: '0',
  },
});

interface Props extends WithStyles<typeof styles> {
  tournaments: Tournament[];
}

class TournamentList extends React.PureComponent<Props> {
  render() {
    const { tournaments = [], classes } = this.props;

    return (
      <Card classes={{ root: classes.container }}>
        <List className={classes.list}>
          {tournaments.map(t => (
            <TournamentListItem key={t.id} tournament={t} />
          ))}
        </List>
      </Card>
    );
  }
}

export default withStyles(styles)(TournamentList);
