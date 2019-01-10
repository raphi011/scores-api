import AppBar from '@material-ui/core/AppBar';
import Button from '@material-ui/core/Button';
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
import Link from 'next/link';
import React from 'react';

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

  onOpenMenu: () => void;
  onLogout: () => Promise<void>;
}

function ButtonAppBar({
  onOpenMenu,
  title,
  isLoggedIn,
  onLogout,
  bodyScrolled,
  classes,
}: Props) {
  const button = isLoggedIn ? (
    <Button color="inherit" onClick={onLogout}>
      Logout
    </Button>
  ) : null;

  let className = classes.appBar;

  if (!bodyScrolled) {
    className += ` ${classes.bodyNotScrolled}`;
  }

  return (
    <AppBar position="fixed" className={className}>
      <Toolbar>
        <IconButton
          color="inherit"
          onClick={onOpenMenu}
          className={classes.menuButton}
          aria-label="Menu"
        >
          <MenuIcon />
        </IconButton>
        <Link href={title.href}>
          <Typography variant="h6" className={classes.logo} color="inherit">
            {title.text}
          </Typography>
        </Link>
        <div className={classes.flex} />

        {button}
      </Toolbar>
    </AppBar>
  );
}

export default withStyles(styles)(ButtonAppBar);
