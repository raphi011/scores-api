import React from 'react';

import Link from 'next/link';

import AppBar from '@material-ui/core/AppBar';
import Avatar from '@material-ui/core/Avatar';
import Hidden from '@material-ui/core/Hidden';
import IconButton from '@material-ui/core/IconButton';
import {
  createStyles,
  Theme,
  WithStyles,
  withStyles,
} from '@material-ui/core/styles';
import Toolbar from '@material-ui/core/Toolbar';
import Typography from '@material-ui/core/Typography';
import MenuIcon from '@material-ui/icons/Menu';

import { User } from '../types';

const styles = (theme: Theme) =>
  createStyles({
    appBar: {
      backgroundColor: 'white',
      borderBottom: `1px solid ${theme.palette.grey[100]}`,
      zIndex: theme.zIndex.drawer + 1,
    },
    bodyNotScrolled: {
      boxShadow: 'none',
    },
    flex: {
      flex: 1,
    },
    logo: {
      cursor: 'pointer',
    },
    menuButton: {
      marginLeft: -12,
      marginRight: 20,
    },
    row: {
      display: 'flex',
      justifyContent: 'center',
    },
  });

interface Props extends WithStyles<typeof styles> {
  bodyScrolled: boolean;
  title: { text: string; href: string };
  isLoggedIn: boolean;
  user: User;

  onToggleDrawer: () => void;
}

function ButtonAppBar({
  onToggleDrawer,
  title,
  user,
  bodyScrolled,
  classes,
}: Props) {
  let className = classes.appBar;

  if (!bodyScrolled) {
    className += ` ${classes.bodyNotScrolled}`;
  }

  return (
    <AppBar position="fixed" className={className}>
      <Toolbar>
        <Hidden mdUp>
          <IconButton
            color="inherit"
            onClick={onToggleDrawer}
            className={classes.menuButton}
            aria-label="Menu"
          >
            <MenuIcon />
          </IconButton>
        </Hidden>
        <Link href={title.href}>
          <Typography variant="h6" className={classes.logo} color="inherit">
            {title.text}
          </Typography>
        </Link>
        <div className={classes.flex} />
        <Link href={{ pathname: '/user', query: { id: user.id } }}>
          <IconButton>
            <Avatar src={user.profileImageUrl} />
          </IconButton>
        </Link>
      </Toolbar>
    </AppBar>
  );
}

export default withStyles(styles)(ButtonAppBar);
