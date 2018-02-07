// @flow

import React from 'react';
import { connect } from 'react-redux';
import Router from 'next/router';
import Waypoint from 'react-waypoint';
import { CircularProgress } from 'material-ui/Progress';

import MatchList from '../components/MatchList';
import MatchOptionsDialog from '../components/MatchOptionsDialog';
import { deleteMatchAction } from '../redux/actions/entities';

import type { Match } from '../types';

type Props = {
  matches: Array<Match>,
  deleteMatch: Match => void,
  onLoadMore: () => void,
  loading: boolean,
  hasMore: boolean,
};

type State = {
  selectedMatch: ?Match,
};

class MatchListContainer extends React.Component<Props, State> {
  state = {
    selectedMatch: null,
  };

  onCloseDialog = () => {
    this.setState({ selectedMatch: null });
  };

  onShowPlayer = (playerId: number) => {
    this.setState({ selectedMatch: null });
    Router.push(`/player?id=${playerId}`);
  };

  onOpenDialog = (selectedMatch: Match) => {
    this.setState({ selectedMatch });
  };

  onDeleteMatch = () => {
    const { deleteMatch } = this.props;
    const { selectedMatch } = this.state;

    if (!selectedMatch) return;

    deleteMatch(selectedMatch);

    this.setState({ selectedMatch: null });
  };

  render() {
    const { matches, onLoadMore, loading, hasMore } = this.props;
    const { selectedMatch } = this.state;

    return (
      <React.Fragment>
        <MatchList matches={matches} onMatchClick={this.onOpenDialog} />
        <div style={{ height: '50px', textAlign: 'center' }}>
          {loading ? (
            <CircularProgress />
          ) : hasMore ? (
            <Waypoint onEnter={onLoadMore} />
          ) : null}
        </div>
        <MatchOptionsDialog
          open={selectedMatch != null}
          match={selectedMatch}
          onClose={this.onCloseDialog}
          onDelete={this.onDeleteMatch}
          onShowPlayer={this.onShowPlayer}
        />
      </React.Fragment>
    );
  }
}

const mapDispatchToProps = {
  deleteMatch: deleteMatchAction,
};

export default connect(null, mapDispatchToProps)(MatchListContainer);
