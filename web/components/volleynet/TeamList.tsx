import List from '@material-ui/core/List';
import { createStyles, withStyles } from '@material-ui/core/styles';
import React from 'react';

import TeamListItem from './TeamListItem';

import { VolleynetTeam } from '../../types';

const styles = createStyles({
  root: {
    width: '100%',
  },
});

interface Props {
  teams: VolleynetTeam[];
  emptyMessage?: string;
  classes: any;
}

const sortByRankOrSeed = (a, b) => (a.rank || a.seed) - (b.rank || b.seed);

class TeamList extends React.PureComponent<Props> {
  render() {
    const { teams = [], emptyMessage = "", classes } = this.props;

    if (!teams || !teams.length) {
      return emptyMessage;
    }

    return (
      <List className={classes.root}>
        {teams.sort(sortByRankOrSeed).map(t => (
          <TeamListItem key={t.player1.id + t.player2.id} team={t} />
        ))}
      </List>
    );
  }
}

export default withStyles(styles)(TeamList);
