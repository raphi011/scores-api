import React from 'react';

import { NextComponentClass, NextContext } from 'next';
import Error from 'next/error';
import Router from 'next/router';
import { connect } from 'react-redux';
import { Dispatch, Store, Action } from 'redux';

import { dispatchAction, dispatchActions } from '../redux/actions';
import { userOrLoginRouteAction } from '../redux/auth/actions';
import { userSelector } from '../redux/auth/selectors';

interface Props {
  store: Store;
  fromServer: boolean;
  dispatch: Dispatch;
  error?: { responseCode: number };
}

export interface Context extends NextContext {
  store: Store;
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
        const { store, res, req, query, pathname, asPath } = ctx;

        const isServer = !!req;

        const dispatch = store.dispatch;

        let isLoggedIn: boolean;

        if (isServer) {
          const result = await dispatchWithContext(
            ctx,
            userOrLoginRouteAction(),
          );

          isLoggedIn = !!result.response.user;
        } else {
          const authState = userSelector(store.getState());

          isLoggedIn = authState.isLoggedIn;
        }

        let props = {
          dispatch,
          fromServer: isServer,
          isLoggedIn,
        };

        if (!isLoggedIn) {
          if (pathname !== '/login') {
            const redir = asPath ? `?r=${encodeURIComponent(asPath)}` : '';

            const path = `/login${redir}`;

            redirectWithContext(ctx, path);

            return {};
          }
        }

        // All good, return props!
        if (Component.getParameters) {
          const parameters = await Component.getParameters(query);

          props = {
            ...props,
            ...parameters,
          };
        }

        if (Component.getInitialProps) {
          const initialProps = await Component.getInitialProps(ctx);

          props = {
            ...print,
            ...initialProps,
          };
        }

        // Execute these only on the server side to avoid waiting for
        // api calls before rendering the page
        if (isServer && Component.buildActions) {
          const actions = Component.buildActions(props);

          await dispatchActions(dispatch, actions, req, res);
        }

        return props;
      } catch (e) {
        return { error: e };
      }
    }

    async componentDidMount() {
      const { fromServer, dispatch } = this.props;

      if (!Component.buildActions || fromServer) {
        return;
      }

      const actions = Component.buildActions(this.props);
      await dispatchActions(dispatch, actions);
    }

    async componentDidUpdate(nextProps: any, nextState: any) {
      if (
        !Component.shouldComponentUpdate ||
        !Component.buildActions ||
        !Component.shouldComponentUpdate(nextProps, nextState)
      ) {
        return;
      }

      const { dispatch } = nextProps;

      const actions = Component.buildActions(nextProps);

      await dispatchActions(dispatch, actions);
    }

    render() {
      const { error, ...props } = this.props;

      if (error) {
        return <Error statusCode={error.responseCode} />;
      }
      return <Component {...props} />;
    }
  }

  return connect(
    Component.mapStateToProps,
    Component.mapDispatchToProps,
  )(Auth);
};
