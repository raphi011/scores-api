import React from 'react';
import Link from 'next/link';

import { withStyles } from '@material-ui/core/styles';
import Button from '@material-ui/core/Button';

import withAuth from '../../containers/AuthContainer';
import Layout from '../../containers/LayoutContainer';
import TournamentView from '../../components/volleynet/TournamentView';

import { loadTournamentAction } from '../../redux/actions/entities';
import { tournamentSelector } from '../../redux/reducers/entities';
import { userSelector } from '../../redux/reducers/auth';

import { Tournament, User } from '../../types';

const styles = () => ({});

interface Props {
  tournament?: Tournament;
  user: User;
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
    const { tournament, user } = this.props;

    return (
      <Layout title="Tournament">
        <Link prefetch href="/volleynet">
          <Button color="primary">Back</Button>
        </Link>
        <TournamentView tournament={tournament} user={user} />
      </Layout>
    );
  }
}

export default withAuth(withStyles(styles)(ShowTournament));
