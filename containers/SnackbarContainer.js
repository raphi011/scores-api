import { connect } from "react-redux";
import { withStyles } from "material-ui/styles";
import { statusSelector } from "../redux/reducers/reducer";
import { clearStatusAction } from "../redux/actions/action";
import Snackbar from "../components/Snackbar";

function mapStateToProps(state) {
  const status = statusSelector(state);

  return {
    status,
    open: !!status
  };
}

const mapDispatchToProps = {
  onClose: clearStatusAction
};

export default connect(mapStateToProps, mapDispatchToProps)(Snackbar);
