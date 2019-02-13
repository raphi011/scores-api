import React, { ReactNode } from 'react';

import { connect } from 'react-redux';

import Layout from '../components/Layout';

interface Props {
  title: { text: string; href: string };
  children: ReactNode;
}

interface State {
  drawerOpen: boolean;
}

class LayoutContainer extends React.Component<Props, State> {
  state = {
    drawerOpen: false,
  };

  onToggleDrawer = () => {
    this.setState({
      drawerOpen: !this.state.drawerOpen,
    });
  };

  render() {
    const { title, children } = this.props;

    return (
      <Layout
        onToggleDrawer={this.onToggleDrawer}
        drawerOpen={this.state.drawerOpen}
        title={title}
      >
        {children}
      </Layout>
    );
  }
}

function mapStateToProps() {
  return {};
}

export default connect(mapStateToProps)(LayoutContainer);
