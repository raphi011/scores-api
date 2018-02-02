// @flow

import React from 'react';
import { withStyles } from 'material-ui/styles';
import withRedux from 'next-redux-wrapper';
import Button from 'material-ui/Button';
import AddIcon from 'material-ui-icons/Add';
import Tooltip from 'material-ui/Tooltip';
import Link from 'next/link';

import withRoot from '../styles/withRoot';
import Layout from '../components/Layout';
import MatchList from '../containers/MatchListContainer';
import initStore, { dispatchActions } from '../redux/store';
import { matchesSelector } from '../redux/reducers/reducer';
import {
  loadMatchesAction,
  setStatusAction,
  userOrLoginRouteAction,
} from '../redux/actions/action';
import type { Match, Classes } from '../types';

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
  classes: Classes,
  matches: Array<Match>,
};

class Index extends React.Component<Props> {
  static async getInitialProps({ store, req, res, isServer }) {
    const actions = [loadMatchesAction(), userOrLoginRouteAction()];

    await dispatchActions(store.dispatch, isServer, req, res, actions);
  }

  render() {
    const { matches, classes } = this.props;

    return (
      <Layout title="Matches">
        <div className={classes.matchListContainer}>
          <MatchList matches={matches} />
        </div>
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
};

export default withRedux(initStore, mapStateToProps, mapDispatchToProps)(
  withRoot(withStyles(styles)(Index)),
);
