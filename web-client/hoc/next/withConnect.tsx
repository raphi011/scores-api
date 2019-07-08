import React from 'react';

import { NextComponentType, NextPageContext } from 'next';
import Error from 'next/error';
import { connect } from 'react-redux';
import { Dispatch, Store as ReduxStore } from 'redux';

import { Store } from '../../redux/store';
import { dispatchActions } from '../../redux/actions';
import { userSelector } from '../../redux/auth/selectors';

interface Props {
  reduxStore: ReduxStore;
  fromServer: boolean;
  dispatch: Dispatch;
  initialReduxState: Store;
  error?: { responseCode: number };
}

export interface Context extends NextPageContext {
  store: ReduxStore;
}

export interface ClientContext extends Context {
  fromServer: boolean;
}

export default (Component: any): NextComponentType<Props> => {
  class WithConnect extends React.Component<Props> {
    static async getInitialProps(ctx: Context) {
      try {
        const { store, res, req } = ctx;

        const isServer = !!req;
        const initialReduxState = store.getState();

        let props = {
          initialReduxState,
          fromServer: isServer,
          reduxStore: store,
          user: userSelector(initialReduxState),
        };

        if (typeof Component.getInitialProps === 'function') {
          const initialProps = await Component.getInitialProps(ctx);

          props = {
            ...props,
            ...initialProps,
          };
        }

        // Execute these only on the server side to avoid waiting for
        // api calls before rendering the page
        if (isServer && Component.buildActions) {
          const actions = Component.buildActions(props);

          await dispatchActions(store.dispatch, actions, req, res);
        }

        return props;
      } catch (error) {
        return { error };
      }
    }

    async componentDidMount() {
      const { fromServer, reduxStore, error } = this.props;

      if (error || !Component.buildActions || fromServer) {
        return;
      }

      const actions = Component.buildActions(this.props);
      await dispatchActions(reduxStore.dispatch, actions);
    }

    async componentDidUpdate(nextProps: any, nextState: any) {
      if (
        nextProps.error ||
        !Component.buildActions ||
        !Component.shouldComponentFetch ||
        !Component.shouldComponentFetch(nextProps, nextState)
      ) {
        return;
      }

      const { reduxStore } = nextProps;

      const actions = Component.buildActions(nextProps);

      await dispatchActions(reduxStore.dispatch, actions);
    }

    render() {
      const { error, ...props } = this.props;

      if (error) {
        return <Error title="An error occured" statusCode={error.responseCode} />;
      }

      return <Component {...props} />;
    }
  }

  // @ts-ignore
  return connect(
    Component.mapStateToProps,
    Component.mapDispatchToProps,
  )(WithConnect);
};
