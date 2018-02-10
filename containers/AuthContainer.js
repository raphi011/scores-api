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
        url = req.url;
        isLoggedIn = !!user;
      } else {
        const authState = userSelector(store.getState());

        url = Router.asPath;
        isLoggedIn = authState.isLoggedIn;
        user = authState.user;
      }

      if (!isLoggedIn) {
        if (!url.includes('/login')) {
          // redirect to '/login'
          const redir = url ? `?r=${encodeURIComponent(url)}` : '';

          if (isServer) {
            // TODO: improve this
            const protocol =
              process.env.NODE_ENV === 'development' ? 'http' : 'https';

            const host = req.headers.host;
            const loginUrl = `${protocol}://${host}/login${redir}`;
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

        return { user, isLoggedIn, loginRoute };
      }

      const props = {
        user,
        isLoggedIn,
      };

      // All good, auth okay!
      if (WrappedComponent.getInitialProps) {
        const wrappedProps = await WrappedComponent.getInitialProps(ctx);

        return {
          ...props,
          ...wrappedProps,
        };
      }

      return props;
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
