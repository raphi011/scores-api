import React from "react";
import Router from "next/router";
import withRedux from "next-redux-wrapper";

import { loggedInAction, setStatusAction } from "../redux/actions/action";
import initStore from "../redux/store";

class LoggedIn extends React.Component {
  static async getInitialProps(props) {
    const { store, query, isServer } = props;

    const { username, error } = query;

    if (!error) {
      await store.dispatch(loggedInAction(username));
    } else {
      await store.dispatch(setStatusAction("User not found"));
    }
  }

  componentDidMount() {
    Router.replace("/");
  }

  render() {
    return null;
  }
}

export default withRedux(initStore)(LoggedIn);
