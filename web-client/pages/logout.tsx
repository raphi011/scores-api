import React from 'react';

import Router from 'next/router';
import { Store } from 'redux';

import { Context } from '../containers/AuthContainer';
import { logoutAction } from '../redux/auth/actions';

interface Props {
  store: Store;
}

export default class Logout extends React.Component<Props> {
  static async getInitialProps(ctx: Context) {
    const { store } = ctx;

    return { store };
  }

  async componentDidMount() {
    const { store } = this.props;

    await store.dispatch(logoutAction());
    Router.push('/login');
  }

  render() {
    return null;
  }
}
