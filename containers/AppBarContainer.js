// @flow

import { connect } from "react-redux";
import { userSelector, loginRouteSelector } from "../redux/reducers/reducer";
import { logoutAction } from "../redux/actions/action";
import AppBar from "../components/AppBar";

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
  onLogout: logoutAction
};

export default connect(mapStateToProps, mapDispatchToProps)(AppBar);
