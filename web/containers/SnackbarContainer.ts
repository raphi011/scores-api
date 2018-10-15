import { connect } from 'react-redux';
import Snackbar from '../components/Snackbar';
import { clearStatusAction } from '../redux/actions/status';
import { statusSelector } from '../redux/reducers/status';

function mapStateToProps(state) {
  const status = statusSelector(state);

  return {
    status,
    open: !!status,
  };
}

const mapDispatchToProps = {
  onClose: clearStatusAction,
};

export default connect(
  mapStateToProps,
  mapDispatchToProps,
)(Snackbar);
