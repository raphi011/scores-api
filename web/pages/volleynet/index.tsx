import Router from 'next/router';
import React from 'react';

import { createStyles, withStyles } from '@material-ui/core/styles';
import CenteredLoading from '../../components/CenteredLoading';
import DayHeader from '../../components/DayHeader';
import GroupedList from '../../components/GroupedList';
import LeagueSelect from '../../components/volleynet/LeagueSelect';
import TournamentList from '../../components/volleynet/TournamentList';
import withAuth from '../../containers/AuthContainer';
import Layout from '../../containers/LayoutContainer';
import { loadTournamentsAction } from '../../redux/actions/entities';
import { tournamentsByLeagueSelector } from '../../redux/reducers/entities';
import * as ArrayUtils from '../../utils/array';

import { Tournament } from '../../types';

const defaultLeagues = ['AMATEUR TOUR', 'PRO TOUR', 'JUNIOR TOUR'];

const styles = createStyles({
  root: {
    margin: '0 10px',
  },
});

interface IProps {
  tournaments: Tournament[];
  loadTournaments: (
    filters: { gender: string; league: string; season: string },
  ) => void;
  leagues: string[];
  classes: any;
}

const thisYear = new Date().getFullYear().toString();

function sortDescending(a, b) {
  return new Date(b.start).getTime() - new Date(a.start).getTime();
}

function sameDay(d1: Date, d2: Date): boolean {
  return (
    d1.getFullYear() === d2.getFullYear() &&
    d1.getMonth() === d2.getMonth() &&
    d1.getDay() === d2.getDay()
  );
}

function groupTournaments(tournaments: Tournament[]) {
  const grouped = [];

  let previous = null;

  tournaments.forEach(t => {
    if (!previous || !sameDay(new Date(previous.start), new Date(t.start))) {
      grouped.push([t]);
    } else {
      grouped[grouped.length - 1].push(t);
    }

    previous = t;
  });

  return grouped;
}

function renderHeader(tournaments: Tournament[]) {
  return (
    <DayHeader
      appendix={`(${tournaments.length})`}
      date={new Date(tournaments[0].start)}
    />
  );
}

class Volleynet extends React.Component<IProps> {
  static mapDispatchToProps = {
    loadTournaments: loadTournamentsAction,
  };
  static buildActions({ leagues = [] }: IProps) {
    return leagues.map(league =>
      loadTournamentsAction({
        gender: 'M',
        league,
        season: thisYear,
      }),
    );
  }

  static getParameters(query) {
    let { leagues = ['AMATEUR TOUR'] } = query;

    if (!Array.isArray(leagues)) {
      leagues = [leagues];
    }

    // TODO: remove uniques
    leagues = leagues.filter(l => defaultLeagues.includes(l));

    return { leagues };
  }

  static mapStateToProps(state, { leagues }: IProps) {
    const tournaments = tournamentsByLeagueSelector(state, leagues);

    return { tournaments };
  }

  renderList = (tournaments: Tournament[]) => {
    const { classes } = this.props;

    return (
      <div className={classes.root}>
        <TournamentList
          tournaments={tournaments}
          onTournamentClick={this.onTournamentClick}
        />
      </div>
    );
  };

  componentDidUpdate(prevProps) {
    const { loadTournaments, leagues } = this.props;

    if (!ArrayUtils.equals(leagues, prevProps.leagues) && leagues) {
      leagues.forEach(league => {
        loadTournaments({ gender: 'M', league, season: thisYear });
      });
    }
  }

  onTournamentClick = (t: Tournament) => {
    Router.push({ pathname: '/volleynet/tournament', query: { id: t.id } });
  };

  onLeagueChange = (_, selectedLeagues) => {
    Router.push({
      pathname: '/volleynet',
      query: { leagues: selectedLeagues },
    });
  };

  render() {
    const { leagues, tournaments } = this.props;

    let content = <CenteredLoading />;

    if (tournaments) {
      content = (
        <GroupedList<Tournament>
          groupItems={groupTournaments}
          items={tournaments.sort(sortDescending)}
          renderHeader={renderHeader}
          renderList={this.renderList}
        />
      );
    }

    return (
      <Layout title={{ text: 'Volleynet', href: '' }}>
        <LeagueSelect selected={leagues} onChange={this.onLeagueChange} />
        {content}
      </Layout>
    );
  }
}

export default withStyles(styles)(withAuth(Volleynet));
