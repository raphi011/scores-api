// @flow

import React from 'react';
import { withStyles } from 'material-ui/styles';
import withRedux from 'next-redux-wrapper';
import Button from 'material-ui/Button';
import AddIcon from 'material-ui-icons/Add';
import Tooltip from 'material-ui/Tooltip';
import Link from 'next/link';

import withRoot from '../styles/withRoot';
import Layout from '../components/Layout';
import MatchList from '../containers/MatchListContainer';
import initStore, { dispatchActions } from '../redux/store';
import { matchesSelector } from '../redux/reducers/reducer';
import {
  loadMatchesAction,
  setStatusAction,
  userOrLoginRouteAction,
} from '../redux/actions/action';
import type { Match, Classes } from '../types';

const styles = theme => ({
  matchListContainer: {
    marginBottom: '70px',
  },
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
  loadMatches: string => Promise<any>,
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
      await loadMatches(after);
    } catch (e) {
      newState.hasMore = false;
    } finally {
      this.setState(newState);
    }
  };

  render() {
    const { matches, classes } = this.props;
    const { loading, hasMore } = this.state;

    return (
      <Layout title="Matches">
        <div className={classes.matchListContainer}>
          <MatchList
            matches={matches}
            onLoadMore={this.onLoadMore}
            loading={loading}
            hasMore={hasMore}
          />
        </div>
        <Tooltip title="Create new Match" className={classes.button}>
          <Button fab color="primary" aria-label="add">
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
  const matches = matchesSelector(state);

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
