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
  <div>
    <Typography variant="headline">Please login</Typography>
    <Button color="secondary" href={loginRoute}>
      Login with Google
    </Button>
  </div>
);

export default withAuth(Login);
