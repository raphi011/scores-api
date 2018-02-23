// @flow

import React from 'react';
import Tabs, { Tab } from 'material-ui/Tabs';
import Typography from 'material-ui/Typography';
import { CircularProgress } from 'material-ui/Progress';

import MatchList from '../containers/MatchListContainer';
import withAuth from '../containers/AuthContainer';
import StatisticList from '../components/StatisticList';

import Layout from '../components/Layout';
import {
  loadPlayerAction,
  loadPlayersAction,
  loadPlayerTeamStatisticAction,
  loadPlayerStatisticAction,
  loadPlayerMatchesAction,
} from '../redux/actions/entities';
import { multiApiAction } from '../redux/actions/api';
import PlayerView from '../components/PlayerView';
import {
  playerSelector,
  statisticByPlayerSelector,
  statisticByPlayerTeamSelector,
  matchesByPlayerSelector,
} from '../redux/reducers/entities';
import type { Player, Statistic } from '../types';

type Props = {
  player: Player,
  statistic: Statistic,
  teamStatistic: Array<Statistic>,
  matches: Array<MatchList>,
  playerId: number,
  loadMatches: (number, ?string) => Promise<any>,
};

type State = {
  tabOpen: number,
  loading: boolean,
  hasMore: boolean,
};

class PlayerInfo extends React.Component<Props, State> {
  static getParameters(query) {
    const { id } = query;

    const playerId = Number.parseInt(id, 10);

    if (Number.isInteger(playerId)) {
      return { playerId };
    }

    return {};
  }

  static shouldComponentUpdate(lastProps, nextProps) {
    return lastProps.playerId !== nextProps.playerId;
  }

  static buildActions({ playerId }) {
    let action;

    if (playerId) {
      action = multiApiAction([
        loadPlayerAction(playerId),
        loadPlayerMatchesAction(playerId),
        loadPlayerStatisticAction(playerId),
        loadPlayerTeamStatisticAction(playerId),
      ]);
    } else {
      action = loadPlayersAction();
    }

    return [action];
  }

  static mapStateToProps(state, ownProps) {
    const { playerId } = ownProps;
    const player = playerSelector(state, playerId);
    const statistic = statisticByPlayerSelector(state, playerId);
    const matches = matchesByPlayerSelector(state, playerId);
    const teamStatistic = statisticByPlayerTeamSelector(state, playerId);

    return {
      player,
      statistic,
      matches,
      teamStatistic,
    };
  }

  static mapDispatchToProps = {
    loadMatches: loadPlayerMatchesAction,
  };

  state = {
    tabOpen: 0,
    loading: false,
    hasMore: true,
  };

  onLoadMore = async () => {
    const { playerId, loadMatches, matches } = this.props;

    this.setState({ loading: true });

    const lastElement = matches[matches.length - 1];

    const after = lastElement ? lastElement.createdAt : '';

    const newState = {
      loading: false,
      hasMore: true,
    };

    try {
      const result = await loadMatches(playerId, after);
      newState.hasMore = !result.empty;
    } catch (e) {
      newState.hasMore = false;
    } finally {
      this.setState(newState);
    }
  };

  onTabClick = (event, index) => {
    this.setState({ tabOpen: index });
  };

  render() {
    const { player, matches, statistic, teamStatistic, playerId } = this.props;
    const { loading, hasMore } = this.state;

    if (!playerId) {
      return (
        <Layout title="Players">
          <Typography align="center" variant="display4">
            Players: todo!
          </Typography>
        </Layout>
      );
    }

    const loadingPlayer = !(player && statistic);

    return (
      <Layout title="Players">
        {loadingPlayer ? (
          <CircularProgress />
        ) : (
          <PlayerView player={player} statistic={statistic} />
        )}
        <Tabs
          onChange={this.onTabClick}
          value={this.state.tabOpen}
          textColor="primary"
          fullWidth
        >
          <Tab label={`Matches (${matches.length})`} />
          <Tab label="Teams" />
        </Tabs>
        {this.state.tabOpen === 0 ? (
          <MatchList
            matches={matches}
            highlightPlayerId={playerId}
            onLoadMore={this.onLoadMore}
            loading={loading}
            hasMore={hasMore}
          />
        ) : (
          <StatisticList statistics={teamStatistic} />
        )}
      </Layout>
    );
  }
}

export default withAuth(PlayerInfo);
