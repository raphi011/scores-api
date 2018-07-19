/* eslint-disable prefer-destructuring */

import React from 'react';
import { Element, ReactNode } from 'react';

import Router from 'next/router';
import { connect } from 'react-redux';

import { dispatchAction, dispatchActions } from '../redux/store';
import { userSelector } from '../redux/reducers/auth';
import { userOrLoginRouteAction } from '../redux/actions/auth';

import { User } from '../types';

type Props = {
  isLoggedIn: boolean,
  fromServer: boolean,
  user: User,
  dispatch: any,
};

type WrappedProps = {};

function withAuth<C: ComponentType<WrappedProps>>(
  WrappedComponent: C,
): Element<C> {
  class Auth extends React.Component<Props> {
    async componentDidMount() {
      const { fromServer, dispatch } = this.props;

      if (!WrappedComponent.buildActions || fromServer) {
        return;
      }

      const actions = WrappedComponent.buildActions(this.props);
      await dispatchActions(dispatch, actions, false);
    }

    async componentWillUpdate(nextProps) {
      if (
        !WrappedComponent.shouldComponentUpdate ||
        !WrappedComponent.buildActions ||
        !WrappedComponent.shouldComponentUpdate(this.props, nextProps)
      ) {
        return;
      }

      const { dispatch } = nextProps;

      const actions = WrappedComponent.buildActions(nextProps);

      await dispatchActions(dispatch, actions, false);
    }

    static async getInitialProps(ctx) {
      const { isServer, store, res, req, query } = ctx;

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
        user,
        isLoggedIn,
        loginRoute: '',
        fromServer: isServer,
        dispatch,
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
      if (WrappedComponent.getParameters) {
        const parameters = await WrappedComponent.getParameters(query);

        props = {
          ...props,
          ...parameters,
        };
      }

      // Execute these only on the server side to avoid waiting for
      // api calls before rendering the page
      if (isServer && WrappedComponent.buildActions) {
        const actions = WrappedComponent.buildActions(props);

        await dispatchActions(dispatch, actions, isServer, req, res);
      }

      return props;
    }

    render() {
      return <WrappedComponent {...this.props} />;
    }
  }

  return connect(
      WrappedComponent.mapStateToProps,
      WrappedComponent.mapDispatchToProps,
    )(Auth);
}

export default withAuth;
