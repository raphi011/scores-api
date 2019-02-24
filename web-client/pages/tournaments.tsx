import React from 'react';

import { QueryStringMapObject } from 'next';
import Router from 'next/router';

import ExpansionPanel from '@material-ui/core/ExpansionPanel';
import ExpansionPanelDetails from '@material-ui/core/ExpansionPanelDetails';
import ExpansionPanelSummary from '@material-ui/core/ExpansionPanelSummary';
import {
  createStyles,
  Theme,
  WithStyles,
  withStyles,
} from '@material-ui/core/styles';
import { Breakpoint } from '@material-ui/core/styles/createBreakpoints';
import Typography from '@material-ui/core/Typography';
import withWidth from '@material-ui/core/withWidth';
import ExpandMoreIcon from '@material-ui/icons/ExpandMore';

import DayHeader from '../components/DayHeader';
import GroupedList from '../components/GroupedList';
import TournamentFilters, {
  Filters,
} from '../components/volleynet/filters/TournamentFilters';
import TournamentList from '../components/volleynet/TournamentList';
import withAuth from '../containers/AuthContainer';
import Layout from '../containers/LayoutContainer';
import { userSelector } from '../redux/auth/selectors';
import {
  loadTournamentAction,
  loadTournamentsAction,
} from '../redux/entities/actions';
import { filteredTournamentsSelector } from '../redux/entities/selectors';
import { Store } from '../redux/store';
import { Gender, Tournament, User } from '../types';
import { sameDay } from '../utils/date';

const defaultLeagues = ['amateur-tour', 'pro-tour', 'junior-tour'];

const styles = (theme: Theme) =>
  createStyles({
    found: {
      color: theme.palette.grey[400],
    },
    primary: {
      flexGrow: 1,
    },
    root: {
      display: 'flex',

      flexDirection: 'row',
      marginTop: '40px',

      [theme.breakpoints.down('xs')]: {
        flexDirection: 'column',
      },
    },
    secondary: {
      [theme.breakpoints.up('sm')]: {
        paddingRight: '50px',
        width: '250px',
      },
    },
  });

interface Props extends WithStyles<typeof styles> {
  tournaments: Tournament[];
  league: string[];
  season: number;
  gender: Gender[];
  user: User;
  width: Breakpoint;

  loadTournaments: (filters: {
    gender: Gender[];
    league: string[];
    season: number;
  }) => void;
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
    return [
      loadTournamentsAction({
        gender,
        league,
        season,
      }),
    ];
  }

  static getParameters(query: QueryStringMapObject) {
    const { season = '2019' } = query;
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
      pathname: '/tournament',
      query: { id: t.id },
    });
  };

  onFilter = async (filters: Filters) => {
    this.setState({ loading: true });

    const { loadTournaments } = this.props;

    const query = filters;

    Router.push({
      pathname: '/',
      query,
    });

    await loadTournaments(query);

    this.setState({ loading: false });
  };

  renderFilters = () => {
    const { league, gender, season, width } = this.props;
    const { loading } = this.state;

    const filters = (
      <TournamentFilters
        loading={loading}
        league={league}
        gender={gender}
        season={season}
        onFilter={this.onFilter}
      />
    );

    if (width === 'xs') {
      return (
        <ExpansionPanel>
          <ExpansionPanelSummary expandIcon={<ExpandMoreIcon />}>
            <Typography style={{ fontSize: '20px' }}>Filters</Typography>
          </ExpansionPanelSummary>
          <ExpansionPanelDetails>{filters}</ExpansionPanelDetails>
        </ExpansionPanel>
      );
    }

    return filters;
  };

  render() {
    const { tournaments, classes } = this.props;

    return (
      <Layout title={{ text: 'Tournaments', href: '' }}>
        <Typography variant="h1">
          Tournaments{' '}
          <span className={classes.found}>({tournaments.length})</span>
        </Typography>
        <div className={classes.root}>
          <div className={classes.secondary}>{this.renderFilters()}</div>
          <div className={classes.primary}>
            <GroupedList<Tournament>
              groupItems={groupTournaments}
              items={tournaments}
              renderHeader={renderHeader}
              renderList={this.renderList}
            />
          </div>
        </div>
      </Layout>
    );
  }
}

function groupTournaments(tournaments: Tournament[]) {
  const grouped: Tournament[][] = [];

  let previous: Tournament | null = null;

  tournaments.forEach((t: Tournament) => {
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

export default withStyles(styles)(withAuth(withWidth()(Volleynet)));
