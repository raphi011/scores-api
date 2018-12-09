import React from 'react';

import List from '@material-ui/core/List';
import { createStyles, withStyles } from '@material-ui/core/styles';

import { User } from '../../types';
import UserListItem from './UserListItem';

const styles = createStyles({
  root: {
    width: '100%',
  },
});

type Props = {
  users: User[];
  classes: any;
  onClick: (user: User) => void;
};

class UserList extends React.PureComponent<Props> {
  render() {
    const { users = [], onClick, classes } = this.props;

    return (
      <List className={classes.root}>
        {users.map(u => (
          <UserListItem key={u.id} onClick={onClick} user={u} />
        ))}
      </List>
    );
  }
}

export default withStyles(styles)(UserList);
