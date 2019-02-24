/* eslint-disable no-underscore-dangle */

import { createGenerateClassName, Theme } from '@material-ui/core/styles';
import { SheetsRegistry, GenerateClassName } from 'jss';
import theme from './styles/theme';

export interface PageContext {
  generateClassName: GenerateClassName;
  sheetsManager: any; // Map<Theme, SheetManagerTheme>
  sheetsRegistry: SheetsRegistry;
  theme: Theme;
}

function createPageContext(): PageContext {
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

declare const process: any;
declare const global: any;

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
