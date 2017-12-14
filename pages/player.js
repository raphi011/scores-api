// @flow

import React from 'react';
import withRedux from 'next-redux-wrapper';

import withRoot from '../components/withRoot';
import Layout from '../components/Layout';
import initStore, { dispatchActions } from '../redux/store';
import {
  loadPlayerAction,
  userOrLoginRouteAction,
  loadPlayersAction,
  loadStatisticAction,
} from '../redux/actions/action';
import PlayerView from '../components/PlayerView';
import { playerSelector, statisticSelector } from '../redux/reducers/reducer';
import type { Player, Statistic } from '../types';

type Props = {
  player: Player,
  statistic: Statistic,
};

class PlayerInfo extends React.Component<Props> {
  static async getInitialProps({ store, query, req, res, isServer }) {
    const actions = [userOrLoginRouteAction()];

    const { id } = query;

    if (id) {
      actions.push(loadPlayerAction(id), loadStatisticAction(id));
    } else {
      actions.push(loadPlayersAction());
    }

    await dispatchActions(store.dispatch, isServer, req, res, actions);

    return { playerId: id };
  }

  render() {
    const { player, statistic } = this.props;

    return (
      <Layout title="Players">
        <PlayerView player={player} statistic={statistic} />
      </Layout>
    );
  }
}

function mapStateToProps(state, ownProps) {
  const { playerId } = ownProps;
  const player = playerSelector(state, playerId);
  const statistic = statisticSelector(state, playerId);

  return {
    player,
    statistic,
  };
}

const mapDispatchToProps = {};

export default withRedux(initStore, mapStateToProps, mapDispatchToProps)(
  withRoot(PlayerInfo),
);
