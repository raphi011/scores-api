import React from 'react';

import ListItem from '@material-ui/core/ListItem';
import ListItemText from '@material-ui/core/ListItemText';

import { User } from '../../types';

interface IProps {
  user: User;
  onClick?: (User) => void;
}

const onClickHandler = handler => {
  if (!handler) {
    return undefined;
  }

  return handler;
};

const UserListItem = ({ user, onClick }: IProps) => (
  <ListItem button onClick={onClickHandler(onClick)}>
    <ListItemText primary={user.email} />
  </ListItem>
);

export default UserListItem;
