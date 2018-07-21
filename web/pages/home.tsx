import React from 'react';
import { withStyles, Theme, createStyles } from '@material-ui/core/styles';

import Typography from '@material-ui/core/Typography';
import Paper from '@material-ui/core/Paper';
import withAuth from '../containers/AuthContainer';

import Layout from '../containers/LayoutContainer';
import { userSelector } from '../redux/reducers/auth';
import { playerSelector } from '../redux/reducers/entities';
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

interface Props {
  player: Player;
  classes: any;
}

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
    const { user } = userSelector(state);

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
