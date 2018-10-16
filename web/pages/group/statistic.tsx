import IconButton from '@material-ui/core/IconButton';
import Menu from '@material-ui/core/Menu';
import MenuItem from '@material-ui/core/MenuItem';
import Paper from '@material-ui/core/Paper';
import { createStyles, withStyles } from '@material-ui/core/styles';
import Toolbar from '@material-ui/core/Toolbar';
import Typography from '@material-ui/core/Typography';
import DateRangeIcon from '@material-ui/icons/DateRange';
import Router from 'next/router';
import React from 'react';

import StatisticList from '../../components/StatisticList';
import withAuth from '../../containers/AuthContainer';
import Layout from '../../containers/LayoutContainer';
import { userOrLoginRouteAction } from '../../redux/actions/auth';
import { loadGroupStatisticsAction } from '../../redux/actions/entities';
import { statisticByGroupSelector } from '../../redux/reducers/entities';
import { PlayerStatistic, StatisticFilter } from '../../types';

const styles = createStyles({
  title: {
    flex: '0 0 auto',
  },
  toolbar: {
    justifyContent: 'space-between',
  },
});

interface IProps {
  groupId: number;
  filter: StatisticFilter;
  statistics: PlayerStatistic[];
  classes: any;
}

interface IState {
  filterMenuOpen: boolean;
  anchorEl?: HTMLElement;
}

class Statistics extends React.Component<IProps, IState> {
  static getParameters(query) {
    let { filter = 'month', groupId } = query;

    groupId = Number.parseInt(groupId, 10) || 0;
    filter = filter.toLowerCase();

    return { filter, groupId };
  }

  static buildActions(parameters) {
    const { filter, groupId } = parameters;

    const actions = [
      loadGroupStatisticsAction(groupId, filter),
      userOrLoginRouteAction(),
    ];

    return actions;
  }

  static shouldComponentUpdate(lastProps, nextProps) {
    return lastProps.filter !== nextProps.filter;
  }

  static mapStateToProps(state, { groupId }) {
    const statistics = statisticByGroupSelector(state, groupId);

    return {
      statistics,
    };
  }

  state = {
    anchorEl: null,
    filterMenuOpen: false,
  };

  onOpenFilterMenu = event => {
    this.setState({ filterMenuOpen: true, anchorEl: event.currentTarget });
  };

  onCloseFilterMenu = () => {
    this.setState({ filterMenuOpen: false, anchorEl: null });
  };

  onSetFilter = (filter: string) => {
    const { groupId } = this.props;

    this.onCloseFilterMenu();
    Router.push(`/group/statistic?groupId=${groupId}&filter=${filter}`);
  };

  onSetTodayFilter = () => this.onSetFilter('today');

  onSetMonthFilter = () => this.onSetFilter('month');

  onSetThisYearFilter = () => this.onSetFilter('thisyear');

  onSetAllFilter = () => this.onSetFilter('all');

  onRowClick = playerId => {
    Router.push(`/player?id=${playerId}`);
  };

  timeFilter = () => {
    const { filter } = this.props;

    if (filter === 'all') {
      return 'Ranks';
    }

    return `Ranks by ${filter}`;
  };

  render() {
    const { statistics, classes } = this.props;

    return (
      <Layout title={{ text: 'Players', href: '' }}>
        <Toolbar className={classes.toolbar}>
          <Typography variant="h6">{this.timeFilter()}</Typography>
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
