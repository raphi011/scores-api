import React, { ReactNode } from 'react';
import { connect } from 'react-redux';

import Layout from '../components/Layout';

import { playerSelector } from '../redux/reducers/entities';
import { userSelector } from '../redux/reducers/auth';

import { Player } from '../types';

interface Props {
  title: string;
  children: ReactNode;
  userPlayer: Player;
}

interface State {
  drawerOpen: boolean;
  groupOpen: { [key: number]: boolean };
}

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

  onOpenDrawer = () => {
    this.setState({ drawerOpen: true });
  };

  onCloseDrawer = () => {
    this.setState({ drawerOpen: false });
  };

  render() {
    const { title, userPlayer, children } = this.props;
    return (
      <Layout
        userPlayer={userPlayer}
        onCloseDrawer={this.onCloseDrawer}
        onOpenDrawer={this.onOpenDrawer}
        drawerOpen={this.state.drawerOpen}
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
