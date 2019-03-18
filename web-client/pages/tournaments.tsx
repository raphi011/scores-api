import React from 'react';

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

import * as Query from '../utils/query';
import DayHeader from '../components/DayHeader';
import GroupedList from '../components/GroupedList';
import TournamentFilters, {
  Filters,
} from '../components/volleynet/filters/TournamentFilters';
import TournamentList from '../components/volleynet/TournamentList';
import withAuth from '../hoc/next/withAuth';
import Layout from '../containers/LayoutContainer';
import { userSelector } from '../redux/auth/selectors';
import { loadTournamentsAction } from '../redux/entities/actions';
import { filteredTournamentsSelector } from '../redux/entities/selectors';
import { Store } from '../redux/store';
import { Tournament, User } from '../types';
import { sameDay } from '../utils/date';
import {
  tournamentFilterSelector,
  filterOptionsSelector,
} from '../redux/tournaments/selectors';
import { FilterOptions, Filter } from '../redux/tournaments/reducer';
import {
  setFilterAction,
  loadFilterOptionsAction,
} from '../redux/tournaments/actions';
import { QueryStringMapObject } from 'next';
import { dispatchAction } from '../redux/actions';
import withConnect, { ClientContext } from '../hoc/next/withConnect';

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
  leagues: string[];
  genders: string[];
  season: string;
  options: FilterOptions;
  filters: Filter;
  user: User;
  width: Breakpoint;

  setFilter: (filters: Filter) => void;
  loadTournaments: (filters: Filter) => void;
}

interface State {
  loading: boolean;
  filterDialogOpen: boolean;
}

function filtersFromQuery(query: QueryStringMapObject): Filters {
  const seasons = Query.one(query, 'season');
  const genders = Query.multiple(query, 'gender');
  const leagues = Query.multiple(query, 'league');

  return { leagues, genders, seasons };
}

class Volleynet extends React.Component<Props, State> {
  static async getInitialProps({ req, res, store, query }: ClientContext) {
    let options = filterOptionsSelector(store.getState());

    if (!options) {
      await dispatchAction(store.dispatch, loadFilterOptionsAction(), req, res);
      options = filterOptionsSelector(store.getState());
    }

    const queryFilters = filtersFromQuery(query);
    await store.dispatch(setFilterAction(queryFilters));

    const state = store.getState();

    const filters = tournamentFilterSelector(state);

    return { filters, options };
  }

  static buildActions({ filters }: Props) {
    return [loadTournamentsAction(filters)];
  }

  static mapDispatchToProps = {
    loadTournaments: loadTournamentsAction,
    setFilter: setFilterAction,
  };

  static mapStateToProps(state: Store) {
    const tournaments = filteredTournamentsSelector(state);
    const user = userSelector(state);

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

    const { setFilter, loadTournaments } = this.props;

    const query = filters;

    Router.push({
      pathname: '/',
      query,
    });

    await loadTournaments(query);
    setFilter(query);

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
    const { filters, options, tournaments, width, classes } = this.props;
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
        filters={filters}
        options={options}
        loading={loading}
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

    const tournamentCount = tournaments ? tournaments.length : 0;

    return (
      <Layout title={{ text: 'Tournaments', href: '' }}>
        <Typography variant="h1">
          Tournaments <span className={classes.found}>({tournamentCount})</span>
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

export default withAuth(
  withConnect(withWidth()(withStyles(styles)(Volleynet))),
);
