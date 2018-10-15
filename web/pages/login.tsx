import React, { SyntheticEvent } from 'react';
import Button from '@material-ui/core/Button';
import Typography from '@material-ui/core/Typography';
import TextField from '@material-ui/core/TextField';
import Paper from '@material-ui/core/Paper';
import Router from 'next/router';
import Red from '@material-ui/core/colors/red';
import WarningIcon from '@material-ui/icons/Warning';
import FormGroup from '@material-ui/core/FormGroup';

import withAuth from '../containers/AuthContainer';
import { loginWithPasswordAction } from '../redux/actions/auth';
import { setStatusAction } from '../redux/actions/status';
import Snackbar from '../containers/SnackbarContainer';
import LoadingButton from '../components/LoadingButton';

interface Props {
  r: string;
  loginRoute: string;
  fromServer: boolean;
  error: string;
  classes: any;
  loginWithPassword: (
    credentials: { email: string; password: string },
  ) => Promise<any>;
  setStatus: (string) => void;
}

interface State {
  email: string;
  password: string;
  loggingIn: boolean;
}

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
