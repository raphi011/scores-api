import React from 'react';
import fetch from 'isomorphic-unfetch';
import withRedux from 'next-redux-wrapper';
import { withStyles } from 'material-ui/styles';

import withRoot from '../components/withRoot';
import Layout from '../components/Layout';
import initStore, { dispatchActions } from '../redux/store';
import {
  loadPlayerAction,
  userOrLoginRouteAction,
  loadPlayersAction,
} from '../redux/actions/action';
import PlayerView from '../components/PlayerView';
import { playerSelector } from '../redux/reducers/reducer';

const styles = theme => ({});

class Player extends React.Component {
  static async getInitialProps({ store, query, req, res, isServer }) {
    const actions = [userOrLoginRouteAction()];

    const { id } = query;

    if (id) {
      actions.push(loadPlayerAction(id));
    } else {
      actions.push(loadPlayersAction());
    }

    await dispatchActions(store.dispatch, isServer, req, res, actions);

    return { playerId: id };
  }

  render() {
    const { player, playerId, classes } = this.props;

    return (
      <Layout title="Players">
        <PlayerView player={player} />
      </Layout>
    );
  }
}

function mapStateToProps(state, ownProps) {
  const { playerId } = ownProps;
  console.log(playerId)
  const player = playerSelector(state, playerId);

  return {
    player,
  };
}

const mapDispatchToProps = {};

export default withRedux(initStore, mapStateToProps, mapDispatchToProps)(
  withRoot(withStyles(styles)(Player)),
);
