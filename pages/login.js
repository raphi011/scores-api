// @flow 

import React from "react";
import Router from "next/router";
import withRedux from "next-redux-wrapper";

import Button from 'material-ui/Button';
import { userOrLoginRouteAction } from "../redux/actions/action";
import initStore, { dispatchActions } from "../redux/store";
import { userSelector, loginRouteSelector } from "../redux/reducers/reducer";

type Props = {
  isLoggedIn: boolean,
  loginRoute: ?string,
  redirect: ?string,
};

class Login extends React.Component<Props> {
  static async getInitialProps({ store, query, isServer, req, res }) {
    const actions = [userOrLoginRouteAction()];

    const { redirect } = query;

    await dispatchActions(store.dispatch, isServer, req, res, actions);

    return { redirect };
  }

  componentDidMount() {
    const { isLoggedIn, redirect } = this.props;
 
    if (isLoggedIn) {
      Router.replace(redirect || "/");
    }
  }

  render() {
    const { loginRoute } = this.props;

    return (
      <div>
        <Button color="contrast" href={loginRoute}>
          Login
        </Button>
      </div>
    );
  }
}

function mapStateToProps(state) {
  const { isLoggedIn } = userSelector(state);
  const loginRoute = loginRouteSelector(state);

  return {
    isLoggedIn,
    loginRoute,
  };
}

export default withRedux(initStore, mapStateToProps)(Login);
