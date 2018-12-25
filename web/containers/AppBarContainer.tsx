import React from 'react';

import Router from 'next/router';
import { connect } from 'react-redux';

import AppBar from '../components/AppBar';
import { logoutAction } from '../redux/auth/actions';
import { loginRouteSelector, userSelector } from '../redux/auth/selectors';
import { Store } from '../redux/store';

interface Props {
  isLoggedIn: boolean;
  loginRoute: string;
  onOpenMenu: () => void;
  title: { text: string; href: string };
  logout: () => void;
}

class AppBarContainer extends React.Component<Props> {
  onLogout = async () => {
    const { logout } = this.props;

    await Router.push('/login');
    await logout();
  };

  render() {
    return <AppBar {...this.props} onLogout={this.onLogout} />;
  }
}

function mapStateToProps(state: Store) {
  const { isLoggedIn } = userSelector(state);
  const loginRoute = loginRouteSelector(state);

  return {
    isLoggedIn,
    loginRoute,
  };
}

const mapDispatchToProps = {
  logout: logoutAction,
};

export default connect(
  mapStateToProps,
  mapDispatchToProps,
)(AppBarContainer);
