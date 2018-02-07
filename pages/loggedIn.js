// @flow

import React from 'react';
import Router from 'next/router';
import withRedux from 'next-redux-wrapper';

import { loggedInAction } from '../redux/actions/auth';
import { setStatusAction } from '../redux/actions/status';
import initStore from '../redux/store';

class LoggedIn extends React.Component<null, null> {
  static async getInitialProps(props) {
    const { store, query } = props;

    const { username, error } = query;

    if (!error) {
      await store.dispatch(loggedInAction(username));
    } else {
      await store.dispatch(setStatusAction('User not found'));
    }
  }

  componentDidMount() {
    Router.replace('/');
  }

  render() {
    return null;
  }
}

export default withRedux(initStore)(LoggedIn);
