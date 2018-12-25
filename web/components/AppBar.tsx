import AppBar from '@material-ui/core/AppBar';
import Button from '@material-ui/core/Button';
import Hidden from '@material-ui/core/Hidden';
import IconButton from '@material-ui/core/IconButton';
import { createStyles, WithStyles, withStyles } from '@material-ui/core/styles';
import Toolbar from '@material-ui/core/Toolbar';
import Typography from '@material-ui/core/Typography';
import MenuIcon from '@material-ui/icons/Menu';
import Link from 'next/link';
import React from 'react';

const styles = createStyles({
  flex: {
    cursor: 'pointer',
    flex: 1,
  },
  menuButton: {
    marginLeft: -12,
    marginRight: 20,
  },
  root: {
    width: '100%',
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
            <Typography variant="h6" color="inherit" className={classes.flex}>
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
