import React, { ChangeEvent } from 'react';

import { Dialog, DialogTitle, TextField, Button } from '@material-ui/core';
import { createStyles, withStyles } from '@material-ui/core/styles';

import { User } from '../../types';

const styles = createStyles({
  userForm: {
    display: 'flex',
    flexDirection: 'column',
    padding: '10px',
  },
});

type Props = {
  user?: User;
  open: boolean;
  isNew: boolean;
  classes: any;

  email: string;
  password: string;

  onClose: () => void;
  onChangeEmail: (event: ChangeEvent<HTMLInputElement>) => void;
  onChangePassword: (event: ChangeEvent<HTMLInputElement>) => void;
};

export default withStyles(styles)(
  ({
    isNew,
    onClose,
    open,
    email,
    onChangeEmail,
    password,
    onChangePassword,
    classes,
  }: Props) => {
    const title = isNew ? 'New User' : 'Edit User';

    return (
      <Dialog onClose={onClose} open={open}>
        <DialogTitle>{title}</DialogTitle>
        <form className={classes.userForm} noValidate autoComplete="off">
          <TextField
            label="Email"
            disabled={!isNew}
            value={email}
            onChange={onChangeEmail}
          />
          <TextField
            label="Password"
            type="password"
            margin="normal"
            value={password}
            onChange={onChangePassword}
          />
          <Button color="primary" variant="contained">
            Submit
          </Button>
        </form>
        <div />
      </Dialog>
    );
  },
);
