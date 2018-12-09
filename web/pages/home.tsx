import React from 'react';

import Paper from '@material-ui/core/Paper';
import { createStyles, Theme, withStyles } from '@material-ui/core/styles';
import Typography from '@material-ui/core/Typography';
import withAuth from '../containers/AuthContainer';

import Layout from '../containers/LayoutContainer';
import { userSelector } from '../redux/auth/selectors';
import { playerSelector } from '../redux/entities/selectors';
import { Player } from '../types';

const styles = (theme: Theme) =>
  createStyles({
    container: {
      margin: theme.spacing.unit,
    },
    paper: {
      margin: '10px 0',
      padding: theme.spacing.unit * 2,
    },
  });

type Props = {
  player: Player;
  classes: any;
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

  static mapStateToProps(state) {
    const { user, isLoggedIn } = userSelector(state);

    if (!isLoggedIn) {
      return {};
    }

    const { playerId } = user;

    const player = playerSelector(state, playerId);

    return {
      player,
      user,
    };
  }

  render() {
    const { classes } = this.props;

    return (
      <Layout title={{ text: 'Home', href: '' }}>
        <div className={classes.container}>
          <Typography variant="h5">Home</Typography>
          <Paper className={classes.paper}>
            <Typography variant="h6">News</Typography>
          </Paper>
          <Paper className={classes.paper}>
            <Typography variant="h6">Volleynet</Typography>
            <Typography variant="subtitle1">Upcoming tournaments</Typography>
            <Typography variant="subtitle1">Past tournaments</Typography>
          </Paper>
        </div>
      </Layout>
    );
  }
}

export default withAuth(withStyles(styles)(Home));
