// @flow

import React from 'react';
import { withStyles } from 'material-ui/styles';
import withRedux from 'next-redux-wrapper';
import Button from 'material-ui/Button';
import AddIcon from 'material-ui-icons/Add';
import Tooltip from 'material-ui/Tooltip';
import Toolbar from 'material-ui/Toolbar';
import Link from 'next/link';

import withRoot from '../styles/withRoot';
import Layout from '../components/Layout';
import MatchList from '../containers/MatchListContainer';
import initStore, { dispatchActions } from '../redux/store';
import { allMatchesSelector } from '../redux/reducers/entities';
import { loadMatchesAction } from '../redux/actions/entities';
import { setStatusAction } from '../redux/actions/status';
import { userOrLoginRouteAction } from '../redux/actions/auth';

import type { Match, Classes } from '../types';

const styles = theme => ({
  button: {
    margin: theme.spacing.unit,
    position: 'fixed',
    right: '24px',
    bottom: '24px',
  },
});

type Props = {
  classes: Classes,
  matches: Array<Match>,
  loadMatches: (?string) => Promise<{ empty: boolean }>,
};

type State = {
  loading: boolean,
  hasMore: boolean,
};

class Index extends React.Component<Props, State> {
  static async getInitialProps({ store, req, res, isServer }) {
    const actions = [loadMatchesAction(), userOrLoginRouteAction()];

    await dispatchActions(store.dispatch, isServer, req, res, actions);
  }

  state = {
    loading: false,
    hasMore: true,
  };

  onRefresh = async () => {
    const { loadMatches } = this.props;
    try {
      await loadMatches();
      this.setState({ loading: false, hasMore: true });
    } catch (e) {
      this.setState({ loading: false, hasMore: false });
    }
  };

  onLoadMore = async () => {
    const { loadMatches, matches } = this.props;

    this.setState({ loading: true });

    const lastElement = matches[matches.length - 1];

    const after = lastElement ? lastElement.createdAt : '';

    const newState = {
      loading: false,
      hasMore: true,
    };

    try {
      const result = await loadMatches(after);
      newState.hasMore = !result.empty;
    } catch (e) {
      newState.hasMore = false;
    } finally {
      console.log('has more: ' + newState.hasMore);
      this.setState(newState);
    }
  };

  render() {
    const { matches, classes } = this.props;
    const { loading, hasMore } = this.state;

    return (
      <Layout title="Matches">
        <div>
          <Toolbar>
            <Button color="secondary" onClick={this.onRefresh} variant="raised">
              Refresh
            </Button>
          </Toolbar>

          <MatchList
            matches={matches}
            onLoadMore={this.onLoadMore}
            loading={loading}
            hasMore={hasMore}
          />
        </div>

        <Tooltip title="Create new Match" className={classes.button}>
          <Button variant="fab" color="primary" aria-label="add">
            <Link prefetch href="/createMatch">
              <AddIcon />
            </Link>
          </Button>
        </Tooltip>
      </Layout>
    );
  }
}

function mapStateToProps(state) {
  const matches = allMatchesSelector(state);

  return {
    matches,
  };
}

const mapDispatchToProps = {
  loadMatches: loadMatchesAction,
  setStatus: setStatusAction,
};

export default withRedux(initStore, mapStateToProps, mapDispatchToProps)(
  withRoot(withStyles(styles)(Index)),
);
