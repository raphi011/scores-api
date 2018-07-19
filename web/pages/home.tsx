// @flow

import React from 'react';
import { withStyles } from '@material-ui/core/styles';
// import Router from 'next/router';

import Typography from '@material-ui/core/Typography';
import Paper from '@material-ui/core/Paper';
import withAuth from '../containers/AuthContainer';

import Layout from '../containers/LayoutContainer';
// import {
//   loadPlayerTeamStatisticAction,
//   loadPlayerStatisticAction,
//   loadPlayerMatchesAction,
// } from '../redux/actions/entities';
// import { multiApiAction } from '../redux/actions/api';
import { userSelector } from '../redux/reducers/auth';
import { playerSelector } from '../redux/reducers/entities';
import type { Player, Classes } from '../types';

const styles = theme => ({
  container: {
    margin: theme.spacing.unit,
  },
  paper: {
    margin: '10px 0',
    padding: theme.spacing.unit * 2,
  },
});

type Props = {
  player: Player,
  classes: Classes,
};

class Home extends React.Component<Props> {
  static getParameters(query) {
    const { id } = query;

    const playerId = Number.parseInt(id, 10);

    if (Number.isInteger(playerId)) {
      return { playerId };
    }

    return {};
  }

  //   static buildActions({ playerId, user }) {
  // return [multiApiAction([])];
  //   }

  static mapStateToProps(state) {
    const { user } = userSelector(state);

    const { playerId } = user;

    const player = playerSelector(state, playerId);

    return {
      player,
      user,
    };
  }

  //   static mapDispatchToProps = {
  //     loadMatches: loadPlayerMatchesAction,
  //   };

  render() {
    const { player, classes } = this.props;

    return (
      <Layout title="Home">
        <div className={classes.container}>
          <Typography variant="headline">Home</Typography>
          <Paper className={classes.paper}>
            <Typography variant="title">News</Typography>
          </Paper>
          <Paper className={classes.paper}>
            <Typography variant="title">Groups</Typography>
          </Paper>
          <Paper className={classes.paper}>
            <Typography variant="title">Volleynet</Typography>
            <Typography variant="subheading">Upcoming tournaments</Typography>
            <Typography variant="subheading">Past tournaments</Typography>
          </Paper>
        </div>
      </Layout>
    );
  }
}

export default withAuth(withStyles(styles)(Home));
