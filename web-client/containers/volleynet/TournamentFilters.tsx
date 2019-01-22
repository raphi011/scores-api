import React from 'react';

import { connect } from 'react-redux';

import TournamentFilters from '../../components/volleynet/filters/TournamentFilters';

interface Props {
  onSubmit: () => void;
}

interface State {
  bodyScrolled: boolean;
}

class TournamentFiltersContainer extends React.Component<Props, State> {

  onSubmit = (e: React.FormEvent<HTMLFormElement>) => {
    e.preventDefault();
  };

  render() {
    return (
      <TournamentFilters onSubmit={this.onSubmit} />
    );
  }
}

function mapStateToProps(/* state */) {
  return {};
}

const mapDispatchToProps = {
};

export default connect(mapStateToProps, mapDispatchToProps)(TournamentFiltersContainer);
