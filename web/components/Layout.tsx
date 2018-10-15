import { createStyles, Theme, withStyles } from '@material-ui/core/styles';
import React, { ReactNode } from 'react';

import AppBar from '../containers/AppBarContainer';
import Drawer from './Drawer';

import { Player } from '../types';

interface Props {
  title: { text: string; href: string };
  children: ReactNode;
  userPlayer: Player;
  drawerOpen: boolean;
  onCloseDrawer: () => void;
  onOpenDrawer: () => void;
  classes: any;
}

const styles = (theme: Theme) =>
  createStyles({
    container: {
      flexGrow: 1,
      zIndex: 1,
      overflowY: 'hidden',
      position: 'relative',
      display: 'flex',
      width: '100%',
      backgroundColor: theme.palette.background.default,
    },
    content: {
      maxHeight: 'calc(100vh - 56px)',
      overflowY: 'scroll',
      scrollBehavior: 'smooth',
      '-webkit-overflow-scrolling': 'touch',
      flexGrow: 1,
    },
  });

const Layout = ({
  title,
  userPlayer,
  children,
  drawerOpen,
  onCloseDrawer,
  onOpenDrawer,
  classes,
}: Props) => (
  <div>
    <AppBar onOpenMenu={onOpenDrawer} title={title} />
    <div className={classes.container}>
      <Drawer
        userPlayer={userPlayer}
        onClose={onCloseDrawer}
        onOpen={onOpenDrawer}
        open={drawerOpen}
      />
      <main className={classes.content}>{children}</main>
    </div>
  </div>
);

export default withStyles(styles)(Layout);
