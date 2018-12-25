import React from 'react';

import Router from 'next/router';

import { createStyles, WithStyles, withStyles } from '@material-ui/core/styles';

import { Fade } from '@material-ui/core';
import CenteredLoading from '../../components/CenteredLoading';
import DayHeader from '../../components/DayHeader';
import GroupedList from '../../components/GroupedList';
import LeagueSelect from '../../components/volleynet/LeagueSelect';
import TournamentList from '../../components/volleynet/TournamentList';
import TournamentView from '../../components/volleynet/TournamentView';
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

const defaultLeagues = ['AMATEUR TOUR', 'PRO TOUR', 'JUNIOR TOUR'];

const styles = createStyles({
  left: {
    flexGrow: 1,
    maxWidth: '500px',
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

  static getParameters(query) {
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
    const { leagues } = this.props;

    Router.push({
      pathname: '/volleynet',
      query: { tournamentId: t.id, leagues },
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

    let rightContent: any = <span>Select a tournament</span>;

    if (tournament) {
      rightContent = <TournamentView tournament={tournament} user={user} />;
    }

    return (
      <Layout title={{ text: 'Volleynet', href: '' }}>
        <div className={classes.root}>
          <div className={classes.left}>
            <LeagueSelect selected={leagues} onChange={this.onLeagueChange} />
            {leftContent}
          </div>
          <div className={classes.right}>
            <Fade in={!!tournament}>{rightContent}</Fade>
          </div>
        </div>
      </Layout>
    );
  }
}

const thisYear = new Date().getFullYear().toString();

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
