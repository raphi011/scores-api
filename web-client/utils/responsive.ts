import { Breakpoint } from '@material-ui/core/styles/createBreakpoints';

export function isMobile(width: Breakpoint): boolean {
  return ['xs', 'sm'].indexOf(width) !== -1;
}
