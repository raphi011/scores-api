// @flow

import React from 'react';
import { withStyles } from 'material-ui/styles';
// import Button from 'material-ui/Button';

import withAuth from '../containers/AuthContainer';
import Layout from '../containers/LayoutContainer';
import TournamentView from '../components/volleynet/TournamentView';

import type { FullTournament } from '../types';

const styles = theme => ({});

type Props = {
  tournamentId: string,
};

type State = {
  tournament: ?FullTournament,
};

class Tournament extends React.Component<Props, State> {
  static getParameters(query) {
    const { tournamentId } = query;

    return { tournamentId };
  }

  state = {
    tournament: null,
  };

  async componentDidMount() {
    const response = await fetch(
      'http://localhost:3000/api/volleynet/tournaments/1',
    );

    const tournament = await response.json();

    this.setState({ tournament });
  }

  render() {
    const { tournament } = this.state;

    return (
      <Layout title="New Match">
        <div style={{ paddingTop: '60px' }}>
          <TournamentView tournament={tournament} />
        </div>
      </Layout>
    );
  }
}

export default withAuth(withStyles(styles)(Tournament));
