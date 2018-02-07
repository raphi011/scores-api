// @flow

import { connect } from 'react-redux';
import { loginRouteSelector, userSelector } from '../redux/reducers/auth';
import { logoutAction } from '../redux/actions/auth';
import AppBar from '../components/AppBar';

function mapStateToProps(state) {
  const { isLoggedIn, user } = userSelector(state);
  const loginRoute = loginRouteSelector(state);

  return {
    isLoggedIn,
    user,
    loginRoute,
  };
}

const mapDispatchToProps = {
  onLogout: logoutAction,
};

export default connect(mapStateToProps, mapDispatchToProps)(AppBar);
