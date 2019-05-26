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
import classNames from 'classnames';
import Modal from '@material-ui/core/Modal';

const styles = (theme: Theme) =>
  createStyles({
    active: {
      color: theme.palette.primary.main,
    },
    avatar: {
      padding: 0,
    },
    inactive: {
      color: theme.palette.grey[500],
    },
    appBar: {
      borderBottom: `1px solid ${theme.palette.grey[100]}`,
    },
    mobile: {
      display: 'flex',
      flexDirection: 'column',
      alignItems: 'stretch',
      justifyContent: 'flexStart',
    },
    removeDropShadow: {
      borderBottom: `0px`,
      boxShadow: 'none',
    },
    drawer: {
      margin: `0 ${theme.spacing(2)}px`,
    },
    flex: {
      flex: 1,
    },
    links: {
      alignItems: 'baseline',
      display: 'flex',
      flexDirection: 'row',
    },
    link: {
      fontWeight: 400,
      cursor: 'pointer',
      marginRight: '20px',
    },
    logo: {
      color: theme.palette.grey[900],
      cursor: 'pointer',
      display: 'inline-block',
      fontSize: '20px',
      fontWeight: 500,
    },
    menuButton: {
      color: theme.palette.primary.light,
      marginRight: 20,
    },
    row: {
      display: 'flex',
      justifyContent: 'center',
    },
    toolbar: {
      margin: 'auto',
      maxWidth: maxContentWidth,
      height: '80px',
      minHeight: '80px',
      width: '100%',
    },
  });

interface Props extends WithStyles<typeof styles> {
  bodyScrolled: boolean;
  title: { text: string; href: string };
  isLoggedIn: boolean;
  drawerOpen: boolean;
  isMobile: boolean;
  user: User | null;

  onCloseDrawer: (e: React.SyntheticEvent) => void;
  onToggleDrawer: () => void;
}

const links = [
  {
    href: '/',
    name: 'Tournaments',
    altHref: '/tournament',
  },
  {
    href: '/ladder',
    name: 'Ladder',
  },
  {
    href: '/admin',
    name: 'Admin',
    adminOnly: true,
  },
];

function ButtonAppBar({
  onToggleDrawer,
  onCloseDrawer,
  isMobile,
  user,
  drawerOpen,
  bodyScrolled,
  classes,
}: Props) {
  const className = classNames(classes.appBar, {
    [classes.removeDropShadow]: !bodyScrolled && !drawerOpen,
    [classes.mobile]: isMobile,
    // [classes.desktop]: !isMobile,
    // [classes.drawerOpen]: drawerOpen,
  });

  const renderedLinks = links.map(l => (
    <span key={l.href} className={classes.link}>
      <RenderLink {...l} classes={classes} />
    </span>
  ));

  let navbarLinks = null;
  let drawerLinks = null;

  if (isMobile) {
    drawerLinks = renderedLinks;
  } else {
    navbarLinks = renderedLinks;
  }

  const userId = user ? user.id : 0;
  const profileImageUrl = user ? user.profileImageUrl : '';

  const appBar = (
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
          <Typography component="a" className={classes.logo} color="inherit">
            Scores
          </Typography>
        </Link>
        <div className={classes.flex} />
        {navbarLinks ? (
          <span className={classes.links}>{navbarLinks}</span>
        ) : null}
        <Link href={{ pathname: '/user', query: { id: userId } }}>
          <IconButton className={classes.avatar}>
            <Avatar
              alt="Your profile picture"
              src={profileImageUrl ? `${profileImageUrl}?sz=40` : ''}
            />
          </IconButton>
        </Link>
      </Toolbar>
      {drawerLinks && drawerOpen ? (
        <div className={classes.drawer}>{drawerLinks}</div>
      ) : null}
    </AppBar>
  );

  if (drawerOpen) {
    return (
      <Modal onClose={onCloseDrawer} open={drawerOpen}>
        {appBar}
      </Modal>
    );
  }

  return appBar;
}

interface RenderLinkProps extends WithStyles<typeof styles> {
  href: string;
  adminOnly?: boolean;
  altHref?: string;
  name: string;
}

function RenderLink({
  href,
  altHref,
  adminOnly,
  name,
  classes,
}: RenderLinkProps) {
  const link = (
    <ActiveLink
      altHref={altHref}
      activeClassName={classes.active}
      prefetch
      href={href}
    >
      <Typography variant="subtitle1">{name}</Typography>
    </ActiveLink>
  );

  if (adminOnly) {
    return <AdminOnly>{link}</AdminOnly>;
  }

  return link;
}

export default withStyles(styles)(ButtonAppBar);
