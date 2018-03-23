// @flow

import React from 'react';
import { withStyles } from 'material-ui/styles';
import Button from 'material-ui/Button';
import AddIcon from 'material-ui-icons/Add';
import Toolbar from 'material-ui/Toolbar';
import Link from 'next/link';

import withAuth from '../containers/AuthContainer';
import Layout from '../containers/LayoutContainer';
import MatchList from '../containers/MatchListContainer';
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
  static buildActions() {
    const actions = [loadMatchesAction()];

    return actions;
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
            <Button
              color="primary"
              variant="raised"
              size="small"
              onClick={this.onRefresh}
            >
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

        <Button
          variant="fab"
          color="primary"
          aria-label="add"
          className={classes.button}
        >
          <Link prefetch href="/createMatch">
            <AddIcon />
          </Link>
        </Button>
      </Layout>
    );
  }
}

export default withAuth(withStyles(styles)(Index));
