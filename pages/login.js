// @flow

import React from 'react';
import Router from 'next/router';
import Button from 'material-ui/Button';
import Typography from 'material-ui/Typography';

import withAuth from '../containers/AuthContainer';

type Props = {
  isLoggedIn: boolean,
  loginRoute: ?string,
  redirect: ?string,
};

class Login extends React.Component<Props> {
  // componentDidMount() {
  //   const { isLoggedIn, redirect } = this.props;

  //   if (isLoggedIn) {
  //     Router.replace(redirect || '/');
  //   }
  // }

  render() {
    const { loginRoute } = this.props;

    return (
      <div>
        <Typography variant="headline">Please login</Typography>
        <Button color="secondary" href={loginRoute}>
          Login with Google
        </Button>
      </div>
    );
  }
}

export default withAuth(Login);
