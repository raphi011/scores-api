// @flow weak

import React from 'react';
import Document, { Head, Main, NextScript } from 'next/document';
import JssProvider from 'react-jss/lib/JssProvider';
import getContext from '../styles/getContext';

class MyDocument extends Document {
  render() {
    return (
      <html lang="en" dir="ltr">
        <Head>
          <title>Scores</title>
          <meta charSet="utf-8" />
          <meta
            name="viewport"
            content={
              'user-scalable=0, initial-scale=1, ' +
              'minimum-scale=1, width=device-width, height=device-height'
            }
          />
          <link rel="manifest" href="/static/manifest.json" />
          <meta name="theme-color" content={this.props.stylesContext.theme.palette.primary[500]} />
          <link
            rel="stylesheet"
            href="https://fonts.googleapis.com/css?family=Roboto:300,400,500"
          />
        </Head>
        <body>
          <Main />
          <NextScript />
        </body>
      </html>
    );
  }
}

MyDocument.getInitialProps = ctx => {

  // Get the context to collected side effects.
  const context = getContext();
  const page = ctx.renderPage(Component => props => (
    <JssProvider registry={context.sheetsRegistry} jss={context.jss}>
      <Component {...props} />
    </JssProvider>
  ));

  return {
    ...page,
    stylesContext: context,
    styles: (
      <style
        id="jss-server-side"
        // eslint-disable-next-line react/no-danger
        dangerouslySetInnerHTML={{ __html: context.sheetsRegistry.toString() }}
      />
    ),
  };
};

export default MyDocument;
