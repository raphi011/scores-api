import React from 'react';

// import Button from '@material-ui/core/Button';
import { createStyles, Theme, withStyles } from '@material-ui/core/styles';
// import Toolbar from '@material-ui/core/Toolbar';
// import AddIcon from '@material-ui/icons/Add';
// import Link from 'next/link';

import { Typography } from '@material-ui/core';
import withAuth from '../../containers/AuthContainer';
import Layout from '../../containers/LayoutContainer';
// import MatchList from '../../containers/MatchListContainer';
import {
  loadGroupAction,
  loadMatchesAction,
} from '../../redux/entities/actions';
import {
  // groupPlayersSelector,
  groupSelector,
  matchesByGroupSelector,
} from '../../redux/entities/selectors';
import { setStatusAction } from '../../redux/status/actions';

import { Match /*, User, Group*/ } from '../../types';

const styles = (theme: Theme) => createStyles({});

interface IProps {
  groupId: number;
  // user: User;
  matches: Match[];
  // group: Group;
  loadMatches: (groupId: number, after?: string) => Promise<{ empty: boolean }>;
  classes: any;
}

interface IState {
  loading: boolean;
  hasMore: boolean;
}

class Index extends React.Component<IProps, IState> {
  static mapDispatchToProps = {
    loadMatches: loadMatchesAction,
    setStatus: setStatusAction,
  };
  static getParameters(query) {
    let { groupId } = query;

    groupId = Number(groupId) || 0;

    return { groupId };
  }

  static buildActions({ groupId }) {
    const actions = [loadGroupAction(groupId)];

    return actions;
  }

  static mapStateToProps(state, ownProps) {
    const { groupId } = ownProps;

    const matches = matchesByGroupSelector(state, groupId);
    const group = groupSelector(state, groupId);

    return { matches, group };
  }

  state = {
    hasMore: true,
    loading: false,
  };

  render() {
    const { matches, user, group, classes } = this.props;
    const { loading, hasMore } = this.state;

    return (
      <Layout title={{ text: 'Matches', href: '' }}>
        <div>
          <Typography>{group.name}</Typography>
          <Typography>Last 5 Matches</Typography>
          <Typography>First 5 Players</Typography>
          {/* <Toolbar>
            <Button
              color="primary"
              variant="contained"
              size="small"
              onClick={this.onRefresh}
            >
              Refresh
            </Button>
          </Toolbar> */}

          {/* <MatchList
            highlightPlayerId={user.playerId}
            matches={matches}
            onLoadMore={this.onLoadMore}
            loading={loading}
            hasMore={hasMore}
          /> */}
        </div>

        {/* <Button
          variant="fab"
          color="primary"
          aria-label="add"
          className={classes.button}
        >
          <Link prefetch href="/createMatch">
            <AddIcon />
          </Link>
        </Button> */}
      </Layout>
    );
  }
}

export default withAuth(withStyles(styles)(Index));
