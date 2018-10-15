import React from 'react';

import TournamentView from '../../components/volleynet/TournamentView';
import withAuth from '../../containers/AuthContainer';
import Layout from '../../containers/LayoutContainer';

import { loadTournamentAction } from '../../redux/actions/entities';
import { userSelector } from '../../redux/reducers/auth';
import { tournamentSelector } from '../../redux/reducers/entities';

import { Tournament, User } from '../../types';

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
    const { user } = userSelector(state);

    return { tournament, user };
  }

  render() {
    const { tournament, user } = this.props;

    return (
      <Layout title={{ text: 'Tournaments', href: '/volleynet' }}>
        <TournamentView tournament={tournament} user={user} />
      </Layout>
    );
  }
}

export default withAuth(ShowTournament);
