/* eslint-disable prefer-destructuring */

import React from 'react';

import Router from 'next/router';
import { connect } from 'react-redux';

import { userOrLoginRouteAction } from '../redux/actions/auth';
import { userSelector } from '../redux/reducers/auth';
import { dispatchAction, dispatchActions } from '../redux/store';

import { User } from '../types';

interface IProps {
  isLoggedIn: boolean;
  fromServer: boolean;
  user: User;
  dispatch: any;
}

const withAuth = Component => {
  class Auth extends React.Component<IProps> {
    static async getInitialProps(ctx) {
      try {
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
        return {};
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
      return <Component {...this.props} />;
    }
  }

  return connect(
    Component.mapStateToProps,
    Component.mapDispatchToProps,
  )(Auth);
};

export default withAuth;
