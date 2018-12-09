import React from 'react';

import ListItem from '@material-ui/core/ListItem';
import ListItemText from '@material-ui/core/ListItemText';

import { User } from '../../types';

interface IProps {
  user: User;
  onClick?: (User) => void;
}

const onClickHandler = (handler, user: User) => {
  if (!handler) {
    return undefined;
  }

  return () => handler(user);
};

const UserListItem = ({ user, onClick }: IProps) => (
  <ListItem button onClick={onClickHandler(onClick, user)}>
    <ListItemText primary={user.email} />
  </ListItem>
);

export default UserListItem;
