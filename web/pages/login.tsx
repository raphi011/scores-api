import React, { SyntheticEvent } from 'react';

import Router from 'next/router';

import Button from '@material-ui/core/Button';
import Red from '@material-ui/core/colors/red';
import FormGroup from '@material-ui/core/FormGroup';
import Paper from '@material-ui/core/Paper';
import TextField from '@material-ui/core/TextField';
import Typography from '@material-ui/core/Typography';
import WarningIcon from '@material-ui/icons/Warning';

import LoadingButton from '../components/LoadingButton';
import withAuth from '../containers/AuthContainer';
import Snackbar from '../containers/SnackbarContainer';
import { loginWithPasswordAction } from '../redux/auth/actions';
import { setStatusAction } from '../redux/status/actions';

interface Props {
  r: string;
  loginRoute: string;
  fromServer: boolean;
  error: string;

  loginWithPassword: (
    credentials: { email: string; password: string },
  ) => Promise<any>;
  setStatus: (status: string) => void;
}

interface State {
  email: string;
  password: string;
  loggingIn: boolean;
}

class Login extends React.Component<Props, State> {
  static mapDispatchToProps = {
    loginWithPassword: loginWithPasswordAction,
    setStatus: setStatusAction,
  };

  static getParameters(query) {
    const { error, r } = query;

    return { error, r };
  }

  state = {
    email: '',
    loggingIn: false,
    password: '',
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

  loginWithPassword = async (e: SyntheticEvent<HTMLFormElement>) => {
    e.preventDefault();

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
    } catch ({ responseCode }) {
      this.setState({ loggingIn: false });

      if (responseCode === 401) {
        setStatus('Wrong username or password.');
      } else {
        setStatus('Something went wrong there.');
      }
    }
  };

  render() {
    const { loginRoute, error } = this.props;
    const { email, password, loggingIn } = this.state;

    const errorBox = error ? (
      <span
        style={{
          alignItems: 'center',
          display: 'flex',
          flexDirection: 'row',
          fontWeight: 'bold',
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
          alignItems: 'center',
          display: 'flex',
          height: '100%',
          justifyContent: 'center',
        }}
      >
        <Paper style={{ textAlign: 'center', padding: '30px' }}>
          <form onSubmit={this.loginWithPassword}>
            <Typography variant="h3">Welcome</Typography>
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
            <LoadingButton loading={loggingIn}>Login</LoadingButton>
            <div style={{ margin: '20px 0' }}>- or -</div>
            <Button
              color="primary"
              disabled={!loginRoute}
              fullWidth
              variant="contained"
              href={loginRoute}
            >
              Login with Google
            </Button>
            <div
              style={{ color: Red['800'], height: '50px', paddingTop: '20px' }}
            >
              {errorBox}
            </div>
          </form>
        </Paper>
        <Snackbar />
      </div>
    );
  }
}

export default withAuth(Login);
