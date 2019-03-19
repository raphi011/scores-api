import React from 'react';

import { NextComponentClass } from 'next';
import Router from 'next/router';
import { Dispatch, Store, Action } from 'redux';

import { dispatchAction } from '../../redux/actions';
import { userOrLoginRouteAction } from '../../redux/auth/actions';
import { userSelector } from '../../redux/auth/selectors';
import { Context } from './withConnect';

interface Props {
  store: Store;
  dispatch: Dispatch;
  error?: { responseCode: number };
}

export async function dispatchWithContext(ctx: Context, action: Action) {
  const { store, res, req } = ctx;

  const dispatch = store.dispatch;

  return await dispatchAction(dispatch, action, req, res);
}

export function redirectWithContext(ctx: Context, path: string) {
  const { res, req } = ctx;

  if (req && res) {
    const protocol = 'https';
    const host = req.headers.host;
    const loginUrl = `${protocol}://${host}${path}`;
    res.writeHead(302, {
      Location: loginUrl,
    });
    res.end();
    res.finished = true;
  } else {
    Router.push(path);
  }
}

export default (Component: any): NextComponentClass<Props> => {
  class Auth extends React.Component<Props> {
    static async getInitialProps(ctx: Context) {
      try {
        const { store, req, asPath } = ctx;

        const isServer = !!req;

        let isLoggedIn: boolean;

        if (isServer) {
          const result = await dispatchWithContext(
            ctx,
            userOrLoginRouteAction(),
          );

          isLoggedIn = !!result.response.user;
        } else {
          const user = userSelector(store.getState());

          isLoggedIn = !!user;
        }

        if (!isLoggedIn) {
          const redir =
            asPath && asPath !== '/' ? `?r=${encodeURIComponent(asPath)}` : '';

          const path = `/login${redir}`;

          redirectWithContext(ctx, path);

          return {};
        }

        if (typeof Component.getInitialProps === 'function') {
          return await Component.getInitialProps(ctx);
        }

        return {};
      } catch (e) {
        return { error: e };
      }
    }

    render() {
      return <Component {...this.props} />;
    }
  }

  // @ts-ignore
  return Auth;
};
