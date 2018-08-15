import React from 'react';
import TextField from '@material-ui/core/TextField';
import { withStyles, createStyles } from '@material-ui/core/styles';
import Switch from '@material-ui/core/Switch';
import FormControlLabel from '@material-ui/core/FormControlLabel';

import Button from '@material-ui/core/Button';
import DoneIcon from '@material-ui/icons/Done';

const styles = createStyles({
  container: {
    padding: '0 10px',
  },
});

interface Props {
  onLogin: (username: string, password: string, rememberMe: boolean) => void;
  classes: any;
  username?: string;
}

interface State {
  username: string;
  password: string;
  rememberMe: boolean;
  usernameValidation: string;
}

class Login extends React.Component<Props, State> {
  constructor(props) {
    super(props);

    this.state = {
      username: props.username || '',
      password: '',
      rememberMe: true,
      usernameValidation: '',
    };
  }

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
    const { username, password, rememberMe } = this.state;

    if (username && this.loginRegex.test(username) && password) {
      onLogin(username, password, rememberMe);
    }
  };

  onChangeRememberMe = event => {
    const rememberMe = event.target.checked;

    this.setState({ rememberMe });
  };

  render() {
    const { classes } = this.props;
    const { username, password, rememberMe, usernameValidation } = this.state;

    return (
      <div className={classes.container}>
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
