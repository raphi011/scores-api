import { createMuiTheme } from '@material-ui/core/styles';
import { Palette } from '@material-ui/core/styles/createPalette';
import { TypographyOptions } from '@material-ui/core/styles/createTypography';

export const fontPalette: { [key: string]: string } = {
  900: '32px',
  800: '24px',
  700: '20px',
  600: '18px',
  500: '16px',
  400: '14px',
  300: '12px',
};

export const maxContentWidth = '1200px';

const grey = {
  900: 'hsl(210, 24%, 16%)',
  800: 'hsl(209, 20%, 25%)',
  700: 'hsl(209, 18%, 30%)',
  600: 'hsl(209, 14%, 37%)',
  500: 'hsl(211, 12%, 43%)',
  400: 'hsl(211, 10%, 53%)',
  300: 'hsl(211, 13%, 65%)',
  200: 'hsl(210, 16%, 82%)',
  100: 'hsl(214, 15%, 91%)',
  50: 'hsl(216, 33%, 97%)',
};

const primary = {
  900: 'hsl(204, 96%, 27%)',
  800: 'hsl(203, 87%, 34%)',
  700: 'hsl(202, 83%, 41%)',
  600: 'hsl(201, 79%, 46%)',
  500: 'hsl(199, 84%, 55%)',
  400: 'hsl(197, 92%, 61%)',
  300: 'hsl(196, 94%, 67%)',
  200: 'hsl(195, 97%, 75%)',
  100: 'hsl(195, 100%, 85%)',
  50: 'hsl(195, 100%, 95%)',
};

const spacingUnit = 8;

export default createMuiTheme({
  palette: {
    background: {
      default: 'hsl(160, 0%, 240%)',
    },
    primary,
    grey,
    contrastThreshold: 3,
  },
  overrides: {
    MuiAppBar: {
      colorDefault: {
        backgroundColor: 'hsl(160, 0%, 240%)',
      },
    },
    MuiCard: {
      root: {
        padding: 4 * spacingUnit,
      },
    },
    MuiDivider: {
      root: {
        color: grey[50],
        height: '1px',
      },
    },
  },
  shape: {
    borderRadius: 0,
  },
  spacing: {
    unit: spacingUnit,
  },
  typography: (palette: Palette): TypographyOptions => {
    return {
      body1: {
        fontSize: fontPalette[500],
      },
      body2: {
        fontSize: fontPalette[500],
        fontWeight: 300,
        color: palette.grey[700],
      },
      h1: {
        color: palette.grey[800],
        fontSize: fontPalette[900],
        fontWeight: 400,
      },
      h2: {
        fontSize: fontPalette[800],
      },
      h3: {
        color: palette.grey[400],
        fontSize: fontPalette[700],
        fontWeight: 500,
      },
      h4: {
        color: palette.grey[400],
        fontSize: fontPalette[400],
        fontWeight: 500,
        textTransform: 'uppercase',
      },
      subtitle1: {
        color: palette.grey[400],
        fontSize: fontPalette[500],
        fontWeight: 300,
      },
      subtitle2: {
        color: palette.grey[400],
        fontSize: fontPalette[400],
        fontWeight: 400,
      },
      useNextVariants: true,
    };
  },
});
