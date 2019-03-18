import React from 'react';

import { connect } from 'react-redux';
import { User } from '../types';

import { userSelector } from '../redux/auth/selectors';
import { Store } from '../redux/store';

interface Props {
  children: JSX.Element;
  user: User | null;
}

const AdminOnly = ({ user, children }: Props) => {
  if (user && user.role === 'admin') {
    return children;
  }

  return <div />;
};

function mapStateToProps(state: Store) {
  const user = userSelector(state);

  return { user };
}

export default connect(mapStateToProps)(AdminOnly);
