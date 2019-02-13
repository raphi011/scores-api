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

import AdminOnly from '../containers/AdminOnly';
import { maxContentWidth } from '../styles/theme';
import { User } from '../types';
import ActiveLink from './ActiveLink';

const styles = (theme: Theme) =>
  createStyles({
    active: {
      color: theme.palette.primary[600],
    },
    appBar: {
      borderBottom: `1px solid ${theme.palette.grey[100]}`,
      height: '80px',
    },
    bodyNotScrolled: {
      boxShadow: 'none',
    },
    flex: {
      flex: 1,
    },
    links: {
      alignItems: 'baseline',
      display: 'flex',
      flexDirection: 'row',

      '&> *': {
        color: theme.palette.grey[500],
        fontWeight: 400,
        cursor: 'pointer',
        marginRight: '20px',
      },
    },
    logo: {
      color: theme.palette.grey[900],
      cursor: 'pointer',
      display: 'inline-block',
      fontSize: '20px',
      fontWeight: 500,
    },
    menuButton: {
      color: theme.palette.primary[200],
      marginRight: 20,
    },
    page: {
      color: theme.palette.primary[100],
      display: 'inline-block',
      marginLeft: '7px',
    },
    row: {
      display: 'flex',
      justifyContent: 'center',
    },
    toolbar: {
      margin: 'auto',
      maxWidth: maxContentWidth,
      width: '100%',
    },
  });

interface Props extends WithStyles<typeof styles> {
  bodyScrolled: boolean;
  title: { text: string; href: string };
  isLoggedIn: boolean;
  user: User;

  onToggleDrawer: () => void;
}

function ButtonAppBar({ onToggleDrawer, user, bodyScrolled, classes }: Props) {
  let className = classes.appBar;

  if (!bodyScrolled) {
    className += ` ${classes.bodyNotScrolled}`;
  }

  return (
    <AppBar color="default" position="fixed" className={className}>
      <Toolbar className={classes.toolbar}>
        <Hidden mdUp>
          <IconButton
            onClick={onToggleDrawer}
            className={classes.menuButton}
            aria-label="Menu"
          >
            <MenuIcon />
          </IconButton>
        </Hidden>
        <Link href={'/'}>
          <Typography className={classes.logo} color="inherit">
            Scores
          </Typography>
        </Link>
        <div className={classes.flex} />
        <span className={classes.links}>
          <ActiveLink activeClassName={classes.active} prefetch href="/">
            <Typography variant="subtitle1">Tournaments</Typography>
          </ActiveLink>
          <ActiveLink activeClassName={classes.active} prefetch href="/ladder">
            <Typography variant="subtitle1">Ladder</Typography>
          </ActiveLink>
          <AdminOnly>
            <ActiveLink activeClassName={classes.active} href="/admin">
              <Typography variant="subtitle1">Admin</Typography>
            </ActiveLink>
          </AdminOnly>
        </span>
        <Link href={{ pathname: '/user', query: { id: user.id } }}>
          <IconButton style={{ padding: 0 }}>
            <Avatar alt="Your profile picture" src={user.profileImageUrl} />
          </IconButton>
        </Link>
      </Toolbar>
    </AppBar>
  );
}

export default withStyles(styles)(ButtonAppBar);
