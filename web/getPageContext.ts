/* eslint-disable no-underscore-dangle */

import { SheetsRegistry } from 'jss';
import {
  createMuiTheme,
  createGenerateClassName,
} from '@material-ui/core/styles';
import blue from '@material-ui/core/colors/blue';
import teal from '@material-ui/core/colors/teal';

// A theme with custom primary and secondary color.
// It's optional.
const theme = createMuiTheme({
  palette: {
    primary: {
      dark: blue[700],
      light: blue[300],
      main: blue[500],
    },
    secondary: {
      dark: teal[700],
      light: teal[300],
      main: teal[500],
    },
  },
  typography: {
    useNextVariants: true,
  },
});

function createPageContext() {
  return {
    // The standard class name generator.
    generateClassName: createGenerateClassName(),
    // This is needed in order to deduplicate the injection of CSS in the page.
    sheetsManager: new Map(),
    sheetsRegistry: new SheetsRegistry(),
    theme,
    // This is needed in order to inject the critical CSS.
  };
}

export default function getPageContext() {
  // Make sure to create a new context for every server-side request so that data
  // isn't shared between connections (which would be bad).
  if (!process.browser) {
    return createPageContext();
  }

  // Reuse context on the client-side.
  if (!global.__INIT_MATERIAL_UI__) {
    global.__INIT_MATERIAL_UI__ = createPageContext();
  }

  return global.__INIT_MATERIAL_UI__;
}
