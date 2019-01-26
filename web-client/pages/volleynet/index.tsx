import React from 'react';

import Router from 'next/router';

import { createStyles, WithStyles, withStyles } from '@material-ui/core/styles';

import { QueryStringMapObject } from 'next';
import CenteredLoading from '../../components/CenteredLoading';
import DayHeader from '../../components/DayHeader';
import GroupedList from '../../components/GroupedList';
import TournamentFilters, { Filters } from '../../components/volleynet/filters/TournamentFilters';
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
import { Gender, Tournament, User } from '../../types';

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

  league: string[];
  season: number;
  gender: Gender[];

  user: User;

  loadTournaments: (
    filters: { gender: Gender[]; league: string[]; season: number },
  ) => void;
}

interface State {
  loading: boolean;
}

class Volleynet extends React.Component<Props, State> {
  static mapDispatchToProps = {
    loadTournament: loadTournamentAction,
    loadTournaments: loadTournamentsAction,
  };

  static buildActions({ gender, season, league = [] }: Props) {
      return [loadTournamentsAction({
        gender,
        league,
        season,
      })];
  }

  static getParameters(query: QueryStringMapObject) {
    const { season = "2018" } = query;
    let { gender = ['M', 'W'], league = ['amateur-tour'] } = query;

    if (!Array.isArray(league)) {
      league = [league];
    }
    if (!Array.isArray(gender)) {
      gender = [gender];
    }

    league = league.filter((l: string) => defaultLeagues.includes(l));

    return { league, gender, season: Number(season) };
  }

  static mapStateToProps(state: Store) {
    const tournaments = filteredTournamentsSelector(state);
    const { user } = userSelector(state);

    return { tournaments, user };
  }

  state = {
    loading: false,
  };


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

  onFilter = async (filters: Filters) => {
    this.setState({ loading: true })

    const { loadTournaments } = this.props;

    const query = filters; 

    Router.push({
      pathname: '/volleynet',
      query,
    });

    
    await loadTournaments(query);

    this.setState({ loading: false })
  }

  render() {
    const { league, gender, season, tournaments, classes } = this.props;
    const { loading } = this.state;

    let leftContent = <CenteredLoading />;

    if (tournaments) {
      leftContent = (
        <GroupedList<Tournament>
          groupItems={groupTournaments}
          items={tournaments}
          // items={tournaments.sort(sortDescending)}
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
              loading={loading}
              league={league}
              gender={gender}
              season={season}
              onFilter={this.onFilter}
            />
          </div>
          <div className={classes.right}>{leftContent}</div>
        </div>
      </Layout>
    );
  }
}

// function sortDescending(a: Tournament, b: Tournament) {
//   return new Date(b.start).getTime() - new Date(a.start).getTime();
// }

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
