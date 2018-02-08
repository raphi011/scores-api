// @flow

import React from 'react';
import { withStyles } from 'material-ui/styles';
import AppBar from 'material-ui/AppBar';
import Toolbar from 'material-ui/Toolbar';
import Typography from 'material-ui/Typography';
import Button from 'material-ui/Button';
import IconButton from 'material-ui/IconButton';
import MenuIcon from 'material-ui-icons/Menu';
import Tooltip from 'material-ui/Tooltip';
import Avatar from 'material-ui/Avatar';
import type { User, Classes } from '../types';

const styles = () => ({
  row: {
    display: 'flex',
    justifyContent: 'center',
  },
  root: {
    width: '100%',
  },
  flex: {
    flex: 1,
  },
  menuButton: {
    marginLeft: -12,
    marginRight: 20,
  },
});

type Props = {
  onOpenMenu: () => void,
  title: string,
  isLoggedIn: boolean,
  user: User,
  onLogout: () => void,
  classes: Classes,
};

function ButtonAppBar({
  onOpenMenu,
  title,
  isLoggedIn,
  user,
  onLogout,
  classes,
}: Props) {
  const button = isLoggedIn ? (
    <Tooltip title={user.email} placement="bottom">
      <div className={classes.row}>
        <Avatar src={user.profileImageUrl} />
        <Button color="inherit" onClick={onLogout}>
          Logout
        </Button>
      </div>
    </Tooltip>
  ) : null;

  return (
    <div className={classes.root}>
      <AppBar position="fixed">
        <Toolbar>
          <IconButton
            color="inherit"
            onClick={onOpenMenu}
            className={classes.menuButton}
            aria-label="Menu"
          >
            <MenuIcon />
          </IconButton>
          <Typography variant="title" color="inherit" className={classes.flex}>
            {title}
          </Typography>
          {button}
        </Toolbar>
      </AppBar>
    </div>
  );
}

export default withStyles(styles)(ButtonAppBar);
