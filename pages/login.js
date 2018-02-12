// @flow

import React from 'react';
import Button from 'material-ui/Button';
import Typography from 'material-ui/Typography';

import withAuth from '../containers/AuthContainer';

type Props = {
  loginRoute: string,
};

// TODO: load new loginRoute after logout
// Login.getInitialProps = async ({ }) {

// }

const Login = ({ loginRoute }: Props) => (
  <div style={{ display: "flex", justifyContent: "center", height: "100%", alignItems: "center" }}>
    <div style={{ textAlign: "center" }}>
    <Typography variant="display2">Welcome</Typography>
    <br />
    <Button color="primary" variant="raised" href={loginRoute}>
      Login with Google
    </Button>
    </div>
  </div>
);

export default withAuth(Login);
