import React from 'react';

import Router from 'next/router';
import { connect } from 'react-redux';

import AppBar from '../components/AppBar';
import { logoutAction } from '../redux/auth/actions';
import { loginRouteSelector, userSelector } from '../redux/auth/selectors';
import { Store } from '../redux/store';
import { User } from '../types';

interface Props {
  isLoggedIn: boolean;
  loginRoute: string;
  title: { text: string; href: string };
  user: User;

  logout: () => void;
  onToggleDrawer: () => void;
}

interface State {
  bodyScrolled: boolean;
  anchorEl: Element;
}

class AppBarContainer extends React.Component<Props, State> {
  state = {
    anchorEl: null,
    bodyScrolled: false,
  };

  componentDidMount() {
    window.addEventListener('scroll', this.onScroll);
  }

  onScroll = () => {
    const bodyScrolled = window.scrollY !== 0;

    if (bodyScrolled !== this.state.bodyScrolled) {
      this.setState({ bodyScrolled });
    }
  };

  componentWillUnmount() {
    window.removeEventListener('scroll', this.onScroll);
  }

  handleClick = (event: React.SyntheticEvent) => {
    this.setState({ anchorEl: event.currentTarget });
  };

  handleClose = () => {
    this.setState({ anchorEl: null });
  };

  onLogout = async () => {
    const { logout } = this.props;

    this.handleClose();

    await Router.push('/login');
    await logout();
  };

  render() {
    return (
      <AppBar
        {...this.props}
        anchorEl={this.state.anchorEl}
        onMenuOpen={this.handleClick}
        onMenuClose={this.handleClose}
        bodyScrolled={this.state.bodyScrolled}
        onLogout={this.onLogout}
      />
    );
  }
}

function mapStateToProps(state: Store) {
  const { user, isLoggedIn } = userSelector(state);
  const loginRoute = loginRouteSelector(state);

  return {
    isLoggedIn,
    loginRoute,
    user,
  };
}

const mapDispatchToProps = {
  logout: logoutAction,
};

export default connect(
  mapStateToProps,
  mapDispatchToProps,
)(AppBarContainer);
