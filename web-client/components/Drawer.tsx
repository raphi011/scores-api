import React, { SyntheticEvent } from 'react';

import Link from 'next/link';

import Divider from '@material-ui/core/Divider';
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
import Typography from '@material-ui/core/Typography';
import LogoutIcon from '@material-ui/icons/ExitToApp';
import LadderIcon from '@material-ui/icons/LooksOne';
import SettingsIcon from '@material-ui/icons/SettingsRounded';
import TournamentIcon from '@material-ui/icons/Star';

import AdminOnly from '../containers/AdminOnly';

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
          <Link prefetch href="/">
            <ListItem button>
              <ListItemIcon>
                <TournamentIcon />
              </ListItemIcon>
              <ListItemText inset primary="Tournaments" />
            </ListItem>
          </Link>
          <Link prefetch href="/ladder">
            <ListItem button>
              <ListItemIcon>
                <LadderIcon />
              </ListItemIcon>
              <ListItemText inset primary="Ladder" />
            </ListItem>
          </Link>
          <AdminOnly>
            <Link href="/admin">
              <ListItem button>
                <ListItemIcon>
                  <SettingsIcon />
                </ListItemIcon>
                <ListItemText inset primary="Admin" />
              </ListItem>
            </Link>
          </AdminOnly>
        </List>
      </div>
    );

    const content = (
      <>
        {sideList}
        <div style={{ flex: 1 }} />
        <Divider />
        <Link href="/logout">
          <ListItem button>
            <ListItemIcon>
              <LogoutIcon />
            </ListItemIcon>
            <ListItemText inset primary="Logout" />
          </ListItem>
        </Link>
        <Typography align="center" variant="caption">
          {VERSION}
        </Typography>
      </>
    );

    if (mobile) {
      return (
        <Drawer open={open} variant="temporary" anchor="top" onClose={onClose}>
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
