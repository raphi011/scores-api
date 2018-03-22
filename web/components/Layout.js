// @flow

import React from 'react';
import { withStyles } from 'material-ui/styles';

import Drawer from './Drawer';
import AppBar from '../containers/AppBarContainer';
import Snackbar from '../containers/SnackbarContainer';

import type { Group } from '../types';

type Props = {
  title: string,
  children: React.Node,
  groups: Array<Group>,
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
    const { title, groups = [], children, classes } = this.props;
    return (
      <div className={classes.style}>
        <Drawer
          onRequestClose={this.onCloseDrawer}
          groups={groups}
          open={this.state.open}
        />
        <AppBar onOpenMenu={this.onToggleDrawer} title={title} />
        {children}
        <Snackbar />
      </div>
    );
  }
}

export default withStyles(styles)(Layout);
