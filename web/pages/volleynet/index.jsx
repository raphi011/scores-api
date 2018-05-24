// @flow

import React from 'react';
import Router from 'next/router';

import Tabs from '@material-ui/core/Tabs';
import Tab from '@material-ui/core/Tab';

import withAuth from '../../containers/AuthContainer';
import TournamentList from '../../components/volleynet/TournamentList';
import CenteredLoading from '../../components/CenteredLoading';
import Layout from '../../containers/LayoutContainer';
import { loadTournamentsAction } from '../../redux/actions/entities';
import { tournamentsByLeagueSelector } from '../../redux/reducers/entities';

import type { Tournament } from '../../types';

const leagues = {
  Amateur: 'AMATEUR TOUR',
};

type State = {
  tabOpen: number,
};

type Props = {
  tournaments: Array<Tournament>,
};

class Volleynet extends React.Component<Props, State> {
  static buildActions() {
    return [loadTournamentsAction({ gender: 'M', league: leagues.Amateur })];
  }

  static mapStateToProps(state) {
    const tournaments = tournamentsByLeagueSelector(state, leagues.Amateur);

    return { tournaments };
  }

  state = {
    tabOpen: 0,
  };

  onTournamentClick = (t: Tournament) => {
    Router.push({ pathname: '/volleynet/tournament', query: { id: t.id } });
  };

  onTabClick = (event, tabOpen) => {
    this.setState({ tabOpen });
  };

  orderTournaments = () => {
    let { tournaments } = this.props;

    if (!tournaments) {
      return null;
    }

    tournaments = tournaments.sort(
      (a, b) => new Date(a.start) - new Date(b.start),
    );

    return {
      upcoming: tournaments.filter(t => new Date(t.start) >= Date.now()),
      past: tournaments.filter(t => new Date(t.start) < Date.now()),
      played: [],
    };
  };

  render() {
    const { tabOpen } = this.state;
    const tournaments = this.orderTournaments();

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
      <Layout title="Volleynet">
        <Tabs onChange={this.onTabClick} value={tabOpen} fullWidth>
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
