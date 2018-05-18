// @flow

import React from 'react';
// import Link from 'next/link';
import fetch from 'isomorphic-unfetch';
import Router from 'next/router';

import { withStyles } from '@material-ui/core/styles';
import Typography from '@material-ui/core/Typography';

import withAuth from '../../containers/AuthContainer';
import Layout from '../../containers/LayoutContainer';
import SearchPlayer from '../../components/volleynet/SearchPlayer';
import Login from '../../components/volleynet/Login';

import type { FullTournament, VolleynetPlayer } from '../../types';

const styles = theme => ({});

type Props = {
  id: string,
};

type State = {
  tournament: ?FullTournament,
  partner: ?VolleynetPlayer,
};

class Signup extends React.Component<Props, State> {
  static getParameters(query) {
    const { id } = query;

    return { id };
  }

  state = {
    tournament: null,
    partner: null,
  };

  async componentDidMount() {
    const { id } = this.props;

    const response = await fetch(
      `http://localhost:3000/api/volleynet/tournaments/${id}`,
    );

    const tournament = await response.json();

    this.setState({ tournament });
  }

  onSelectPlayer = partner => {
    this.setState({ partner });
  };

  onSignup = async (username, password) => {
    const { id: tournamentId } = this.props;
    const { partner } = this.state;

    const partnerId = partner && partner.id;
    const partnerName = partner && partner.login;

    const body = {
      username,
      password,
      partnerId,
      tournamentId,
      partnerName,
    };

    const response = await fetch('http://localhost:3000/api/volleynet/signup', {
      body: JSON.stringify(body),
      method: 'POST',
    });

    if (response.status !== 200) {
      // TODO: set message
      await Router.push({
        pathname: '/volleynet/tourname',
        query: { id: tournamentId },
      });
    }
  };

  render() {
    const { tournament, partner } = this.state;

    if (!tournament) {
      return null;
    }

    return (
      <Layout title="Signup">
        <Typography variant="headline">{tournament.name}</Typography>
        {partner ? (
          <>
            <Typography
              variant="title"
              style={{ margin: '20px 0' }}
            >{`Partner: ${partner.firstName} ${partner.lastName}`}</Typography>
            <Login onLogin={this.onSignup} />
          </>
        ) : (
          <SearchPlayer
            gender={tournament.gender}
            onSelectPlayer={this.onSelectPlayer}
          />
        )}
      </Layout>
    );
  }
}

export default withAuth(withStyles(styles)(Signup));
