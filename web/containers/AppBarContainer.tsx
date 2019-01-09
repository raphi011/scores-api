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

interface State {
  bodyScrolled: boolean;
}

class AppBarContainer extends React.Component<Props, State> {
  state = {
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
  onLogout = async () => {
    const { logout } = this.props;

    await Router.push('/login');
    await logout();
  };

  render() {
    return (
      <AppBar
        {...this.props}
        bodyScrolled={this.state.bodyScrolled}
        onLogout={this.onLogout}
      />
    );
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
