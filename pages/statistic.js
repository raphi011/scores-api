// @flow
import React from 'react';
import { withStyles } from 'material-ui/styles';
import DateRangeIcon from 'material-ui-icons/DateRange';
import IconButton from 'material-ui/IconButton';
import Typography from 'material-ui/Typography';
import Toolbar from 'material-ui/Toolbar';
import Menu, { MenuItem } from 'material-ui/Menu';
import Router from 'next/router';

import withAuth from '../containers/AuthContainer';
import Layout from '../components/Layout';
import StatisticList from '../components/StatisticList';
import { dispatchActions } from '../redux/store';
import { userOrLoginRouteAction } from '../redux/actions/auth';
import { loadStatisticsAction } from '../redux/actions/entities';
import { allStatisticSelector } from '../redux/reducers/entities';
import type { Statistic, StatisticFilter } from '../types';

const styles = () => ({
  title: {
    flex: '0 0 auto',
  },
  toolbar: {
    justifyContent: 'space-between',
  },
});

type Props = {
  filter: StatisticFilter,
  statistics: Array<Statistic>,
  classes: Object,
};

type State = {
  filterMenuOpen: boolean,
  anchorEl: ?HTMLElement,
};

class Statistics extends React.Component<Props, State> {
  static async getInitialProps({ store, query, req, res, isServer }) {
    let { filter = 'all' } = query;
    filter = filter.toLowerCase();

    const actions = [loadStatisticsAction(filter), userOrLoginRouteAction()];

    await dispatchActions(store.dispatch, isServer, req, res, actions);

    return { filter };
  }

  static mapStateToProps(state) {
    const statistics = allStatisticSelector(state);

    return {
      statistics,
    };
  }

  state = {
    filterMenuOpen: false,
    anchorEl: null,
  };

  onOpenFilterMenu = event => {
    this.setState({ filterMenuOpen: true, anchorEl: event.currentTarget });
  };

  onCloseFilterMenu = () => {
    this.setState({ filterMenuOpen: false, anchorEl: null });
  };

  onSetWeekFilter = () => {
    this.onCloseFilterMenu();
    Router.push('/statistic?filter=week');
  };

  onSetMonthFilter = () => {
    this.onCloseFilterMenu();
    Router.push('/statistic?filter=month');
  };

  onSet3MonthsFilter = () => {
    this.onCloseFilterMenu();
    Router.push('/statistic?filter=quarter');
  };

  onSetYearFilter = () => {
    this.onCloseFilterMenu();
    Router.push('/statistic?filter=year');
  };

  onSetAllFilter = () => {
    this.onCloseFilterMenu();
    Router.push('/statistic?filter=all');
  };

  onRowClick = playerId => {
    Router.push(`/player?id=${playerId}`);
  };

  timeFilter = () => {
    const { filter } = this.props;

    if (filter === 'all') return 'Ranks';

    return `Ranks by ${filter}`;
  };

  render() {
    const { statistics, classes } = this.props;

    return (
      <Layout title="Players">
        <Toolbar className={classes.toolbar}>
          <Typography variant="title">{this.timeFilter()}</Typography>
          <IconButton onClick={this.onOpenFilterMenu}>
            <DateRangeIcon />
          </IconButton>
          <Menu
            anchorEl={this.state.anchorEl}
            open={this.state.filterMenuOpen}
            onClose={this.onCloseFilterMenu}
          >
            <MenuItem onClick={this.onSetWeekFilter}>Last Week</MenuItem>
            <MenuItem onClick={this.onSetMonthFilter}>Last Month</MenuItem>
            <MenuItem onClick={this.onSet3MonthsFilter}>Last 3 Months</MenuItem>
            <MenuItem onClick={this.onSetYearFilter}>Last Year</MenuItem>
            <MenuItem onClick={this.onSetAllFilter}>All</MenuItem>
          </Menu>
        </Toolbar>
        <StatisticList
          onPlayerClick={this.onRowClick}
          statistics={statistics}
        />
      </Layout>
    );
  }
}

export default withAuth(withStyles(styles)(Statistics));
