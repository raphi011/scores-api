import React from 'react';

import { Context, redirectWithContext } from '../containers/AuthContainer';
import { logoutAction } from '../redux/auth/actions';

interface Props {
  logout: () => Promise<void>;
}

export default class Logout extends React.Component<Props> {
  static async getInitialProps(ctx: Context) {
    const { store } = ctx;

    await store.dispatch(logoutAction());
    redirectWithContext(ctx, '/login');

    return {};
  }

  render() {
    return null;
  }
}
