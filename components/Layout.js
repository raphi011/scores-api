import React from "react";
import Drawer from "./Drawer";
import AppBar from "./AppBar";
import Snackbar from "./Snackbar";

class Layout extends React.Component {
  state = {
    open: false
  };

  onToggleDrawer = () => {
    this.setState({ open: !this.state.open });
  };

  onCloseDrawer = () => {
    this.setState({ open: false });
  };

  render() {
    const { status, title, loginRoute, children } = this.props;
    return (
      <div>
        <Drawer onRequestClose={this.onCloseDrawer} open={this.state.open} />
        <AppBar
          onOpenMenu={this.onToggleDrawer}
          loginRoute={loginRoute}
          title={title}
        />
        {children}
        <Snackbar status={status} />
      </div>
    );
  }
}

export default Layout;
