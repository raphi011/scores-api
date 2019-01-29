import React from 'react';

import withAuth from '../containers/AuthContainer';

import Router from 'next/router';
import { logoutAction } from '../redux/auth/actions';

interface Props {
  logout: () => Promise<void>;
}

class Logout extends React.Component<Props> {
  static mapDispatchToProps = {
    logout: logoutAction,
  };

  async componentWillMount() {
    const { logout } = this.props;

    await logout();
  }

  componentDidMount() {
    Router.push('/login');
  }

  render() {
    return null;
  }
}

export default withAuth(Logout);
