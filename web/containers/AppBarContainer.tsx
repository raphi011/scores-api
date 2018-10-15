import Router from 'next/router';
import React from 'react';
import { connect } from 'react-redux';

import AppBar from '../components/AppBar';
import { logoutAction } from '../redux/actions/auth';
import { loginRouteSelector, userSelector } from '../redux/reducers/auth';

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

function mapStateToProps(state) {
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
