import React from 'react';

import withAuth from '../../containers/AuthContainer';
import Layout from '../../containers/LayoutContainer';

class UserSettings extends React.Component {
  render() {
    return <Layout title={{ text: 'Settings', href: '' }}>TODO</Layout>;
  }
}

export default withAuth(UserSettings);
