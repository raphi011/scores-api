import React from 'react';

import Router from 'next/router';

import { createStyles, WithStyles, withStyles } from '@material-ui/core/styles';

import { QueryStringMapObject } from 'next';
import CenteredLoading from '../../components/CenteredLoading';
import DayHeader from '../../components/DayHeader';
import GroupedList from '../../components/GroupedList';
import TournamentFilters from '../../components/volleynet/filters/TournamentFilters';
import TournamentList from '../../components/volleynet/TournamentList';
import withAuth from '../../containers/AuthContainer';
import Layout from '../../containers/LayoutContainer';
import { userSelector } from '../../redux/auth/selectors';
import {
  loadTournamentAction,
  loadTournamentsAction,
} from '../../redux/entities/actions';
import { filteredTournamentsSelector } from '../../redux/entities/selectors';
import { Store } from '../../redux/store';
import { Tournament, User } from '../../types';

const defaultLeagues = ['amateur-tour', 'pro-tour', 'junior-tour'];

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
  tournaments: Tournament[];
  leagues: string[];
  user: User;

  loadTournaments: (
    filters: { gender: string; league: string; season: string },
  ) => void;
}

class Volleynet extends React.Component<Props> {
  static mapDispatchToProps = {
    loadTournament: loadTournamentAction,
    loadTournaments: loadTournamentsAction,
  };

  static buildActions({ genders: gender, season, leagues: league = [] }: Props) {
      return [loadTournamentsAction({
        gender,
        league,
        season,
      })];
  }

  static getParameters(query: QueryStringMapObject) {
    let { season = "2018", genders = ['M', 'W'], leagues = ['amateur-tour'] } = query;

    if (!Array.isArray(leagues)) {
      leagues = [leagues];
    }
    if (!Array.isArray(genders)) {
      genders = [genders];
    }

    leagues = leagues.filter((l: string) => defaultLeagues.includes(l));

    return { leagues, genders, season };
  }

  static mapStateToProps(state: Store) {
    const tournaments = filteredTournamentsSelector(state);
    const { user } = userSelector(state);

    return { tournaments, user };
  }

  constructor(props) {
    super(props);

    const { leagues, genders, season } = this.props;

    this.state = {
      leagues,
      genders,
      season: Number(season),
    };
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

  onTournamentClick = (t: Tournament) => {
    Router.push({
      pathname: '/volleynet/tournament',
      query: { id: t.id },
    });
  };

  onFilter = (filters) => {
    this.setState(filters);
  }

  onSubmit = () => {
    const { loadTournaments } = this.props;
    const { leagues: league, season, genders: gender } = this.state;

    const query = {
      league,
      season,
      gender,
    };

    Router.push({
      pathname: '/volleynet',
      query,
    });
    
    loadTournaments(query);
  }

  onLeaguesChange = (leagues: string[]) => {
    Router.push({
      pathname: '/volleynet',
      query: { leagues },
    });
  }

  render() {
    const { tournaments, classes } = this.props;
    const { leagues, genders, season } = this.state;

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
            <TournamentFilters
              leagues={leagues}
              genders={genders}
              season={season}
              onChange={this.onFilter}
              onSubmit={this.onSubmit}
            />
          </div>
          <div className={classes.right}>{leftContent}</div>
        </div>
      </Layout>
    );
  }
}

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
