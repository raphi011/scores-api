// @flow

import React from 'react';
import Link from 'next/link';
import fetch from 'isomorphic-unfetch';

import { withStyles } from '@material-ui/core/styles';
import Button from '@material-ui/core/Button';

import withAuth from '../../containers/AuthContainer';
import Layout from '../../containers/LayoutContainer';
import TournamentView from '../../components/volleynet/TournamentView';
import { BACKEND_URL } from '../../utils/env';

import type { FullTournament } from '../../types';

const styles = theme => ({});

type Props = {
  id: string,
};

type State = {
  tournament: ?FullTournament,
};

class Tournament extends React.Component<Props, State> {
  static getParameters(query) {
    const { id } = query;

    return { id };
  }

  state = {
    tournament: null,
  };

  async componentDidMount() {
    const { id } = this.props;

    const response = await fetch(
      `${BACKEND_URL}/api/volleynet/tournaments/${id}`,
    );

    const tournament = await response.json();

    this.setState({ tournament });
  }

  render() {
    const { tournament } = this.state;

    return (
      <Layout title="New Match">
        <Link prefetch href="/volleynet">
          <Button color="primary">Back</Button>
        </Link>
        <TournamentView tournament={tournament} />
      </Layout>
    );
  }
}

export default withAuth(withStyles(styles)(Tournament));
