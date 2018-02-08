// @flow

import React from 'react';
import Tabs, { Tab } from 'material-ui/Tabs';
import Typography from 'material-ui/Typography';

import MatchList from '../containers/MatchListContainer';
import withAuth from '../containers/AuthContainer';

import Layout from '../components/Layout';
import { dispatchAction } from '../redux/store';
import {
  loadPlayerAction,
  loadPlayersAction,
  loadPlayerStatisticAction,
  loadPlayerMatchesAction,
} from '../redux/actions/entities';
import { multiApiAction } from '../redux/actions/api';
import PlayerView from '../components/PlayerView';
import {
  playerSelector,
  statisticByPlayerSelector,
  matchesByPlayerSelector,
} from '../redux/reducers/entities';
import type { Player, Statistic } from '../types';

type Props = {
  player: Player,
  statistic: Statistic,
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
  static async getInitialProps({ store, query, req, res, isServer }) {
    let action;

    const { id } = query;

    if (id) {
      action = multiApiAction([
        loadPlayerAction(id),
        loadPlayerMatchesAction(id),
        loadPlayerStatisticAction(id),
      ]);
    } else {
      action = loadPlayersAction();
    }

    await dispatchAction(store.dispatch, isServer, req, res, action);

    return { playerId: id };
  }

  static mapStateToProps(state, ownProps) {
    const { playerId } = ownProps;
    const player = playerSelector(state, playerId);
    const statistic = statisticByPlayerSelector(state, playerId);
    const matches = matchesByPlayerSelector(state, playerId);

    return {
      player,
      statistic,
      matches,
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
    const { player, matches, statistic, playerId } = this.props;
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

    return (
      <Layout title="Players">
        <PlayerView player={player} statistic={statistic} />
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
            onLoadMore={this.onLoadMore}
            loading={loading}
            hasMore={hasMore}
          />
        ) : (
          <Typography align="center">List of teams</Typography>
        )}
      </Layout>
    );
  }
}

export default withAuth(PlayerInfo);
