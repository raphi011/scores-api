// @flow
/* eslint-disable prefer-destructuring */

import React from 'react';
import Router from 'next/router';
import withRedux from 'next-redux-wrapper';

import initStore, { dispatchAction } from '../redux/store';
import { userSelector } from '../redux/reducers/auth';
import type { User } from '../types';
import withRoot from '../styles/withRoot';
import { userOrLoginRouteAction } from '../redux/actions/auth';

type Props = {
  isLoggedIn: boolean,
  user: User,
};

function withAuth(WrappedComponent) {
  class Auth extends React.Component<Props> {
    static async getInitialProps(ctx) {
      const { isServer, store, res, req } = ctx;

      let user;
      let path;
      let url;
      let isLoggedIn;
      let loginRoute = '';

      if (isServer) {
        const result = await dispatchAction(
          store.dispatch,
          isServer,
          req,
          res,
          userOrLoginRouteAction(),
        );

        user = result.response.user;
        loginRoute = result.response.loginRoute;
        path = req.path;
        url = req.url;
        isLoggedIn = !!user;
      } else {
        const authState = userSelector(store.getState());

        path = Router.pathname;
        url = Router.asPath;
        isLoggedIn = authState.isLoggedIn;
        user = authState.user;
      }

      if (path !== '/login' && !isLoggedIn) {
        const redir = url ? `?r=${encodeURIComponent(url)}` : '';

        if (isServer) {
          const host = req.headers.host;
          const loginUrl = `http://${host}/login${redir}`;
          res.writeHead(302, {
            Location: loginUrl,
          });
          res.end();
          res.finished = true;
        } else {
          Router.push(`/login${redir}`);
        }

        return {};
      }

      // All good, auth okay!
      if (WrappedComponent.getInitialProps) {
        const wrappedProps = await WrappedComponent.getInitialProps(ctx);

        const combinedProps = {
          ...wrappedProps,
          user,
          isLoggedIn,
        };

        return combinedProps;
      }

      return { user, isLoggedIn, loginRoute };
    }

    render() {
      return <WrappedComponent {...this.props} />;
    }
  }

  return withRedux(
    initStore,
    WrappedComponent.mapStateToProps,
    WrappedComponent.mapDispatchToProps,
  )(withRoot(Auth));
}

export default withAuth;
