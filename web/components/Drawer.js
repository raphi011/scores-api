// @flow

import React from 'react';
import { withStyles } from 'material-ui/styles';
import MaterialDrawer from 'material-ui/Drawer';
import List, {
  ListItem,
  ListSubheader,
  ListItemIcon,
  ListItemText,
} from 'material-ui/List';
import Collapse from 'material-ui/transitions/Collapse';
import AddIcon from 'material-ui-icons/Add';
import Avatar from 'material-ui/Avatar';
import SettingsIcon from 'material-ui-icons/Settings';
import StatisticsIcon from 'material-ui-icons/ShowChart';
import FitnessCenterIcon from 'material-ui-icons/FitnessCenter';
import ExpandLessIcon from 'material-ui-icons/ExpandLess';
import ExpandMoreIcon from 'material-ui-icons/ExpandMore';
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
  onRequestClose: Event => void,
  userPlayer: Player,
  classes: Object,
};

function Drawer({
  open,
  userPlayer,
  groupOpen,
  onToggleGroup,
  onRequestClose,
  classes,
}: Props) {
  const { groups = [] } = userPlayer;

  const groupList = groups.map(g => (
    <GroupOptions
      key={g.id}
      open={false}
      onToggleOpen={onToggleGroup}
      open={groupOpen[g.id]}
      group={g}
      nestedClassName={classes.nested}
    />
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
        <Link href="">
          <ListItem button>
            <ListItemIcon>
              <SettingsIcon />
            </ListItemIcon>
            <ListItemText inset primary="Admin" />
          </ListItem>
        </Link>
        <ListSubheader className={classes.header}>My groups</ListSubheader>
        {groupList}
        <ListSubheader className={classes.header}>Other groups</ListSubheader>
      </List>
    </div>
  );

  return (
    <MaterialDrawer open={open} onClose={onRequestClose}>
      <div
        tabIndex={0}
        role="button"
        onClick={onRequestClose}
        onKeyDown={onRequestClose}
      >
        {sideList}
      </div>
    </MaterialDrawer>
  );
}

const GroupOptions = ({ group, open, onToggleOpen, nestedClassName }) => (
  <React.Fragment>
    <ListItem button onClick={() => onToggleOpen(group.id)}>
      <Avatar src={group.imageUrl} />
      <ListItemText primary={group.name} />
      {open ? <ExpandLessIcon /> : <ExpandMoreIcon />}
    </ListItem>
    <Collapse in={open} timeout="auto" unmountOnExit>
      <List component="div" disablePadding className={nestedClassName}>
        <Link
          prefetch
          href={{
            pathname: '/group/createMatch',
            query: { groupId: group.id },
          }}
        >
          <ListItem button>
            <ListItemIcon>
              <AddIcon />
            </ListItemIcon>
            <ListItemText primary="New Match" />
          </ListItem>
        </Link>
        <Link
          prefetch
          href={{ pathname: '/group', query: { groupId: group.id } }}
        >
          <ListItem button>
            <ListItemIcon>
              <FitnessCenterIcon />
            </ListItemIcon>
            <ListItemText primary="Matches" />
          </ListItem>
        </Link>
        <Link
          prefetch
          href={{ pathname: '/group/statistic', query: { groupId: group.id } }}
        >
          <ListItem button>
            <ListItemIcon>
              <StatisticsIcon />
            </ListItemIcon>
            <ListItemText primary="Statistics" />
          </ListItem>
        </Link>
      </List>
    </Collapse>
  </React.Fragment>
);

export default withStyles(styles)(Drawer);
