import React, { ReactNode } from 'react';
import { withStyles } from '@material-ui/core/styles';

import Drawer from './Drawer';
import AppBar from '../containers/AppBarContainer';

import { Player, Classes } from '../types';

interface Props {
  title: string;
  children: ReactNode;
  userPlayer: Player;
  drawerOpen: boolean;
  onCloseDrawer: () => void;
  onOpenDrawer: () => void;
  classes: Classes;
}

const styles = theme => ({
  style: {
    backgroundColor: theme.palette.background.default,
    marginTop: '70px',
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
  <div className={classes.style}>
    <Drawer
      userPlayer={userPlayer}
      onClose={onCloseDrawer}
      onOpen={onOpenDrawer}
      open={drawerOpen}
    />
    <AppBar onOpenMenu={onOpenDrawer} title={title} />
    {children}
  </div>
);

export default withStyles(styles)(Layout);
