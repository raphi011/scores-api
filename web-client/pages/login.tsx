import React from 'react';

import Router from 'next/router';

import { createStyles, withStyles, WithStyles } from '@material-ui/core';
import Button from '@material-ui/core/Button';
import Red from '@material-ui/core/colors/red';
import FormGroup from '@material-ui/core/FormGroup';
import Paper from '@material-ui/core/Paper';
import TextField from '@material-ui/core/TextField';
import Typography from '@material-ui/core/Typography';
import WarningIcon from '@material-ui/icons/Warning';

import LoadingButton from '../components/LoadingButton';
import * as Query from '../utils/query';
import Snackbar from '../containers/SnackbarContainer';
import {
  loginWithPasswordAction,
  userOrLoginRouteAction,
} from '../redux/auth/actions';
import { loginRouteSelector, userSelector } from '../redux/auth/selectors';
import { setStatusAction } from '../redux/status/actions';
import { Store } from '../redux/store';
import withConnect, { Context } from '../hoc/next/withConnect';
import { User } from '../types';

interface Props extends WithStyles<typeof styles> {
  r: string;
  loginRoute: string;
  loginError: string;
  user: User | null;

  loginWithPassword: (credentials: {
    email: string;
    password: string;
  }) => Promise<void>;
  setStatus: (status: string) => void;
}

interface State {
  email: string;
  password: string;
  loggingIn: boolean;
}

const styles = createStyles({
  center: {
    alignItems: 'center',
    display: 'flex',
    height: '100%',
    justifyContent: 'center',
  },
  container: {
    marginTop: '30px',
    padding: '30px',
    textAlign: 'center',
    width: '350px',
  },
});

class Login extends React.Component<Props, State> {
  static mapDispatchToProps = {
    loginWithPassword: loginWithPasswordAction,
    setStatus: setStatusAction,
  };

  static async getInitialProps(ctx: Context): Promise<Partial<Props>> {
    const { query } = ctx;

    const loginError = Query.one(query, 'error');
    const r = Query.one(query, 'r');

    return { loginError, r };
  }

  static mapStateToProps(state: Store) {
    const loginRoute = loginRouteSelector(state);
    const user = userSelector(state);

    return { loginRoute, user };
  }

  state = {
    email: '',
    loggingIn: false,
    password: '',
  };

  static buildActions() {
    return [userOrLoginRouteAction()];
  }

  async componentDidMount() {
    const { user } = this.props;

    if (user) {
      await this.onLoggedIn();
    } else {
      Router.prefetch('/');
    }
  }

  onEmailChange = (e: React.ChangeEvent<HTMLInputElement>) => {
    this.setState({ email: e.target.value });
  };

  onPasswordChange = (e: React.ChangeEvent<HTMLInputElement>) => {
    this.setState({ password: e.target.value });
  };

  onLoggedIn = async () => {
    const { r } = this.props;

    const path = r || '/';
    await Router.push(path);
  };

  loginWithPassword = async (e: React.FormEvent<HTMLFormElement>) => {
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
      await this.onLoggedIn();
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
    const { loginRoute, loginError, classes } = this.props;
    const { email, password, loggingIn } = this.state;

    const errorBox = loginError ? (
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
        {loginError}
      </span>
    ) : null;

    return (
      <div className={classes.center}>
        <Paper className={classes.container}>
          <form onSubmit={this.loginWithPassword}>
            <Typography variant="h3">Scores</Typography>
            <br />
            <FormGroup>
              <TextField
                label="Email"
                type="email"
                value={email}
                onChange={this.onEmailChange}
                autoComplete="email"
                margin="normal"
                autoFocus
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
            <LoadingButton loading={loggingIn}>
              <span>Login</span>
            </LoadingButton>
            <div style={{ margin: '20px 0' }}>- or -</div>
            <Button
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

export default withConnect(withStyles(styles)(Login));
