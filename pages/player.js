// @flow

import React from 'react';
import withRedux from 'next-redux-wrapper';
import Tabs, { Tab } from 'material-ui/Tabs';
import Typography from 'material-ui/Typography';
import MatchList from '../containers/MatchListContainer';

import withRoot from '../components/withRoot';
import Layout from '../components/Layout';
import initStore, { dispatchActions } from '../redux/store';
import {
  loadPlayerAction,
  userOrLoginRouteAction,
  loadPlayersAction,
  loadStatisticAction,
  loadPlayerMatchesAction,
} from '../redux/actions/action';
import PlayerView from '../components/PlayerView';
import { playerSelector, statisticSelector, playerMatchesSelector } from '../redux/reducers/reducer';
import type { Player, Statistic } from '../types';

type Props = {
  player: Player,
  statistic: Statistic,
  matches: Array<MatchList>,
  playerId: number,
};

type State = {
  tabOpen: number,
};

class PlayerInfo extends React.Component<Props, State> {
  static async getInitialProps({ store, query, req, res, isServer }) {
    const actions = [userOrLoginRouteAction()];

    const { id } = query;

    if (id) {
      actions.push(loadPlayerAction(id), loadPlayerMatchesAction(id), loadStatisticAction(id));
    } else {
      actions.push(loadPlayersAction());
    }

    await dispatchActions(store.dispatch, isServer, req, res, actions);

    return { playerId: id };
  }

  state = {
    tabOpen: 0,
  };

  onTabClick = (event, index) => {
    this.setState({ tabOpen: index });
  }

  render() {
    const { player, matches, statistic, playerId } = this.props;

    if (!playerId) {
      return (
        <Layout title="Players">
          <Typography align="center" type="display4">
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
          fullWidth>
          <Tab label={`Matches (${matches.length})`} />
          <Tab label="Teams" />
        </Tabs>
        {this.state.tabOpen === 0 ? (
          <MatchList matches={matches} />
        ) : (
        <Typography align="center">
          List of teams
        </Typography>
        )}
      </Layout>
    );
  }
}

function mapStateToProps(state, ownProps) {
  const { playerId } = ownProps;
  const player = playerSelector(state, playerId);
  const statistic = statisticSelector(state, playerId);
  const matches = playerMatchesSelector(state, playerId);

  return {
    player,
    statistic,
    matches,
  };
}

const mapDispatchToProps = {};

export default withRedux(initStore, mapStateToProps, mapDispatchToProps)(
  withRoot(PlayerInfo),
);
