// @flow

import React from 'react';
import Router from 'next/router';
import fetch from 'isomorphic-unfetch';

import Tabs from '@material-ui/core/Tabs';
import Tab from '@material-ui/core/Tab';

import withAuth from '../../containers/AuthContainer';
import TournamentList from '../../components/volleynet/TournamentList';
import CenteredLoading from '../../components/CenteredLoading';
import Layout from '../../containers/LayoutContainer';
import { buildUrl } from '../../api';

import type { Tournament } from '../../types';

type State = {
  loading: boolean,
  tabOpen: number,
  tournaments: {
    upcoming: Array<Tournament>,
    past: Array<Tournament>,
    played: Array<Tournament>,
  },
};

class Volleynet extends React.Component<null, State> {
  state = {
    loading: false,
    tournaments: null,
    tabOpen: 0,
  };

  componentDidMount() {
    this.loadTournaments();
  }

  loadTournaments = async () => {
    const response = await fetch(buildUrl('volleynet/tournaments'));

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

    let content = <CenteredLoading />;
    let ts = [];

    if (tournaments) {
      switch (tabOpen) {
        case 0:
          ts = tournaments.upcoming;
          break;
        case 1:
          ts = tournaments.past;
          break;
        case 2:
          ts = tournaments.played;
          break;
        default: // this shouldn't happen
      }
      content = (
        <TournamentList
          tournaments={ts}
          onTournamentClick={this.onTournamentClick}
        />
      );
    }

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
        {content}
      </Layout>
    );
  }
}

export default withAuth(Volleynet);
