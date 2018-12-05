import { ReactNode } from 'react';

import { connect } from 'react-redux';
import { User } from '../types';

import { userSelector } from '../redux/auth/selectors';

interface IProps {
  user: User;
  children: ReactNode;
}

const AdminOnly = ({ user, children }: IProps) => {
  if (user.role === 'admin') {
    return children;
  }

  return null;
};

function mapStateToProps(state) {
  const { user } = userSelector(state);

  return { user };
}

export default connect(mapStateToProps)(AdminOnly);
