import React, { SyntheticEvent } from 'react';

import Avatar from '@material-ui/core/Avatar';
import SeasonIcon from '@material-ui/icons/CalendarToday';
import InfoIcon from '@material-ui/icons/Info';
import MDrawer from '@material-ui/core/Drawer';
// import Hidden from '@material-ui/core/Hidden';
import List from '@material-ui/core/List';
import ListItem from '@material-ui/core/ListItem';
import ListItemIcon from '@material-ui/core/ListItemIcon';
import ListItemText from '@material-ui/core/ListItemText';
import ListSubheader from '@material-ui/core/ListSubheader';
import {
  createStyles,
  Theme,
  WithStyles,
  withStyles,
} from '@material-ui/core/styles';
import HomeIcon from '@material-ui/icons/Home';
import LadderIcon from '@material-ui/icons/LooksOne';
import SettingsIcon from '@material-ui/icons/SettingsRounded';
import TournamentIcon from '@material-ui/icons/Star';
import Link from 'next/link';
import AdminOnly from '../containers/AdminOnly';

import { Typography } from '@material-ui/core';
import { Player } from '../types';

const drawerWidth = 300;

// global variable which is inserted by webpack.DefinePlugin()
declare var VERSION: string;

const styles = (theme: Theme) =>
  createStyles({
    drawer: {
      flexShrink: 0,
      width: drawerWidth,
    },
    drawerPaper: {
      width: drawerWidth,
    },
    toolbar: theme.mixins.toolbar,

    // drawerPaper: {
    //   width: drawerWidth,
    //   [theme.breakpoints.up('md')]: {
    //     position: 'relative',
    //   },
    // },
    header: {
      lineHeight: 'inherit',
      marginBottom: '5px',
      marginTop: '25px',
      textTransform: 'uppercase',
    },
    // list: {
    //   background: theme.palette.background.paper,
    // },
    // listFull: {
    //   width: 'auto',
    // },
    // nested: {
    //   paddingLeft: theme.spacing.unit * 4,
    // },
  });

interface Props extends WithStyles<typeof styles> {
  open: boolean;
  onOpen: (event: SyntheticEvent<{}>) => void;
  onClose: (event: SyntheticEvent<{}>) => void;
  userPlayer: Player;
}

function Drawer({ open, userPlayer, onClose, classes }: Props) {
  const sideList = (
    <div className={classes.list}>
      <List>
        <ListItem button>
          <Avatar src={userPlayer.profileImageUrl} />
          <ListItemText inset primary={userPlayer.name} />
        </ListItem>
        {/* </Link> */}
        <ListSubheader className={classes.header}>Navigation</ListSubheader>
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
        <ListSubheader className={classes.header}>Volleynet</ListSubheader>
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

  return (
    // <>
    //   <Hidden mdUp>
    //     <MDrawer
    //       open={open}
    //       onClose={onClose}
    //       ModalProps={{ keepMounted: true }}
    //     >
    //       {content}
    //     </MDrawer>
    //   </Hidden>
    // <Hidden smDown implementation="css">
    <MDrawer
      className={classes.drawer}
      variant="permanent"
      classes={{
        paper: classes.drawerPaper,
      }}
    >
      <div className={classes.toolbar} />
      {content}
    </MDrawer>
    //    </Hidden>
    // </>
  );
}

export default withStyles(styles)(Drawer);
