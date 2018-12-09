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

interface IProps {
  users: User[];
  classes: any;
}

class UserList extends React.PureComponent<IProps> {
  render() {
    const { users = [], classes } = this.props;

    return (
      <List className={classes.root}>
        {users.map(u => (
          <UserListItem key={u.id} user={u} />
        ))}
      </List>
    );
  }
}

export default withStyles(styles)(UserList);
