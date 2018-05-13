// @flow

import React from 'react';
import TextField from 'material-ui/TextField';
import { withStyles } from 'material-ui/styles';

import Button from 'material-ui/Button';
import DoneIcon from 'material-ui-icons/Done';

const styles = () => ({
  container: {
    padding: '0 10px',
  },
});

type Props = {
  onLogin: (string, string) => void,
};

type State = {
  username: string,
  password: string,
  usernameValidation: string,
};

class Login extends React.Component<Props, State> {
  state = {
    username: '',
    password: '',
    usernameValidation: '',
  };

  loginRegex = /^[A-Z][a-z]+\.[A-Z]+$/;

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

  onLogin = () => {
    const { onLogin } = this.props;
    const { username, password } = this.state;

    if (username && this.loginRegex.test(username) && password) {
      onLogin(username, password);
    }
  };

  render() {
    const { classes } = this.props;
    const { username, password, usernameValidation } = this.state;

    return (
      <div className={classes.container}>
        <TextField
          label={usernameValidation || 'Username'}
          error={!!usernameValidation}
          helperText="max.MUSTER"
          margin="normal"
          fullWidth
          onChange={this.onChangeUsername}
          value={username}
        />
        <TextField
          label="Password"
          type="password"
          margin="normal"
          fullWidth
          onChange={this.onChangePassword}
          value={password}
        />
        <Button
          color="primary"
          type="submit"
          onClick={this.onLogin}
          variant="raised"
          fullWidth
        >
          <DoneIcon />
          Signup
        </Button>
      </div>
    );
  }
}

export default withStyles(styles)(Login);
