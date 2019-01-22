import React from 'react';

import FormControlLabel from '@material-ui/core/FormControlLabel';
import { createStyles, WithStyles, withStyles } from '@material-ui/core/styles';
import Switch from '@material-ui/core/Switch';
import TextField from '@material-ui/core/TextField';

import DoneIcon from '@material-ui/icons/Done';
import LoadingButton from '../LoadingButton';

const styles = createStyles({});

interface Props extends WithStyles<typeof styles> {
  username?: string;

  onLogin: (username: string, password: string, rememberMe: boolean) => void;
}

interface State {
  loggingIn: boolean;
  username: string;
  password: string;
  rememberMe: boolean;
  usernameValidation: string;
}

class Login extends React.Component<Props, State> {
  loginRegex = /^[A-Z][a-z]+\.[A-Z]+$/;
  constructor(props) {
    super(props);

    this.state = {
      loggingIn: false,
      password: '',
      rememberMe: true,
      username: props.username || '',
      usernameValidation: '',
    };
  }

  onChangeUsername = event => {
    const username = event.target.value;
    let usernameValidation = '';

    if (username && !this.loginRegex.test(username)) {
      usernameValidation = 'Incorrect username';
    }

    this.setState({ username, usernameValidation });
  };

  onChangePassword = event => {
    const password = event.target.value;

    this.setState({ password });
  };

  onLogin = (e: React.FormEvent<HTMLFormElement>) => {
    e.preventDefault();

    const { onLogin } = this.props;
    const { username, password, rememberMe } = this.state;

    if (username && this.loginRegex.test(username) && password) {
      try {
        this.setState({ loggingIn: true });
        onLogin(username, password, rememberMe);
      } finally {
        this.setState({ loggingIn: false });
      }
    }
  };

  onChangeRememberMe = event => {
    const rememberMe = event.target.checked;

    this.setState({ rememberMe });
  };

  render() {
    const {
      username,
      password,
      rememberMe,
      usernameValidation,
      loggingIn,
    } = this.state;

    return (
      <form onSubmit={this.onLogin}>
        <TextField
          label={usernameValidation || 'Username'}
          error={!!usernameValidation}
          helperText="Max.MUSTER"
          margin="normal"
          fullWidth
          onChange={this.onChangeUsername}
          value={username}
        />
        <TextField
          label="Password"
          type="password"
          helperText="Your password will NOT be saved"
          margin="normal"
          fullWidth
          onChange={this.onChangePassword}
          value={password}
        />
        <FormControlLabel
          control={
            <Switch checked={rememberMe} onChange={this.onChangeRememberMe} />
          }
          label="Remember me"
        />

        <LoadingButton loading={loggingIn}>
          <DoneIcon />
          Signup
        </LoadingButton>
      </form>
    );
  }
}

export default withStyles(styles)(Login);
