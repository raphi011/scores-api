import React, { SyntheticEvent } from 'react';

import Drawer from '@material-ui/core/Drawer';
import List from '@material-ui/core/List';
import ListItem from '@material-ui/core/ListItem';
import ListItemIcon from '@material-ui/core/ListItemIcon';
import ListItemText from '@material-ui/core/ListItemText';
import {
  createStyles,
  Theme,
  WithStyles,
  withStyles,
} from '@material-ui/core/styles';
import SeasonIcon from '@material-ui/icons/CalendarToday';
import HomeIcon from '@material-ui/icons/Home';
import InfoIcon from '@material-ui/icons/Info';
import LadderIcon from '@material-ui/icons/LooksOne';
import SettingsIcon from '@material-ui/icons/SettingsRounded';
import TournamentIcon from '@material-ui/icons/Star';
import Link from 'next/link';
import AdminOnly from '../containers/AdminOnly';

import { Typography } from '@material-ui/core';

const drawerWidth = 300;

// global variable which is inserted by webpack.DefinePlugin()
declare var VERSION: string;

const styles = (theme: Theme) =>
  createStyles({
    header: {
      lineHeight: 'inherit',
      marginBottom: '5px',
      marginTop: '25px',
      textTransform: 'uppercase',
    },
    permanentDrawer: {
      flexShrink: 0,
      width: drawerWidth,
    },
    permanentDrawerPaper: {
      width: drawerWidth,
    },
    toolbar: theme.mixins.toolbar,
  });

interface Props extends WithStyles<typeof styles> {
  open: boolean;
  mobile: boolean;

  onClose?: (event: SyntheticEvent<{}>) => void;
}

export default withStyles(styles)(
  ({ mobile, onClose, open, classes }: Props) => {
    const sideList = (
      <div>
        <List>
          <ListItem button>
            <ListItemIcon>
              <InfoIcon />
            </ListItemIcon>
            <ListItemText inset primary="What's new?" />
          </ListItem>
          <Link href="/home">
            <ListItem button>
              <ListItemIcon>
                <HomeIcon />
              </ListItemIcon>
              <ListItemText inset primary="Home" />
            </ListItem>
          </Link>
          <AdminOnly>
            <Link href="/settings">
              <ListItem button>
                <ListItemIcon>
                  <SettingsIcon />
                </ListItemIcon>
                <ListItemText inset primary="Settings" />
              </ListItem>
            </Link>
          </AdminOnly>
          <Link prefetch href="/volleynet">
            <ListItem button>
              <ListItemIcon>
                <TournamentIcon />
              </ListItemIcon>
              <ListItemText inset primary="Tournaments" />
            </ListItem>
          </Link>
          <Link prefetch href="/volleynet/ranking">
            <ListItem button>
              <ListItemIcon>
                <LadderIcon />
              </ListItemIcon>
              <ListItemText inset primary="Rankings" />
            </ListItem>
          </Link>
          <ListItem button>
            <ListItemIcon>
              <SeasonIcon />
            </ListItemIcon>
            <ListItemText inset primary="My Season" />
          </ListItem>
        </List>
      </div>
    );

    const content = (
      <div tabIndex={0} role="button">
        {sideList}
        <Typography align="center" variant="caption">
          {VERSION}
        </Typography>
      </div>
    );

    if (mobile) {
      return (
        <Drawer
          open={open}
          variant="temporary"
          anchor="top"
          onClose={onClose}
        >
          {content}
        </Drawer>
      );
    }

    return (
      <Drawer
        open={open}
        variant="permanent"
        anchor="left"
        className={classes.permanentDrawer}
        classes={{
          paper: classes.permanentDrawerPaper,
        }}
      >
        <div className={classes.toolbar} />
        {content}
      </Drawer>
    );
  },
);
