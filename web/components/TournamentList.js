// @flow

import React from 'react';
import { withStyles } from 'material-ui/styles';
import List from 'material-ui/List';

import TournamentListItem from './TournamentListItem';

import type { Tournament } from '../types';

const styles = () => ({
  root: {
    width: '100%',
  },
});

type Props = {
  tournaments: Array<Tournament>,
  classes: Object,
};

class TournamentList extends React.PureComponent<Props> {
  render() {
    const { tournaments = [], classes } = this.props;

    return (
      <List className={classes.root}>
        {tournaments.map(t => <TournamentListItem key={t.id} tournament={t} />)}
      </List>
    );
  }
}

export default withStyles(styles)(TournamentList);
