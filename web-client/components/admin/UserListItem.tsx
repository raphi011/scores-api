import React from 'react';

import ListItem from '@material-ui/core/ListItem';
import ListItemText from '@material-ui/core/ListItemText';

import { User } from '../../types';

interface Props {
  user: User;
  onClick?: (User: User) => void;
}

const onClickHandler = (handler: (user: User) => void, user: User) => {
  if (!handler) {
    return undefined;
  }

  return () => handler(user);
};

const UserListItem = ({ user, onClick }: Props) => (
  <ListItem button onClick={onClickHandler(onClick, user)}>
    <ListItemText primary={user.email} />
  </ListItem>
);

export default UserListItem;
