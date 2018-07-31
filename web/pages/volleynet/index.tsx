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
import LeagueSelect from '../../components/LeagueSelect';

import { Tournament } from '../../types';

const leagues = ['AMATEUR TOUR', 'PRO TOUR', 'JUNIOR TEAM'];

interface State {
  tabOpen: number;
}

interface Props {
  tournaments: Tournament[];
  loadTournaments: () => void;
  league: string;
}

const thisYear = new Date().getFullYear().toString();

class Volleynet extends React.Component<Props, State> {
  static buildActions({ league }: Props) {
    return [
      loadTournamentsAction({
        gender: 'M',
        league,
        season: thisYear,
      }),
    ];
  }

  static mapDispatchToProps = {
    loadTournaments: loadTournamentsAction,
  };

  static getParameters(query) {
    let { league = 'AMATEUR TOUR' } = query;

    if (!leagues.includes(league)) {
      league = leagues[0];
    }

    return { league };
  }

  static mapStateToProps(state, { league }: Props) {
    const tournaments = tournamentsByLeagueSelector(state, league);

    return { tournaments };
  }

  state = {
    tabOpen: 0,
  };

  componentDidUpdate(prevProps) {
    const { loadTournaments, league } = this.props;

    if (league !== prevProps.league) {
      loadTournaments({ gender: 'M', league, season: thisYear });
    }
  }

  onTournamentClick = (t: Tournament) => {
    Router.push({ pathname: '/volleynet/tournament', query: { id: t.id } });
  };

  onLeagueChange = event => {
    const league = event.target.value;
    Router.push({
      pathname: '/volleynet',
      query: { league },
    });
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
      upcoming: tournaments.filter(t => t.status === 'upcoming'),
      past: tournaments.filter(t => t.status === 'done'),
      canceled: tournaments.filter(t => t.status === 'canceled'),
      played: [],
    };
  };

  render() {
    const { tabOpen } = this.state;
    const { league } = this.props;

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
          ts = tournaments.canceled;
          break;
        case 3:
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
      <Layout title={{ text: 'Volleynet', href: '' }}>
        <LeagueSelect selected={league} onChange={this.onLeagueChange} />
        <Tabs onChange={this.onTabClick} value={tabOpen} fullWidth>
          <Tab label="Upcoming" />
          <Tab label="Past" />
          <Tab label="Canceled" />
          <Tab label="Played" />
        </Tabs>
        {content}
      </Layout>
    );
  }
}

export default withAuth(Volleynet);
