import React from 'react';
import Link from 'next/link';

import { withStyles, createStyles, Theme } from '@material-ui/core/styles';
import Button from '@material-ui/core/Button';

import withAuth from '../../containers/AuthContainer';
import Layout from '../../containers/LayoutContainer';
import TournamentView from '../../components/volleynet/TournamentView';

import { loadTournamentAction } from '../../redux/actions/entities';
import { tournamentSelector } from '../../redux/reducers/entities';
import { userSelector } from '../../redux/reducers/auth';

import { Tournament, User } from '../../types';

const styles = (theme: Theme) =>
  createStyles({
    backButton: {
      position: 'absolute',
      right: theme.spacing.unit,
    },
    tournamentContainer: {
      paddingTop: theme.spacing.unit * 3,
    },
  });

interface Props {
  tournament?: Tournament;
  user: User;
  classes: any;
}

class ShowTournament extends React.Component<Props> {
  static getParameters(query) {
    const { id } = query;

    const tournamentId = Number.parseInt(id, 10);

    return { tournamentId };
  }

  static buildActions({ tournamentId }) {
    return [loadTournamentAction(tournamentId)];
  }

  static mapStateToProps(state, { tournamentId }) {
    const tournament = tournamentSelector(state, tournamentId);
    const user = userSelector(state);

    return { tournament, user };
  }

  render() {
    const { tournament, user, classes } = this.props;

    return (
      <Layout title="Tournament">
        <div className={classes.backButton}>
          <Link prefetch href="/volleynet">
            <Button variant="outlined" color="primary">
              Back
            </Button>
          </Link>
        </div>
        <div className={classes.tournamentContainer}>
          <TournamentView tournament={tournament} user={user} />
        </div>
      </Layout>
    );
  }
}

export default withAuth(withStyles(styles)(ShowTournament));
