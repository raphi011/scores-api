import { connect } from 'react-redux';

import TournamentFilters from '../../components/volleynet/filters/TournamentFilters';

type Props = {
  //   user: User;
  //   children: ReactElement<any>;
};

function mapStateToProps(state) {
  //   const { user } = userSelector(state);

  return {};
}

const mapDispatchToProps = {};

export default connect(mapStateToProps)(TournamentFilters);
