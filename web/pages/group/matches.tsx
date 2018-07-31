import React from 'react';
import { withStyles, Theme, createStyles } from '@material-ui/core/styles';
import Button from '@material-ui/core/Button';
import AddIcon from '@material-ui/icons/Add';
import Toolbar from '@material-ui/core/Toolbar';
import Link from 'next/link';

import withAuth from '../../containers/AuthContainer';
import Layout from '../../containers/LayoutContainer';
import MatchList from '../../containers/MatchListContainer';
import { matchesByGroupSelector } from '../../redux/reducers/entities';
import {
  loadGroupAction,
  loadMatchesAction,
} from '../../redux/actions/entities';
import { setStatusAction } from '../../redux/actions/status';

import { Match, User } from '../../types';

const styles = (theme: Theme) =>
  createStyles({
    button: {
      margin: theme.spacing.unit,
      position: 'fixed',
      right: '24px',
      bottom: '24px',
    },
  });

interface Props {
  groupId: number;
  user: User;
  matches: Array<Match>;
  loadMatches: (groupId: number, after?: string) => Promise<{ empty: boolean }>;
  classes: any;
}

interface State {
  loading: boolean;
  hasMore: boolean;
}

class Index extends React.Component<Props, State> {
  static getParameters(query) {
    let { groupId } = query;

    groupId = Number.parseInt(groupId, 10) || 0;

    return { groupId };
  }

  static buildActions({ groupId }) {
    const actions = [loadGroupAction(groupId)];

    return actions;
  }

  static mapStateToProps(state, ownProps) {
    const { groupId } = ownProps;
    const matches = matchesByGroupSelector(state, groupId);

    return { matches };
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
    const { loadMatches, groupId } = this.props;
    try {
      await loadMatches(groupId);
      this.setState({ loading: false, hasMore: true });
    } catch (e) {
      this.setState({ loading: false, hasMore: false });
    }
  };

  onLoadMore = async () => {
    const { loadMatches, groupId, matches } = this.props;

    this.setState({ loading: true });

    const lastElement = matches[matches.length - 1];

    const after = lastElement ? lastElement.createdAt : '';

    const newState = {
      loading: false,
      hasMore: true,
    };

    try {
      const result = await loadMatches(groupId, after);
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
      <Layout title={{ text: 'Matches', href: '' }}>
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
