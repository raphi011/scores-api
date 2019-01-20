import React from 'react';

import List from '@material-ui/core/List';
import { createStyles, WithStyles, withStyles } from '@material-ui/core/styles';

import TeamListItem from './TeamListItem';

import { Team } from '../../types';

const styles = createStyles({
  root: {
    width: '100%',
  },
});

interface Props extends WithStyles<typeof styles> {
  teams: Team[];
  emptyMessage?: string;
}

const sortByResultOrSeed = (a: Team, b: Team) =>
  (a.result || a.seed) - (b.result || b.seed);

class TeamList extends React.PureComponent<Props> {
  render() {
    const { teams = [], emptyMessage = '', classes } = this.props;

    if (!teams || !teams.length) {
      return emptyMessage;
    }

    return (
      <List className={classes.root}>
        {teams.sort(sortByResultOrSeed).map(t => (
          <TeamListItem key={t.player1.id + t.player2.id} team={t} />
        ))}
      </List>
    );
  }
}

export default withStyles(styles)(TeamList);
