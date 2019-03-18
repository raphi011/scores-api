import React from 'react';

import withAuth from '../../hoc/next/withAuth';
import Layout from '../../containers/LayoutContainer';

import { userSelector } from '../../redux/auth/selectors';
import { Store } from '../../redux/store';
import { User } from '../../types';
import withConnect from '../../hoc/next/withConnect';

interface Props {
  user: User;
}

class Ranking extends React.Component<Props> {
  static mapStateToProps(state: Store) {
    const user = userSelector(state);

    return { user };
  }

  render() {
    const { user } = this.props;

    return <Layout title={{ text: 'User', href: '' }}>{user.email}</Layout>;
  }
}

export default withAuth(withConnect(Ranking));
