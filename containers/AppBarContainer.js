import { connect } from "react-redux";
import { userSelector } from "../redux/reducers/reducer";
import { logoutAction } from "../redux/actions/action";
import AppBar from "../components/AppBar";

function mapStateToProps(state) {
  const { isLoggedIn, user } = userSelector(state);

  return {
    isLoggedIn,
    user
  };
}

const mapDispatchToProps = {
  onLogout: logoutAction
};

export default connect(mapStateToProps, mapDispatchToProps)(AppBar);
