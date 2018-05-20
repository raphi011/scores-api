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
  id: string,
  tournament: ?FullTournament,
};

class Tournament extends React.Component<Props> {
  static getParameters(query) {
    const { id } = query;

    return { id };
  }

  static buildActions({ id }) {
    return [loadTournamentAction(id)];
  }

  static mapStateToProps(state, { id }) {
    const tournament = tournamentSelector(state, id);

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
