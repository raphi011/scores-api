import React from 'react';

import { QueryStringMapObject } from 'next';
import Router from 'next/router';

import {
  createStyles,
  Theme,
  WithStyles,
  withStyles,
} from '@material-ui/core/styles';
import { Breakpoint } from '@material-ui/core/styles/createBreakpoints';
import Typography from '@material-ui/core/Typography';
import withWidth from '@material-ui/core/withWidth';
import FilterIcon from '@material-ui/icons/FilterList';
import IconButton from '@material-ui/core/IconButton';
import CloseIcon from '@material-ui/icons/Close';
import Fab from '@material-ui/core/Fab';
import { Dialog } from '@material-ui/core';

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
    dialogContainer: {
      padding: theme.spacing.unit * 2,
    },
    dialogBody: {
      marginTop: '20px',
    },
    root: {
      display: 'flex',

      flexDirection: 'row',
      marginTop: '40px',

      [theme.breakpoints.down('xs')]: {
        flexDirection: 'column',
      },
    },
    fab: {
      position: 'fixed',
      bottom: theme.spacing.unit * 2,
      right: theme.spacing.unit * 2,
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
  filterDialogOpen: boolean;
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
    filterDialogOpen: false,
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
    this.setState({ loading: true, filterDialogOpen: false });

    const { loadTournaments } = this.props;

    const query = filters;

    Router.push({
      pathname: '/',
      query,
    });

    await loadTournaments(query);

    this.setState({ loading: false });
  };

  onOpenFilterDialog = () => {
    this.setState({
      filterDialogOpen: true,
    });
  };

  onCloseFilterDialog = () => {
    this.setState({
      filterDialogOpen: false,
    });
  };

  render() {
    const { league, gender, season, tournaments, width, classes } = this.props;
    const { loading } = this.state;

    const list = (
      <GroupedList<Tournament>
        groupItems={groupTournaments}
        items={tournaments}
        renderHeader={renderHeader}
        renderList={this.renderList}
      />
    );

    const filter = (
      <TournamentFilters
        loading={loading}
        league={league}
        gender={gender}
        season={season}
        onFilter={this.onFilter}
      />
    );

    let body;

    if (width === 'xs') {
      body = (
        <>
          {list}
          <Fab
            onClick={this.onOpenFilterDialog}
            className={classes.fab}
            color="primary"
          >
            <FilterIcon />
          </Fab>
          <Dialog
            fullScreen
            open={this.state.filterDialogOpen}
            onClose={this.onCloseFilterDialog}
          >
            <div className={classes.dialogContainer}>
              <IconButton
                color="inherit"
                onClick={this.onCloseFilterDialog}
                aria-label="Close"
              >
                <CloseIcon />
              </IconButton>
              <div className={classes.dialogBody}>{filter}</div>
            </div>
          </Dialog>
          {/* allow scrolling of fab below last tournament */}
          <div style={{ height: '60px' }} />
        </>
      );
    } else {
      body = (
        <>
          <div className={classes.secondary}>{filter}</div>
          <div className={classes.primary}>{list}</div>
        </>
      );
    }

    return (
      <Layout title={{ text: 'Tournaments', href: '' }}>
        <Typography variant="h1">
          Tournaments{' '}
          <span className={classes.found}>({tournaments.length})</span>
        </Typography>
        <div className={classes.root}>{body}</div>
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
