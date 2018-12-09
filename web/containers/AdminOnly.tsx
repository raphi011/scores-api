import { ReactElement } from 'react';

import { connect } from 'react-redux';
import { User } from '../types';

import { userSelector } from '../redux/auth/selectors';

type Props = {
  user: User;
  children: ReactElement<any>;
};

const AdminOnly = ({ user, children }: Props) => {
  if (user.role === 'admin') {
    return children;
  }

  return <div />;
};

function mapStateToProps(state) {
  const { user } = userSelector(state);

  return { user };
}

export default connect(mapStateToProps)(AdminOnly);
