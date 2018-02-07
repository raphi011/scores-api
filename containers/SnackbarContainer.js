// @flow

import { connect } from 'react-redux';
import { statusSelector } from '../redux/reducers/status';
import { clearStatusAction } from '../redux/actions/status';
import Snackbar from '../components/Snackbar';

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

export default connect(mapStateToProps, mapDispatchToProps)(Snackbar);
