// @flow

import React from 'react';
import { withStyles } from 'material-ui/styles';
import withRedux from 'next-redux-wrapper';
import Button from 'material-ui/Button';
import AddIcon from 'material-ui-icons/Add';
import Router from 'next/router';
import Tooltip from 'material-ui/Tooltip';
import Link from 'next/link';

import withRoot from '../components/withRoot';
import Layout from '../components/Layout';
import MatchOptionsDialog from '../components/MatchOptionsDialog';
import MatchList from '../components/MatchList';
import initStore, { dispatchActions } from '../redux/store';
import { matchesSelector } from '../redux/reducers/reducer';
import {
  loadMatchesAction,
  setStatusAction,
  deleteMatchAction,
  userOrLoginRouteAction,
} from '../redux/actions/action';
import type { Match } from '../types';

const styles = theme => ({
  matchListContainer: {
    marginBottom: '70px',
  },
  button: {
    margin: theme.spacing.unit,
    position: 'fixed',
    right: '24px',
    bottom: '24px',
  },
});

type Props = {
  deleteMatch: Match => void,
  matches: Array<Match>,
  classes: Object,
};

type State = {
  selectedMatch: ?Match,
};

class Index extends React.Component<Props, State> {
  static async getInitialProps({ store, req, res, isServer }) {
    const actions = [loadMatchesAction(), userOrLoginRouteAction()];

    await dispatchActions(store.dispatch, isServer, req, res, actions);
  }

  state = {
    selectedMatch: null,
  };

  onCloseDialog = () => {
    this.setState({ selectedMatch: null });
  };

  onShowPlayer = (playerId: number) => {
    Router.push(`/player?id=${playerId}`);
  };

  onOpenDialog = (selectedMatch: Match) => {
    this.setState({ selectedMatch });
  };

  onCreateMatch = () => {
    Router.replace('/createMatch');
  };

  onDeleteMatch = () => {
    const { deleteMatch } = this.props;
    const { selectedMatch } = this.state;

    if (!selectedMatch) return;

    deleteMatch(selectedMatch);

    this.setState({ selectedMatch: null });
  };

  render() {
    const { matches, classes } = this.props;
    const { selectedMatch } = this.state;

    return (
      <Layout title="Matches">
        <div className={classes.matchListContainer}>
          <MatchList matches={matches} onMatchClick={this.onOpenDialog} />
        </div>
        <MatchOptionsDialog
          open={selectedMatch != null}
          match={selectedMatch}
          onClose={this.onCloseDialog}
          onDelete={this.onDeleteMatch}
          onShowPlayer={this.onShowPlayer}
        />
        <Tooltip title="Create new Match" className={classes.button}>
          <Button fab color="primary" aria-label="add">
            <Link prefetch href="/createMatch">
              <AddIcon />
            </Link>
          </Button>
        </Tooltip>
      </Layout>
    );
  }
}

function mapStateToProps(state) {
  const matches = matchesSelector(state);

  return {
    matches,
  };
}

const mapDispatchToProps = {
  loadMatches: loadMatchesAction,
  setStatus: setStatusAction,
  deleteMatch: deleteMatchAction,
};

export default withRedux(initStore, mapStateToProps, mapDispatchToProps)(
  withRoot(withStyles(styles)(Index)),
);
