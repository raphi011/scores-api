import React from 'react';

import { Typography } from '@material-ui/core';
import withAuth from '../../containers/AuthContainer';
import Layout from '../../containers/LayoutContainer';
import {
  loadGroupAction,
  loadMatchesAction,
} from '../../redux/entities/actions';
import {
  groupSelector,
  matchesByGroupSelector,
} from '../../redux/entities/selectors';
import { setStatusAction } from '../../redux/status/actions';

import { Match /*, User, Group*/ } from '../../types';

interface Props {
  groupId: number;
  matches: Match[];

  loadMatches: (groupId: number, after?: string) => Promise<{ empty: boolean }>;
}

interface State {
  loading: boolean;
  hasMore: boolean;
}

class Index extends React.Component<Props, State> {
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
    // const { group } = this.props;

    return (
      <Layout title={{ text: 'Matches', href: '' }}>
        <div>
          {/* <Typography>{group.name}</Typography> */}
          <Typography>Last 5 Matches</Typography>
          <Typography>First 5 Players</Typography>
        </div>
      </Layout>
    );
  }
}

export default withAuth(Index);
