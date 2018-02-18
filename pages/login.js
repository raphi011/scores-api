// @flow

import React from 'react';
import Button from 'material-ui/Button';
import Typography from 'material-ui/Typography';
import Paper from 'material-ui/Paper';
import Router from 'next/router';
import Red from 'material-ui/colors/red';
import WarningIcon from 'material-ui-icons/Warning';

import withAuth from '../containers/AuthContainer';

type Props = {
  loginRoute: string,
  fromServer: boolean,
  error: string,
};

class Login extends React.Component<Props> {
  static getParameters(query) {
    const { error } = query;

    return { error };
  }

  componentDidMount() {
    // TODO: load new loginRoute after logout
    const { fromServer } = this.props;

    if (!fromServer) {
      Router.prefetch('/');
    }
  }

  render() {
    const { loginRoute, error } = this.props;

    const errorBox = error ? (
      <span
        style={{
          display: 'flex',
          fontWeight: 'bold',
          flexDirection: 'row',
          alignItems: 'center',
          justifyContent: 'center',
        }}
      >
        <WarningIcon />
        {error}
      </span>
    ) : null;

    return (
      <div
        style={{
          display: 'flex',
          justifyContent: 'center',
          height: '100%',
          alignItems: 'center',
        }}
      >
        <Paper style={{ textAlign: 'center', padding: '30px' }}>
          <Typography variant="display2">Welcome</Typography>
          <br />
          <Button color="primary" variant="raised" href={loginRoute}>
            Login with Google
          </Button>
          <div
            style={{ color: Red['800'], height: '50px', paddingTop: '20px' }}
          >
            {errorBox}
          </div>
        </Paper>
      </div>
    );
  }
}

export default withAuth(Login);
