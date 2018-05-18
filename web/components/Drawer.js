// @flow

import React from 'react';
import { withStyles } from '@material-ui/core/styles';
import MDrawer from '@material-ui/core/Drawer';
import List, {
  ListItem,
  ListSubheader,
  ListItemIcon,
  ListItemText,
} from '@material-ui/core/List';
import GroupIcon from '@material-ui/icons/Group';
import Avatar from '@material-ui/core/Avatar';
import TournamentIcon from '@material-ui/icons/Star';
import SettingsIcon from '@material-ui/icons/Settings';
import HomeIcon from '@material-ui/icons/Home';
import LadderIcon from '@material-ui/icons/LooksOne';
import Link from 'next/link';

import type { Player } from '../types';

const styles = theme => ({
  list: {
    width: 250,
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
  nested: {
    paddingLeft: theme.spacing.unit * 4,
  },
});

type Props = {
  open: boolean,
  onOpen: Event => void,
  onClose: Event => void,
  userPlayer: Player,
  classes: Object,
};

function Drawer({ open, userPlayer, onClose, onOpen, classes }: Props) {
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

  return (
    <MDrawer open={open} onClose={onClose}>
      <div
        tabIndex={0}
        role="button"
        // onClick={onClose}
        // onKeyDown={onClose}
      >
        {sideList}
      </div>
    </MDrawer>
  );
}

export default withStyles(styles)(Drawer);
