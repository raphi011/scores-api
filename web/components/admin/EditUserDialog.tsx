import React, { ChangeEvent } from 'react';

import { Button, Dialog, DialogTitle, TextField } from '@material-ui/core';
import { createStyles, WithStyles, withStyles } from '@material-ui/core/styles';

import { User } from '../../types';

const styles = createStyles({
  userForm: {
    display: 'flex',
    flexDirection: 'column',
    padding: '20px',
  },
});

interface Props extends WithStyles<typeof styles> {
  user?: User;
  open: boolean;
  isNew: boolean;
  email: string;
  password: string;
  canSubmit: boolean;

  onSubmit: () => void;
  onClose: () => void;
  onChangeEmail: (event: ChangeEvent<HTMLInputElement>) => void;
  onChangePassword: (event: ChangeEvent<HTMLInputElement>) => void;
}

export default withStyles(styles)(
  ({
    isNew,
    onClose,
    open,
    email,
    onChangeEmail,
    password,
    onChangePassword,
    onSubmit,
    canSubmit,
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
          <Button
            color="primary"
            disabled={!canSubmit}
            onClick={onSubmit}
            variant="contained"
          >
            Submit
          </Button>
        </form>
        <div />
      </Dialog>
    );
  },
);
