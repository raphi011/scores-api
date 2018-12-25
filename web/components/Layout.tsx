import {
  createStyles,
  Theme,
  WithStyles,
  withStyles,
} from '@material-ui/core/styles';
import React, { ReactNode } from 'react';

import AppBar from '../containers/AppBarContainer';
import Drawer from './Drawer';

import { Player } from '../types';

const styles = (theme: Theme) =>
  createStyles({
    appBar: {
      zIndex: theme.zIndex.drawer + 1,
    },
    content: {
      flexGrow: 1,
      padding: theme.spacing.unit * 3,
    },
    root: {
      display: 'flex',
    },
    toolbar: theme.mixins.toolbar,
  });

interface Props extends WithStyles<typeof styles> {
  title: { text: string; href: string };
  children: ReactNode;
  userPlayer: Player;
  drawerOpen: boolean;
  onCloseDrawer: () => void;
  onOpenDrawer: () => void;
}

const Layout = ({
  title,
  userPlayer,
  children,
  drawerOpen,
  onCloseDrawer,
  onOpenDrawer,
  classes,
}: Props) => (
  <div className={classes.root}>
    <AppBar onOpenMenu={onOpenDrawer} title={title} />
    <Drawer
      userPlayer={userPlayer}
      onClose={onCloseDrawer}
      onOpen={onOpenDrawer}
      open={drawerOpen}
    />
    <main className={classes.content}>
      <div className={classes.toolbar} />
      {children}
    </main>
  </div>
);

export default withStyles(styles)(Layout);
