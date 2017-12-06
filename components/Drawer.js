import React from 'react';
import PropTypes from 'prop-types';
import { withStyles } from 'material-ui/styles';
import MaterialDrawer from 'material-ui/Drawer';
import Button from 'material-ui/Button';
import List, { ListItem, ListItemIcon, ListItemText } from 'material-ui/List';
import AddIcon from 'material-ui-icons/Add';
import PersonIcon from 'material-ui-icons/Person';
import PeopleIcon from 'material-ui-icons/People';
import StatisticsIcon from 'material-ui-icons/ShowChart';
import FitnessCenterIcon from 'material-ui-icons/FitnessCenter';
import Divider from 'material-ui/Divider';
import Link from 'next/link';

const styles = theme => ({
  list: {
    width: 250,
    background: theme.palette.background.paper,
  },
  listFull: {
    width: 'auto',
  },
});

function Drawer({ open, onRequestClose, classes }) {
  const sideList = (
    <div className={classes.list}>
      <List>
        <Link prefetch href="/newMatch">
          <ListItem button>
            <ListItemIcon>
              <AddIcon />
            </ListItemIcon>
            <ListItemText primary="New Match" />
          </ListItem>
        </Link>
        <Divider />
        <Link prefetch href="/">
          <ListItem button>
            <ListItemIcon>
              <FitnessCenterIcon />
            </ListItemIcon>
            <ListItemText primary="Matches" />
          </ListItem>
        </Link>
        <Link prefetch href="/statistic">
          <ListItem button>
            <ListItemIcon>
              <StatisticsIcon />
            </ListItemIcon>
            <ListItemText primary="Statistics" />
          </ListItem>
        </Link>
      </List>
    </div>
  );

  return (
    <MaterialDrawer open={open} onRequestClose={onRequestClose}>
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

Drawer.propTypes = {
  classes: PropTypes.object.isRequired,
};

export default withStyles(styles)(Drawer);
