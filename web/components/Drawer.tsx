import React from 'react';
import { withStyles, createStyles, Theme } from '@material-ui/core/styles';
import MDrawer from '@material-ui/core/Drawer';
import Hidden from '@material-ui/core/Hidden';
import List from '@material-ui/core/List';
import ListItem from '@material-ui/core/ListItem';
import ListSubheader from '@material-ui/core/ListSubheader';
import ListItemIcon from '@material-ui/core/ListItemIcon';
import ListItemText from '@material-ui/core/ListItemText';
import GroupIcon from '@material-ui/icons/Group';
import Avatar from '@material-ui/core/Avatar';
import TournamentIcon from '@material-ui/icons/Star';
import SettingsIcon from '@material-ui/icons/SettingsRounded';
import HomeIcon from '@material-ui/icons/Home';
import LadderIcon from '@material-ui/icons/LooksOne';
import Link from 'next/link';

import { Player } from '../types';
import { Typography } from '@material-ui/core';

const drawerWidth = 300;

const styles = (theme: Theme) =>
  createStyles({
    list: {
      background: theme.palette.background.paper,
    },
    listFull: {
      width: 'auto',
    },
    header: {
      marginTop: '25px',
      marginBottom: '5px',
      lineHeight: 'inherit',
      textTransform: 'uppercase',
    },
    drawerPaper: {
      width: drawerWidth,
      [theme.breakpoints.up('md')]: {
        position: 'relative',
      },
    },
    nested: {
      paddingLeft: theme.spacing.unit * 4,
    },
  });

interface Props {
  open: boolean;
  onOpen: (Event) => void;
  onClose: (Event) => void;
  userPlayer: Player;
  classes: any;
}

function Drawer({ open, userPlayer, onClose, classes }: Props) {
  const { groups = [] } = userPlayer;

  const groupList = groups.map(g => (
    <Link
      prefetch
      key={g.id}
      href={{ pathname: '/group/statistic', query: { groupId: g.id } }}
    >
      <ListItem button>
        <ListItemIcon>
          <GroupIcon />
        </ListItemIcon>
        <ListItemText primary={g.name} />
      </ListItem>
    </Link>
  ));

  const sideList = (
    <div className={classes.list}>
      <List>
        <Link
          prefetch
          href={{ pathname: '/player', query: { id: userPlayer.id } }}
        >
          <ListItem button>
            <Avatar src={userPlayer.profileImageUrl} />
            <ListItemText inset primary={userPlayer.name} />
          </ListItem>
        </Link>
        <ListSubheader className={classes.header}>Navigation</ListSubheader>
        <Link href="/home">
          <ListItem button>
            <ListItemIcon>
              <HomeIcon />
            </ListItemIcon>
            <ListItemText inset primary="Home" />
          </ListItem>
        </Link>
        <ListItem button>
          <ListItemIcon>
            <SettingsIcon />
          </ListItemIcon>
          <ListItemText inset primary="Settings" />
        </ListItem>
        <ListSubheader className={classes.header}>Volleynet</ListSubheader>
        <Link href="/volleynet">
          <ListItem button>
            <ListItemIcon>
              <TournamentIcon />
            </ListItemIcon>
            <ListItemText inset primary="Tournaments" />
          </ListItem>
        </Link>
        <Link href="/volleynet/ranking">
          <ListItem button>
            <ListItemIcon>
              <LadderIcon />
            </ListItemIcon>
            <ListItemText inset primary="Rankings" />
          </ListItem>
        </Link>
        <ListSubheader className={classes.header}>My groups</ListSubheader>
        {groupList}
        <ListSubheader className={classes.header}>Other groups</ListSubheader>
        <ListItem button>
          <ListItemText primary="-" />
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
    <>
      <Hidden mdUp>
        <MDrawer
          open={open}
          onClose={onClose}
          ModalProps={{ keepMounted: true }}
        >
          {content}
        </MDrawer>
      </Hidden>
      <Hidden smDown implementation="css">
        <MDrawer
          variant="permanent"
          open
          classes={{
            paper: classes.drawerPaper,
          }}
        >
          {content}
        </MDrawer>
      </Hidden>
    </>
  );
}

export default withStyles(styles)(Drawer);
