import { createStyles, Theme, withStyles } from '@material-ui/core/styles';
import React, { ReactNode } from 'react';

import AppBar from '../containers/AppBarContainer';
import Drawer from './Drawer';

import { Player } from '../types';

interface IProps {
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
      backgroundColor: theme.palette.background.default,
      display: 'flex',
      flexGrow: 1,
      overflowY: 'hidden',
      position: 'relative',
      width: '100%',
      zIndex: 1,
    },
    content: {
      '-webkit-overflow-scrolling': 'touch',
      flexGrow: 1,
      maxHeight: 'calc(100vh - 56px)',
      overflowY: 'scroll',
      scrollBehavior: 'smooth',
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
}: IProps) => (
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
