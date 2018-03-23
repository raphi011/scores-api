// @flow

import React from 'react';
import { connect } from 'react-redux';

import Layout from '../components/Layout';

import { playerSelector } from '../redux/reducers/entities';
import { userSelector } from '../redux/reducers/auth';

import type { Player } from '../types';

type Props = {
  title: string,
  children: React.Node,
  userPlayer: Player,
};

type State = {
  drawerOpen: boolean,
  groupOpen: { [number]: boolean },
};

class LayoutContainer extends React.Component<Props, State> {
  state = {
    drawerOpen: false,
    groupOpen: {},
  };

  onToggleGroup = (groupId: number) => {
    const groupOpen = !this.state.groupOpen[groupId];
    this.setState({
      groupOpen: {
        ...this.state.groupOpen,
        [groupId]: groupOpen,
      },
    });
  };

  onToggleDrawer = () => {
    this.setState({ drawerOpen: !this.state.drawerOpen });
  };

  onCloseDrawer = () => {
    this.setState({ drawerOpen: false });
  };

  render() {
    const { title, userPlayer, children } = this.props;
    return (
      <Layout
        userPlayer={userPlayer}
        onRequestClose={this.onCloseDrawer}
        onToggleGroup={this.onToggleGroup}
        groupOpen={this.state.groupOpen}
        drawerOpen={this.state.drawerOpen}
        onToggleDrawer={this.onToggleDrawer}
        title={title}
      >
        {children}
      </Layout>
    );
  }
}

function mapStateToProps(state) {
  const auth = userSelector(state);
  const userPlayer = playerSelector(state, auth.user.playerId);

  return { userPlayer };
}

export default connect(mapStateToProps)(LayoutContainer);
