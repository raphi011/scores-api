// @flow

import React from 'react';
import Link from 'next/link';
import fetch from 'isomorphic-unfetch';

import { withStyles } from 'material-ui/styles';
import Button from 'material-ui/Button';

import withAuth from '../../containers/AuthContainer';
import Layout from '../../containers/LayoutContainer';
import TournamentView from '../../components/volleynet/TournamentView';

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
      `http://localhost:3000/api/volleynet/tournaments/${id}`,
    );

    const tournament = await response.json();

    this.setState({ tournament });
  }

  render() {
    const { tournament } = this.state;

    return (
      <Layout title="New Match">
        <Link prefetch href="/volleynet">
          <Button href="#flat-buttons">Back</Button>
        </Link>
        <TournamentView tournament={tournament} />
      </Layout>
    );
  }
}

export default withAuth(withStyles(styles)(Tournament));
