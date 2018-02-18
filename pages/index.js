// @flow

import React from 'react';
import { withStyles } from 'material-ui/styles';
import Button from 'material-ui/Button';
import AddIcon from 'material-ui-icons/Add';
import Tooltip from 'material-ui/Tooltip';
import Toolbar from 'material-ui/Toolbar';
import Link from 'next/link';

import withAuth from '../containers/AuthContainer';
import Layout from '../components/Layout';
import MatchList from '../containers/MatchListContainer';
import { dispatchActions } from '../redux/store';
import { allMatchesSelector } from '../redux/reducers/entities';
import { loadMatchesAction } from '../redux/actions/entities';
import { setStatusAction } from '../redux/actions/status';

import type { Match, User, Classes } from '../types';

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
  user: User,
  matches: Array<Match>,
  loadMatches: (?string) => Promise<{ empty: boolean }>,
};

type State = {
  loading: boolean,
  hasMore: boolean,
};

class Index extends React.Component<Props, State> {
  static async getInitialProps({ store, req, res, isServer }) {
    const actions = [loadMatchesAction()];

    await dispatchActions(store.dispatch, isServer, req, res, actions);
  }

  static mapStateToProps(state) {
    const matches = allMatchesSelector(state);

    return {
      matches,
    };
  }

  static mapDispatchToProps = {
    loadMatches: loadMatchesAction,
    setStatus: setStatusAction,
  };

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
      this.setState(newState);
    }
  };

  render() {
    const { matches, user, classes } = this.props;
    const { loading, hasMore } = this.state;

    return (
      <Layout title="Matches">
        <div>
          <Toolbar>
            <Button color="primary" onClick={this.onRefresh}>
              Refresh
            </Button>
          </Toolbar>

          <MatchList
            highlightPlayerId={user.playerId}
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

export default withAuth(withStyles(styles)(Index));
