// @flow

import React from 'react';
import { withStyles } from 'material-ui/styles';

import Drawer from './Drawer';
import AppBar from '../containers/AppBarContainer';
import Snackbar from '../containers/SnackbarContainer';

type Props = {
  title: string,
  children: React.Node,
  classes: Object,
};

const styles = theme => ({
  style: {
    backgroundColor: theme.palette.background.paper,
    marginTop: '56px',
  },
});

type State = {
  open: boolean,
};

class Layout extends React.Component<Props, State> {
  state = {
    open: false,
  };

  onToggleDrawer = () => {
    this.setState({ open: !this.state.open });
  };

  onCloseDrawer = () => {
    this.setState({ open: false });
  };

  render() {
    const { title, children, classes } = this.props;
    return (
      <div className={classes.style}>
        <Drawer onRequestClose={this.onCloseDrawer} open={this.state.open} />
        <AppBar onOpenMenu={this.onToggleDrawer} title={title} />
        {children}
        <Snackbar />
      </div>
    );
  }
}

export default withStyles(styles)(Layout);
