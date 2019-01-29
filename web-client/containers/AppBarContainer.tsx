import React from 'react';

import { connect } from 'react-redux';

import AppBar from '../components/AppBar';
import { userSelector } from '../redux/auth/selectors';
import { Store } from '../redux/store';
import { User } from '../types';

interface Props {
  isLoggedIn: boolean;
  title: { text: string; href: string };
  user: User;

  onToggleDrawer: () => void;
}

interface State {
  bodyScrolled: boolean;
}

class AppBarContainer extends React.Component<Props, State> {
  state = {
    anchorEl: null,
    bodyScrolled: false,
  };

  componentDidMount() {
    window.addEventListener('scroll', this.onScroll);
  }

  onScroll = () => {
    const bodyScrolled = window.scrollY !== 0;

    if (bodyScrolled !== this.state.bodyScrolled) {
      this.setState({ bodyScrolled });
    }
  };

  componentWillUnmount() {
    window.removeEventListener('scroll', this.onScroll);
  }

  render() {
    return <AppBar {...this.props} bodyScrolled={this.state.bodyScrolled} />;
  }
}

function mapStateToProps(state: Store) {
  const { user, isLoggedIn } = userSelector(state);

  return {
    isLoggedIn,
    user,
  };
}

export default connect(mapStateToProps)(AppBarContainer);
