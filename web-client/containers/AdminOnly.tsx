import { ReactElement } from 'react';

import { connect } from 'react-redux';
import { User } from '../types';

import { userSelector } from '../redux/auth/selectors';
import { Store } from '../redux/store';

type Props = {
  children: ReactElement<any>;
  user: User;
  isLoggedIn: boolean;
};

const AdminOnly = ({ isLoggedIn, user, children }: Props) => {
  if (isLoggedIn && user.role === 'admin') {
    return children;
  }

  return <div />;
};

function mapStateToProps(state: Store) {
  const { user, isLoggedIn } = userSelector(state);

  return { user, isLoggedIn };
}

export default connect(mapStateToProps)(AdminOnly);
