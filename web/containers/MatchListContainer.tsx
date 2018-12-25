import React from 'react';

import CircularProgress from '@material-ui/core/CircularProgress';
import Router from 'next/router';
import { connect } from 'react-redux';
import Waypoint from 'react-waypoint';

import MatchList from '../components/MatchList';
import MatchOptionsDialog from '../components/MatchOptionsDialog';

import { Match } from '../types';

interface Props {
  matches: Match[];
  // deleteMatch: (Match) => void;
  onLoadMore: () => void;
  highlightPlayerId: number;
  loading: boolean;
  hasMore: boolean;
}

interface IState {
  selectedMatch?: Match;
  dialogOpen: boolean;
}

class MatchListContainer extends React.PureComponent<Props, IState> {
  state = {
    dialogOpen: false,
    selectedMatch: null,
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
    // const { deleteMatch } = this.props;
    // const { selectedMatch } = this.state;
    // if (!selectedMatch) {
    //   return;
    // }
    // deleteMatch(selectedMatch);
    // this.setState({ selectedMatch: null });
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
      <>
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
      </>
    );
  }
}

const mapDispatchToProps = {};

export default connect(
  null,
  mapDispatchToProps,
)(MatchListContainer);
