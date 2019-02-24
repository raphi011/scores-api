import React, { ReactNode } from 'react';

import { connect } from 'react-redux';

import Layout from '../components/Layout';

interface Props {
  title: { text: string; href: string };
  children: ReactNode;
}

class LayoutContainer extends React.Component<Props> {
  render() {
    const { title, children } = this.props;

    return <Layout title={title}>{children}</Layout>;
  }
}

function mapStateToProps() {
  return {};
}

export default connect(mapStateToProps)(LayoutContainer);
