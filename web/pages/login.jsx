// @flow

import React from 'react';
import { withStyles } from 'material-ui/styles';
import Button from 'material-ui/Button';
import Typography from 'material-ui/Typography';
import TextField from 'material-ui/TextField';
import Paper from 'material-ui/Paper';
import Router from 'next/router';
import { CircularProgress } from 'material-ui/Progress';
import Red from 'material-ui/colors/red';
import WarningIcon from 'material-ui-icons/Warning';
import { FormGroup } from 'material-ui/Form';

import withAuth from '../containers/AuthContainer';
import { loginWithPasswordAction } from '../redux/actions/auth';
import { setStatusAction } from '../redux/actions/status';
import Snackbar from '../containers/SnackbarContainer';

import type { Classes } from '../types';

type Props = {
  r: string,
  loginRoute: string,
  fromServer: boolean,
  error: string,
  classes: Classes,
  loginWithPassword: ({ email: string, password: string }) => Promise<any>,
  setStatus: string => void,
};

type State = {
  email: string,
  password: string,
  loggingIn: boolean,
};

const styles = theme => ({
  wrapper: {
    margin: theme.spacing.unit,
    position: 'relative',
  },
  buttonProgress: {
    position: 'absolute',
    top: '50%',
    left: '50%',
    marginTop: -12,
    marginLeft: -12,
  },
});

class Login extends React.Component<Props, State> {
  static getParameters(query) {
    const { error, r } = query;

    return { error, r };
  }

  static mapDispatchToProps = {
    loginWithPassword: loginWithPasswordAction,
    setStatus: setStatusAction,
  };

  state = {
    email: '',
    password: '',
    loggingIn: false,
  };

  componentDidMount() {
    // TODO: load new loginRoute after logout
    const { fromServer } = this.props;

    if (!fromServer) {
      Router.prefetch('/');
    }
  }

  onEmailChange = e => {
    this.setState({ email: e.target.value });
  };

  onPasswordChange = e => {
    this.setState({ password: e.target.value });
  };

  loginWithPassword = async () => {
    const { setStatus, loginWithPassword, r } = this.props;
    const { email, password } = this.state;

    this.setState({ loggingIn: true });

    const credentials = {
      email,
      password,
    };

    try {
      await loginWithPassword(credentials);
      const path = r || '/';
      await Router.push(path);
    } catch (e) {
      setStatus('Something went wrong there');
      this.setState({ loggingIn: false });
    }
  };

  render() {
    const { loginRoute, error, classes } = this.props;
    const { email, password, loggingIn } = this.state;

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
          <FormGroup>
            <TextField
              label="Email"
              type="email"
              value={email}
              onChange={this.onEmailChange}
              autoComplete="email"
              margin="normal"
            />
            <TextField
              label="Password"
              type="password"
              value={password}
              onChange={this.onPasswordChange}
              autoComplete="current-password"
              margin="normal"
            />
          </FormGroup>
          <div className={classes.wrapper}>
            <Button
              color="primary"
              fullWidth
              variant="raised"
              disabled={loggingIn}
              onClick={this.loginWithPassword}
            >
              Login
            </Button>
            {loggingIn && (
              <CircularProgress size={24} className={classes.buttonProgress} />
            )}
          </div>
          <div style={{ margin: '20px 0' }}>- or -</div>
          <Button
            color="primary"
            disabled={!loginRoute}
            fullWidth
            variant="raised"
            href={loginRoute}
          >
            Login with Google
          </Button>
          <div
            style={{ color: Red['800'], height: '50px', paddingTop: '20px' }}
          >
            {errorBox}
          </div>
        </Paper>
        <Snackbar />
      </div>
    );
  }
}

export default withStyles(styles)(withAuth(Login));
