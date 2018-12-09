import React, { ChangeEvent } from 'react';

import { Dialog, DialogTitle, TextField } from '@material-ui/core';
import { User } from '../../types';

type Props = {
  user?: User;
  open: boolean;
  isNew: boolean;

  email: string;
  password: string;

  onClose: () => void;
  onChangeEmail: (event: ChangeEvent<HTMLInputElement>) => void;
  onChangePassword: (event: ChangeEvent<HTMLInputElement>) => void;
};

export default ({
  isNew,
  onClose,
  open,
  email,
  onChangeEmail,
  password,
  onChangePassword,
}: Props) => {
  const title = isNew ? 'New User' : 'Edit User';

  return (
    <Dialog onClose={onClose} open={open}>
      <DialogTitle>{title}</DialogTitle>
      <form noValidate autoComplete="off">
        <TextField label="Email" value={email} onChange={onChangeEmail} />
        <TextField
          label="Password"
          type="password"
          margin="normal"
          value={password}
          onChange={onChangePassword}
        />
      </form>
      <div />
    </Dialog>
  );
};
