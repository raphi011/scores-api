// @flow

import React from 'react';
import { withStyles } from '@material-ui/core/styles';
import List from '@material-ui/core/List';

import TeamListItem from './TeamListItem';

import type { VolleynetTeam } from '../../types';

const styles = () => ({
  root: {
    width: '100%',
  },
});

type Props = {
  teams: Array<VolleynetTeam>,
  classes: Object,
};

const sortByRankOrSeed = (a, b) => (a.rank || a.seed) - (b.rank || b.seed);

class TeamList extends React.PureComponent<Props> {
  render() {
    const { teams = [], classes } = this.props;

    return (
      <List className={classes.root}>
        {teams
          .sort(sortByRankOrSeed)
          .map(t => (
            <TeamListItem key={t.player1.id + t.player2.id} team={t} />
          ))}
      </List>
    );
  }
}

export default withStyles(styles)(TeamList);
