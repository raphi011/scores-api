import {
  createStyles,
  Theme,
  WithStyles,
  withStyles,
} from '@material-ui/core/styles';
import React, { ReactNode } from 'react';

import AppBar from '../containers/AppBarContainer';
import Drawer from '../containers/DrawerContainer';
import { maxContentWidth } from '../styles/theme';

const styles = (theme: Theme) =>
  createStyles({
    appBar: {
      zIndex: theme.zIndex.drawer + 1,
    },
    content: {
      flexGrow: 1,
      margin: '80px auto 0 auto',
      maxWidth: maxContentWidth,
      padding: theme.spacing.unit * 3,

      [theme.breakpoints.down('xs')]: {
        padding: theme.spacing.unit * 1,
      },
    },
    root: {
      display: 'flex',
    },
  });

interface Props extends WithStyles<typeof styles> {
  title: { text: string; href: string };
  children: ReactNode;
  drawerOpen: boolean;

  onToggleDrawer: () => void;
}

const Layout = ({
  title,
  children,
  drawerOpen,
  onToggleDrawer,
  classes,
}: Props) => (
  <div className={classes.root}>
    <AppBar onToggleDrawer={onToggleDrawer} title={title} />
    <Drawer open={drawerOpen} />
    <main className={classes.content}>{children}</main>
  </div>
);

export default withStyles(styles)(Layout);
