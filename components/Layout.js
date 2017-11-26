import React from "react";
import Head from "next/head";
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
  }

  render() {
    const { status, title, children } = this.props;
    return (
      <div>
        <Head>
          <meta name="viewport" content="width=device-width, initial-scale=1" />
          <meta charSet="utf-8" />
          <style>
            {`body {
              margin: 0;
            }`}
            </style>
        </Head>
        <Drawer onRequestClose={this.onCloseDrawer} open={this.state.open} />
        <AppBar onOpenMenu={this.onToggleDrawer} title={title} />
        {children}
        <Snackbar status={status} />
      </div>
    );
  }
}

export default Layout;
