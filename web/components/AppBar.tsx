import AppBar from '@material-ui/core/AppBar';
import Button from '@material-ui/core/Button';
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
import Link from 'next/link';
import React from 'react';

const styles = (theme: Theme) =>
  createStyles({
    appBar: {
      zIndex: theme.zIndex.drawer + 1,
    },
    flex: {
      cursor: 'pointer',
      flex: 1,
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
  onOpenMenu: () => void;
  title: { text: string; href: string };
  isLoggedIn: boolean;
  onLogout: () => Promise<void>;
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
    <AppBar position="fixed" className={classes.appBar}>
      <Toolbar>
        <Hidden mdUp>
          <IconButton
            color="inherit"
            onClick={onOpenMenu}
            className={classes.menuButton}
            aria-label="Menu"
          >
            <MenuIcon />
          </IconButton>
        </Hidden>
        <Link href={title.href}>
          <Typography variant="h6" color="inherit" className={classes.flex}>
            {title.text}
          </Typography>
        </Link>
        {button}
      </Toolbar>
    </AppBar>
  );
}

export default withStyles(styles)(ButtonAppBar);
