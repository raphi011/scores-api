import React from 'react';

import Router from 'next/router';

import { createStyles, WithStyles, withStyles } from '@material-ui/core/styles';

import { QueryStringMapObject } from 'next';
import CenteredLoading from '../../components/CenteredLoading';
import DayHeader from '../../components/DayHeader';
import GroupedList from '../../components/GroupedList';
import TournamentList from '../../components/volleynet/TournamentList';
import withAuth from '../../containers/AuthContainer';
import Layout from '../../containers/LayoutContainer';
import { userSelector } from '../../redux/auth/selectors';
import {
  loadTournamentAction,
  loadTournamentsAction,
} from '../../redux/entities/actions';
import {
  tournamentsByLeagueSelector,
  tournamentSelector,
} from '../../redux/entities/selectors';
import { Store } from '../../redux/store';
import { Tournament, User } from '../../types';
import * as ArrayUtils from '../../utils/array';
import TournamentFilters from '../../components/volleynet/filters/TournamentFilters';

const defaultLeagues = ['AMATEUR TOUR', 'PRO TOUR', 'JUNIOR TOUR'];

const styles = createStyles({
  left: {
    width: '300px',
  },
  right: {
    flexGrow: 1,
  },
  root: {
    display: 'flex',
    flexDirection: 'row',
  },
});

interface Props extends WithStyles<typeof styles> {
  tournament?: Tournament;
  tournaments: Tournament[];
  tournamentId: string;
  leagues: string[];
  user: User;

  loadTournament: (tournamentId: string) => void;
  loadTournaments: (
    filters: { gender: string; league: string; season: string },
  ) => void;
}

class Volleynet extends React.Component<Props> {
  static mapDispatchToProps = {
    loadTournament: loadTournamentAction,
    loadTournaments: loadTournamentsAction,
  };

  static buildActions({ leagues = [], tournamentId }: Props) {
    const actions = leagues.map(league =>
      loadTournamentsAction({
        gender: 'M',
        league,
        season: thisYear,
      }),
    );

    if (tournamentId) {
      actions.push(loadTournamentAction(tournamentId));
    }

    return actions;
  }

  static getParameters(query: QueryStringMapObject) {
    const { tournamentId } = query;
    let { leagues = ['AMATEUR TOUR'] } = query;

    if (!Array.isArray(leagues)) {
      leagues = [leagues];
    }

    leagues = leagues.filter((l: string) => defaultLeagues.includes(l));

    return { leagues, tournamentId };
  }

  static mapStateToProps(state: Store, { leagues, tournamentId }: Props) {
    const tournaments = tournamentsByLeagueSelector(state, leagues);
    const tournament = tournamentSelector(state, Number(tournamentId));
    const { user } = userSelector(state);

    return { tournament, tournaments, user };
  }

  renderList = (tournaments: Tournament[]) => {
    return (
      <div key={tournaments[0].id}>
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
    Router.push({
      pathname: '/volleynet/tournament',
      query: { id: t.id },
    });
  };

  onLeagueChange = (_, selectedLeagues) => {
    Router.push({
      pathname: '/volleynet',
      query: { leagues: selectedLeagues },
    });
  };

  render() {
    const { leagues, user, tournaments, tournament, classes } = this.props;

    let leftContent = <CenteredLoading />;

    if (tournaments) {
      leftContent = (
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
        <div className={classes.root}>
          <div className={classes.left}>
            <TournamentFilters />
          </div>
          <div className={classes.right}>{leftContent}</div>
        </div>
      </Layout>
    );
  }
}

const thisYear = '2018'; // new Date().getFullYear().toString();

function sortDescending(a: Tournament, b: Tournament) {
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
      key={tournaments[0].start}
      appendix={`(${tournaments.length})`}
      date={new Date(tournaments[0].start)}
    />
  );
}

export default withStyles(styles)(withAuth(Volleynet));
