import React from 'react';

import ListItem from '@material-ui/core/ListItem';
import ListItemText from '@material-ui/core/ListItemText';

import { User } from '../../types';

interface IProps {
  user: User;
  onClick: (User) => void;
}

const UserListItem = ({ user, onClick }: IProps) => (
  <ListItem button onClick={() => onClick(user)}>
    <ListItemText primary={user.email} />
  </ListItem>
);

export default UserListItem;
