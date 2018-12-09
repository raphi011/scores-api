/* eslint-disable prefer-destructuring */

import React from 'react';

import { NextComponentClass, NextContext } from 'next';
import Error from 'next/error';
import Router from 'next/router';
import { connect } from 'react-redux';
import { Dispatch, Store } from 'redux';

import { userOrLoginRouteAction } from '../redux/auth/actions';
import { userSelector } from '../redux/auth/selectors';
import { dispatchAction, dispatchActions } from '../redux/store';

type Props = {
  store: Store;
  fromServer: boolean;
  dispatch: Dispatch;
  error: any;
};

interface Context extends NextContext {
  store: Store;
}

export default (Component): NextComponentClass<Props> => {
  class Auth extends React.Component<Props> {
    static async getInitialProps(ctx: Context) {
      try {
        const { store, res, req, query } = ctx;

        const isServer = !!req;

        const dispatch = store.dispatch;

        let user;
        let url;
        let isLoggedIn;
        let loginRoute = '';

        if (isServer) {
          const result = await dispatchAction(
            dispatch,
            userOrLoginRouteAction(),
            isServer,
            req,
            res,
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

        let props = {
          dispatch,
          fromServer: isServer,
          isLoggedIn,
          loginRoute: '',
          user,
        };

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

          props.loginRoute = loginRoute;
        }

        // All good, return props!
        if (Component.getParameters) {
          const parameters = await Component.getParameters(query);

          props = {
            ...props,
            ...parameters,
          };
        }

        // Execute these only on the server side to avoid waiting for
        // api calls before rendering the page
        if (isServer && Component.buildActions) {
          const actions = Component.buildActions(props);

          await dispatchActions(dispatch, actions, isServer, req, res);
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
      await dispatchActions(dispatch, actions, false);
    }

    async componentWillUpdate(nextProps) {
      if (
        !Component.shouldComponentUpdate ||
        !Component.buildActions ||
        !Component.shouldComponentUpdate(this.props, nextProps)
      ) {
        return;
      }

      const { dispatch } = nextProps;

      const actions = Component.buildActions(nextProps);

      await dispatchActions(dispatch, actions, false);
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
