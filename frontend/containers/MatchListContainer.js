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

const Fragment = React.Fragment;

type Props = {
  matches: Array<Match>,
  deleteMatch: Match => void,
  onLoadMore: () => void,
  highlightPlayerId: number,
  loading: boolean,
  hasMore: boolean,
};

type State = {
  selectedMatch: ?Match,
  dialogOpen: boolean,
};

class MatchListContainer extends React.PureComponent<Props, State> {
  state = {
    selectedMatch: null,
    dialogOpen: false,
  };

  onCloseDialog = () => {
    this.setState({ dialogOpen: false });
  };

  onShowPlayer = async (playerId: number) => {
    this.onCloseDialog();
    await Router.push(`/player?id=${playerId}`);
    scroll(0, 0);
  };

  onOpenDialog = (selectedMatch: Match) => {
    this.setState({ dialogOpen: true, selectedMatch });
  };

  onDeleteMatch = () => {
    const { deleteMatch } = this.props;
    const { selectedMatch } = this.state;

    if (!selectedMatch) return;

    deleteMatch(selectedMatch);

    this.setState({ selectedMatch: null });
  };

  render() {
    const {
      matches,
      onLoadMore,
      loading,
      hasMore,
      highlightPlayerId,
    } = this.props;
    const { selectedMatch, dialogOpen } = this.state;

    return (
      <Fragment>
        <MatchList
          matches={matches}
          onMatchClick={this.onOpenDialog}
          highlightPlayerId={highlightPlayerId}
        />
        <div style={{ height: '50px', textAlign: 'center' }}>
          {loading ? (
            <CircularProgress />
          ) : hasMore ? (
            <Waypoint onEnter={onLoadMore} />
          ) : null}
        </div>
        <MatchOptionsDialog
          open={dialogOpen}
          match={selectedMatch}
          onClose={this.onCloseDialog}
          onDelete={this.onDeleteMatch}
          onShowPlayer={this.onShowPlayer}
        />
      </Fragment>
    );
  }
}

const mapDispatchToProps = {
  deleteMatch: deleteMatchAction,
};

export default connect(null, mapDispatchToProps)(MatchListContainer);
