import React from 'react';
import { withStyles, createStyles } from '@material-ui/core/styles';
import AppBar from '@material-ui/core/AppBar';
import Toolbar from '@material-ui/core/Toolbar';
import Hidden from '@material-ui/core/Hidden';
import Typography from '@material-ui/core/Typography';
import Button from '@material-ui/core/Button';
import IconButton from '@material-ui/core/IconButton';
import MenuIcon from '@material-ui/icons/Menu';
import Link from 'next/link';

const styles = createStyles({
  row: {
    display: 'flex',
    justifyContent: 'center',
  },
  root: {
    width: '100%',
  },
  flex: {
    flex: 1,
    cursor: 'pointer',
  },
  menuButton: {
    marginLeft: -12,
    marginRight: 20,
  },
});

interface Props {
  onOpenMenu: () => void;
  title: { text: string; href: string };
  isLoggedIn: boolean;
  onLogout: () => Promise<void>;
  classes: any;
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
      <AppBar position="static">
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
            <Typography
              variant="h6"
              color="inherit"
              className={classes.flex}
            >
              {title.text}
            </Typography>
          </Link>
          {button}
        </Toolbar>
      </AppBar>
    </div>
  );
}

export default withStyles(styles)(ButtonAppBar);
