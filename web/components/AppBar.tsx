import React from 'react';
import { withStyles } from '@material-ui/core/styles';
import AppBar from '@material-ui/core/AppBar';
import Toolbar from '@material-ui/core/Toolbar';
import Typography from '@material-ui/core/Typography';
import Button from '@material-ui/core/Button';
import IconButton from '@material-ui/core/IconButton';
import MenuIcon from '@material-ui/icons/Menu';
import { Classes } from '../types';

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

interface Props {
  onOpenMenu: () => void;
  title: string;
  isLoggedIn: boolean;
  onLogout: () => void;
  classes: Classes;
}

function ButtonAppBar({
  onOpenMenu,
  title,
  isLoggedIn,
  onLogout,
  classes,
}: Props) {
  const button = isLoggedIn ? (
    <Button color="inherit" onClick={onLogout}>
      Logout
    </Button>
  ) : null;

  return (
    <div className={classes.root}>
      <AppBar>
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
