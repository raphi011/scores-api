import React from 'react';

import Router from 'next/router';
import { connect } from 'react-redux';

import { logoutAction } from '../redux/auth/actions';

interface Props {
  logout: () => Promise<void>;
}

class Logout extends React.Component<Props> {
  async componentDidMount() {
    const { logout } = this.props;

    await logout();
    Router.push('/login');
  }

  render() {
    return null;
  }
}

const mapDispatchToProps = dispatch => ({
  logout: () => dispatch(logoutAction()),
});

export default connect(
  null,
  mapDispatchToProps,
)(Logout);
