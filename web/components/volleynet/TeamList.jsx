// @flow

import React from 'react';
import { withStyles } from 'material-ui/styles';
import List from 'material-ui/List';

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

class TeamList extends React.PureComponent<Props> {
  render() {
    const { teams = [], classes } = this.props;

    return (
      <List className={classes.root}>
        {teams.map(t => (
          <TeamListItem key={t.player1.id + t.player2.id} team={t} />
        ))}
      </List>
    );
  }
}

export default withStyles(styles)(TeamList);
