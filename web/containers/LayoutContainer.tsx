import React, { ReactNode } from 'react';

import { connect } from 'react-redux';

import Layout from '../components/Layout';
import { userSelector } from '../redux/auth/selectors';
import { Player } from '../types';

interface IProps {
  title: { text: string; href: string };
  children: ReactNode;
  userPlayer: Player;
}

interface IState {
  drawerOpen: boolean;
  groupOpen: { [key: number]: boolean };
}

class LayoutContainer extends React.Component<IProps, IState> {
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
  const {
    user: { player },
  } = userSelector(state);

  return { userPlayer: player };
}

export default connect(mapStateToProps)(LayoutContainer);
