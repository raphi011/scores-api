// @flow

import React from 'react';
import Router from 'next/router';
import fetch from 'isomorphic-unfetch';

import { CircularProgress } from 'material-ui/Progress';

import withAuth from '../containers/AuthContainer';
import TournamentList from '../components/volleynet/TournamentList';
import Layout from '../containers/LayoutContainer';

import type { Tournament } from '../types';

type State = {
  loading: boolean,
  tournaments: Array<Tournament>,
};

class Volleynet extends React.Component<null, State> {
  state = {
    loading: false,
    tournaments: [],
  };

  componentDidMount() {
    this.loadTournaments();
  }

  loadTournaments = async () => {
    const response = await fetch(
      'http://localhost:3000/api/volleynet/tournaments',
    );

    const tournaments = await response.json();

    this.setState({ tournaments });
  };

  onTournamentClick = (t: Tournament) => {
    Router.push('/tournament');
  };

  render() {
    const { tournaments } = this.state;

    const content = !tournaments.length ? (
      <CircularProgress />
    ) : (
      <TournamentList
        tournaments={tournaments}
        onTournamentClick={this.onTournamentClick}
      />
    );

    return <Layout title="Players">{content}</Layout>;
  }
}

export default withAuth(Volleynet);
