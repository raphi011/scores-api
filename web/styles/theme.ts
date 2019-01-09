import { Color } from '@material-ui/core';
import { PaletteColorOptions } from '@material-ui/core/styles/createPalette';

// Palette 1 of Refactoring UI Palette
export const primary: PaletteColorOptions = {
  900: 'hsl(184, 91%, 17%)',
  800: 'hsl(185, 84%, 25%)',
  700: 'hsl(185, 81%, 29%)',
  600: 'hsl(184, 77%, 34%)',
  500: 'hsl(185, 62%, 45%)',
  400: 'hsl(185, 57%, 50%)',
  300: 'hsl(184, 65%, 59%)',
  200: 'hsl(184, 80%, 74%)',
  100: 'hsl(185, 94%, 87%)',
  50: 'hsl(186, 100%, 94%)',
  A100: 'hsl(224, 67%, 76%)',
  A200: 'hsl(330, 70%, 36%)',
  A400: 'hsl(360, 77%, 78%)',
  A700: 'hsl(45, 86%, 81%)',
};

export const secondary: PaletteColorOptions = {
  900: 'hsl(227, 42%, 51%)', // tournament status open
  800: 'hsl(360, 64%, 55%)', // tournament status closed
  700: 'hsl(184, 77%, 34%)', // tournament status done

  // copied from primary .. fix me
  A100: 'hsl(224, 67%, 76%)',
  A200: 'hsl(330, 70%, 36%)',
  A400: 'hsl(360, 77%, 78%)',
  A700: 'hsl(45, 86%, 81%)',
};

export const grey: Partial<Color> = {
  900: 'hsl(209, 61%, 16%)',
  800: 'hsl(211, 39%, 23%)',
  700: 'hsl(209, 34%, 30%)',
  600: 'hsl(209, 28%, 39%)',
  500: 'hsl(210, 22%, 49%)',
  400: 'hsl(209, 23%, 60%)',
  300: 'hsl(211, 27%, 70%)',
  200: 'hsl(210, 31%, 80%)',
  100: 'hsl(212, 33%, 89%)',
  50: 'hsl(210, 36%, 96%)',
};

export const background = {
  default: grey[50],
};
