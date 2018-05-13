// @flow

import React from 'react';
import Router from 'next/router';
import fetch from 'isomorphic-unfetch';

import Tabs, { Tab } from 'material-ui/Tabs';
import { CircularProgress } from 'material-ui/Progress';

import withAuth from '../../containers/AuthContainer';
import TournamentList from '../../components/volleynet/TournamentList';
import Layout from '../../containers/LayoutContainer';

import type { Tournament } from '../../types';

type State = {
  loading: boolean,
  tabOpen: number,
  tournaments: Array<Tournament>,
};

class Volleynet extends React.Component<null, State> {
  state = {
    loading: false,
    tournaments: [],
    tabOpen: 0,
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
    Router.push({ pathname: '/volleynet/tournament', query: { id: t.id } });
  };

  onTabClick = (event, tabOpen) => {
    this.setState({ tabOpen });
  };

  render() {
    const { tournaments, tabOpen } = this.state;

    const content = !tournaments.length ? (
      <CircularProgress />
    ) : (
      <TournamentList
        tournaments={tournaments}
        onTournamentClick={this.onTournamentClick}
      />
    );

    return (
      <Layout title="Players">
        <Tabs
          onChange={this.onTabClick}
          value={tabOpen}
          textColor="primary"
          fullWidth
        >
          <Tab label="Upcoming" />
          <Tab label="Past" />
          <Tab label="Played" />
        </Tabs>
        {tabOpen === 0 ? content : 'This is not done yet 💩'}
      </Layout>
    );
  }
}

export default withAuth(Volleynet);
