import React from 'react';

import { connect } from 'react-redux';
import withWidth from '@material-ui/core/withWidth';

import AppBar from '../components/AppBar';
import { userSelector } from '../redux/auth/selectors';
import { Store } from '../redux/store';
import { User } from '../types';
import * as responsive from '../utils/responsive';
import { Breakpoint } from '@material-ui/core/styles/createBreakpoints';

interface Props {
  isLoggedIn: boolean;
  title: { text: string; href: string };
  user: User | null;
  width: Breakpoint;
}

interface State {
  bodyScrolled: boolean;
  drawerOpen: boolean;
}

class AppBarContainer extends React.Component<Props, State> {
  state = {
    anchorEl: null,
    bodyScrolled: false,
    drawerOpen: false,
  };

  componentDidMount() {
    window.addEventListener('scroll', this.onScroll);
  }

  onToggleDrawer = () => {
    this.setState({
      drawerOpen: !this.state.drawerOpen,
    });
  };

  onScroll = () => {
    const bodyScrolled = window.scrollY !== 0;

    if (bodyScrolled !== this.state.bodyScrolled) {
      this.setState({ bodyScrolled });
    }
  };

  onCloseDrawer = () => {
    this.setState({ drawerOpen: false });
  };

  componentWillUnmount() {
    window.removeEventListener('scroll', this.onScroll);
  }

  render() {
    const { width } = this.props;

    const isMobile = responsive.isMobile(width);

    return (
      <AppBar
        {...this.props}
        isMobile={isMobile}
        drawerOpen={isMobile && this.state.drawerOpen}
        onCloseDrawer={this.onCloseDrawer}
        onToggleDrawer={this.onToggleDrawer}
        bodyScrolled={this.state.bodyScrolled}
      />
    );
  }
}

function mapStateToProps(state: Store) {
  const user = userSelector(state);

  return {
    isLoggedIn: !!user,
    user,
  };
}

export default connect(mapStateToProps)(withWidth()(AppBarContainer));
