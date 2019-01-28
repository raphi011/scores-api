import {
  createStyles,
  Theme,
  WithStyles,
  withStyles,
} from '@material-ui/core/styles';
import React, { ReactNode } from 'react';

import AppBar from '../containers/AppBarContainer';
import Drawer from '../containers/DrawerContainer';

const styles = (theme: Theme) =>
  createStyles({
    appBar: {
      zIndex: theme.zIndex.drawer + 1,
    },
    content: {
      flexGrow: 1,
      padding: theme.spacing.unit * 3,

      [theme.breakpoints.down('xs')]: {
        padding: theme.spacing.unit * 1,
      }
    },
    root: {
      display: 'flex',
    },
    toolbar: theme.mixins.toolbar,
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
    <Drawer
      open={drawerOpen}
    />
    <main className={classes.content}>
      <div className={classes.toolbar} />
      {children}
    </main>
  </div>
);

export default withStyles(styles)(Layout);
