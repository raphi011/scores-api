import Router from 'next/router';
import React from 'react';

import Tab from '@material-ui/core/Tab';
import Tabs from '@material-ui/core/Tabs';

import CenteredLoading from '../../components/CenteredLoading';
import LeagueSelect from '../../components/LeagueSelect';
import TournamentList from '../../components/volleynet/TournamentList';
import withAuth from '../../containers/AuthContainer';
import Layout from '../../containers/LayoutContainer';
import { loadTournamentsAction } from '../../redux/actions/entities';
import { tournamentsByLeagueSelector } from '../../redux/reducers/entities';

import { Tournament } from '../../types';

const leagues = ['AMATEUR TOUR', 'PRO TOUR', 'JUNIOR TEAM'];

interface IState {
  tabOpen: number;
}

interface IProps {
  tournaments: Tournament[];
  loadTournaments: (
    filters: { gender: string; league: string; season: string },
  ) => void;
  league: string;
  classes: any;
}

const thisYear = new Date().getFullYear().toString();

class Volleynet extends React.Component<IProps, IState> {

  static mapDispatchToProps = {
    loadTournaments: loadTournamentsAction,
  };
  static buildActions({ league }: IProps) {
    return [
      loadTournamentsAction({
        gender: 'M',
        league,
        season: thisYear,
      }),
    ];
  }

  static getParameters(query) {
    let { league = 'AMATEUR TOUR' } = query;

    if (!leagues.includes(league)) {
      league = leagues[0];
    }

    return { league };
  }

  static mapStateToProps(state, { league }: IProps) {
    const tournaments = tournamentsByLeagueSelector(state, league);

    return { tournaments };
  }

  static sortAscending = (a, b) =>
    new Date(a.start).getTime() - new Date(b.start).getTime();

  static sortDescending = (a, b) =>
    new Date(b.start).getTime() - new Date(a.start).getTime();

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

  onTabClick = (_, tabOpen) => {
    this.setState({ tabOpen });
  };

  orderTournaments = () => {
    const { tournaments } = this.props;

    if (!tournaments) {
      return null;
    }

    return {
      upcoming: tournaments
        .filter(
          t =>
            t.status === 'upcoming' ||
            (t.status === 'canceled' && new Date(t.end) >= new Date()),
        )
        .sort(Volleynet.sortAscending),

      past: tournaments
        .filter(
          t =>
            t.status === 'done' ||
            (t.status === 'canceled' && new Date(t.end) < new Date()),
        )
        .sort(Volleynet.sortDescending),
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
        </Tabs>
        {content}
      </Layout>
    );
  }
}

export default withAuth(Volleynet);
