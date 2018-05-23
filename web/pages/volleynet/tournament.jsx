// @flow

import React from 'react';
import Link from 'next/link';

import { withStyles } from '@material-ui/core/styles';
import Button from '@material-ui/core/Button';

import withAuth from '../../containers/AuthContainer';
import Layout from '../../containers/LayoutContainer';
import TournamentView from '../../components/volleynet/TournamentView';

import { loadTournamentAction } from '../../redux/actions/entities';
import { tournamentSelector } from '../../redux/reducers/entities';

import type { FullTournament } from '../../types';

const styles = () => ({});

type Props = {
  tournamentId: number,
  tournament: ?FullTournament,
};

class Tournament extends React.Component<Props> {
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

    return { tournament };
  }

  render() {
    const { tournament } = this.props;

    return (
      <Layout title="Tournament">
        <Link prefetch href="/volleynet">
          <Button color="primary">Back</Button>
        </Link>
        <TournamentView tournament={tournament} />
      </Layout>
    );
  }
}

export default withAuth(withStyles(styles)(Tournament));
