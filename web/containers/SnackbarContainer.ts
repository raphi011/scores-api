import { connect } from 'react-redux';

import Snackbar from '../components/Snackbar';
import { clearStatusAction } from '../redux/status/actions';
import { statusSelector } from '../redux/status/reducer';

function mapStateToProps(state) {
  const status = statusSelector(state);

  return {
    open: !!status,
    status,
  };
}

const mapDispatchToProps = {
  onClose: clearStatusAction,
};

export default connect(
  mapStateToProps,
  mapDispatchToProps,
)(Snackbar);
