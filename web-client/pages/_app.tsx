import React from 'react';

import App, { Container } from 'next/app';
import { Provider } from 'react-redux';
import { Store } from 'redux';

import CssBaseline from '@material-ui/core/CssBaseline';
import { ThemeProvider } from '@material-ui/styles';

import Snackbar from '../containers/SnackbarContainer';
import theme from '../styles/theme';
import withReduxStore from '../hoc/next/withReduxStore';

interface Props {
  store: Store;
}

class MyApp extends App<Props> {
  componentDidMount() {
    // Remove the server-side injected CSS.
    const jssStyles = document.querySelector('#jss-server-side');
    if (jssStyles && jssStyles.parentNode) {
      jssStyles.parentNode.removeChild(jssStyles);
    }
  }

  render() {
    const { Component, pageProps, store } = this.props;
    return (
      <Container>
        <ThemeProvider theme={theme}>
          {/* CssBaseline kickstart an elegant, consistent, and simple baseline to build upon. */}
          <CssBaseline />
          {/* Pass pageContext to the _document though the renderPage enhancer
                to render collected styles on server side. */}
          <Provider store={store}>
            <>
              <Component {...pageProps} />
              <Snackbar />
            </>
          </Provider>
        </ThemeProvider>
      </Container>
    );
  }
}

export default withReduxStore(MyApp);
