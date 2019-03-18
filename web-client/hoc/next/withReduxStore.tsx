import React from 'react';
import initializeStore, { Store } from '../../redux/store';
import { Context } from './withConnect';

const isServer = typeof window === 'undefined';
const __NEXT_REDUX_STORE__ = '__NEXT_REDUX_STORE__';

declare global {
  interface Window {
    __NEXT_REDUX_STORE__: Store;
  }
}

function getOrCreateStore(initialState?: Store) {
  // Always make a new store if server, otherwise state is shared between requests
  if (isServer) {
    return initializeStore(initialState);
  }

  // Create store if unavailable on the client and set it on the window object
  if (!window[__NEXT_REDUX_STORE__]) {
    window[__NEXT_REDUX_STORE__] = initializeStore(initialState);
  }
  return window[__NEXT_REDUX_STORE__];
}

interface Props {
  initialReduxState: Store;
}

export default (App: React.ComponentType) => {
  return class AppWithRedux extends React.Component<Props> {
    static async getInitialProps(appContext: Context) {
      // Get or Create the store with `undefined` as initialState
      // This allows you to set a custom default initialState
      const store = getOrCreateStore();

      // Provide the store to getInitialProps of pages
      // @ts-ignore
      appContext.ctx.store = store;

      let appProps = {};
      // @ts-ignore
      if (typeof App.getInitialProps === 'function') {
        // @ts-ignore
        appProps = await App.getInitialProps(appContext);
      }

      return {
        ...appProps,
        initialReduxState: store.getState(),
      };
    }

    store: Store;

    constructor(props: Props) {
      super(props);
      this.store = getOrCreateStore(props.initialReduxState);
    }

    render() {
      // @ts-ignore
      return <App {...this.props} store={this.store} />;
    }
  };
};
