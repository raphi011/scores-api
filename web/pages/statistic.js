// @flow
import React from 'react';
import { withStyles } from 'material-ui/styles';
import DateRangeIcon from 'material-ui-icons/DateRange';
import IconButton from 'material-ui/IconButton';
import Typography from 'material-ui/Typography';
import Paper from 'material-ui/Paper';
import Toolbar from 'material-ui/Toolbar';
import Menu, { MenuItem } from 'material-ui/Menu';
import Router from 'next/router';

import withAuth from '../containers/AuthContainer';
import Layout from '../containers/LayoutContainer';
import StatisticList from '../components/StatisticList';
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
  static getParameters(query) {
    let { filter = 'month' } = query;

    filter = filter.toLowerCase();

    return { filter };
  }

  static buildActions(parameters) {
    const { filter } = parameters;

    const actions = [loadStatisticsAction(filter), userOrLoginRouteAction()];

    return actions;
  }

  static shouldComponentUpdate(lastProps, nextProps) {
    return lastProps.filter !== nextProps.filter;
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

  onSetTodayFilter = () => {
    this.onCloseFilterMenu();
    Router.push('/statistic?filter=today');
  };

  onSetMonthFilter = () => {
    this.onCloseFilterMenu();
    Router.push('/statistic?filter=month');
  };

  onSetThisYearFilter = () => {
    this.onCloseFilterMenu();
    Router.push('/statistic?filter=thisyear');
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
            <MenuItem onClick={this.onSetTodayFilter}>Today</MenuItem>
            <MenuItem onClick={this.onSetMonthFilter}>Last month</MenuItem>
            <MenuItem onClick={this.onSetThisYearFilter}>This year</MenuItem>
            <MenuItem onClick={this.onSetAllFilter}>All</MenuItem>
          </Menu>
        </Toolbar>
        <Paper>
          <StatisticList
            onPlayerClick={this.onRowClick}
            statistics={statistics}
          />
        </Paper>
      </Layout>
    );
  }
}

export default withAuth(withStyles(styles)(Statistics));
