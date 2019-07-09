import React from 'react';

import { ServerStyleSheets } from '@material-ui/styles';
import theme from '../styles/theme';
import Document, {
  Head,
  Main,
  NextScript,
  DocumentContext,
} from 'next/document';
import flush from 'styled-jsx/server';

interface Props {
  url: string;
}

class MyDocument extends Document<Props> {
  static getInitialProps = async (ctx: DocumentContext): Promise<any> => {
    const sheets = new ServerStyleSheets();
    const originalRenderPage = ctx.renderPage;

    ctx.renderPage = () =>
      originalRenderPage({
        enhanceApp: App => props => sheets.collect(<App {...props} />),
      });

    const initialProps = await Document.getInitialProps(ctx);

    return {
      ...initialProps,
      // Styles fragment is rendered after the app and page rendering finish.
      styles: (
        <>
          {sheets.getStyleElement()}
          {flush() || null}
        </>
      ),
    };
  };

  render() {
    const directives = [
      "default-src 'self'",
      "script-src 'self'",
      'img-src *',
      "child-src 'none'",
      "object-src 'none'",
      "font-src 'self' https://fonts.gstatic.com",
      "style-src 'self' 'unsafe-inline' https://fonts.googleapis.com",
      // 'report-uri /api/csp-violation-report', TODO .. this does not work via <meta> elements
    ];

    if (process.env.NODE_ENV === 'development') {
      // webpack's HMR needs this in development mode
      directives[1] = "script-src 'unsafe-eval' 'unsafe-inline' 'self'";
    }

    const csp = directives.join(';');

    return (
      <html lang="en" dir="ltr">
        <Head>
          <meta charSet="utf-8" />
          <meta
            name="viewport"
            content={
              'user-scalable=0, initial-scale=1, ' +
              'minimum-scale=1, width=device-width, height=device-height'
            }
          />
          <meta httpEquiv="Content-Security-Policy" content={csp} />
          <meta name="theme-color" content={theme.palette.primary.main} />
          <link
            rel="stylesheet"
            href="https://fonts.googleapis.com/css?family=Roboto:300,400,500"
          />
        </Head>
        <body style={{ overflowY: 'scroll' }}>
          <Main />
          <NextScript />
        </body>
      </html>
    );
  }
}

export default MyDocument;
