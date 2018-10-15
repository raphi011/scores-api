import React from 'react';

import { createStyles, withStyles } from '@material-ui/core/styles';
import Typography from '@material-ui/core/Typography';
import Link from 'next/link';

import Login from '../../components/volleynet/Login';
import SearchPlayer from '../../components/volleynet/SearchPlayer';
import withAuth from '../../containers/AuthContainer';
import Layout from '../../containers/LayoutContainer';
import {
  loadTournamentAction,
  tournamentSignupAction,
} from '../../redux/actions/entities';
import { userSelector } from '../../redux/reducers/auth';
import { tournamentSelector } from '../../redux/reducers/entities';
import { card, link, title } from '../../styles/shared';

import { Card, Theme } from '@material-ui/core';
import { Tournament, User, VolleynetPlayer } from '../../types';

const styles = (theme: Theme) =>
  createStyles({
    title: title(theme),
    link,
    card,
    container: {
      padding: theme.spacing.unit * 2,
    },
  });

interface IProps {
  tournamentId: number;
  tournament?: Tournament;
  user: User;
  signup: (
    info: {
      username: string;
      password: string;
      partnerId: number;
      tournamentId: number;
      partnerName: string;
      rememberMe: boolean;
    },
  ) => void;
  classes: any;
}

interface State {
  partner?: VolleynetPlayer;
}

class Signup extends React.Component<IProps, State> {

  static mapDispatchToProps = {
    signup: tournamentSignupAction,
  };
  static getParameters(query) {
    const { id } = query;

    const tournamentId = Number.parseInt(id, 10);

    return { tournamentId };
  }

  static buildActions({ tournamentId }) {
    return [loadTournamentAction(tournamentId)];
  }

  static mapStateToProps(state, { tournamentId }) {
    const tournament = tournamentSelector(state, tournamentId);
    const { user } = userSelector(state);

    return { tournament, user };
  }

  state = {
    partner: null,
  };

  onSelectPlayer = partner => {
    this.setState({ partner });
  };

  onSignup = async (username, password, rememberMe) => {
    const { tournamentId, signup } = this.props;
    const { partner } = this.state;

    const partnerId = partner && partner.id;
    const partnerName = partner && partner.login;

    const body = {
      username,
      password,
      partnerId,
      tournamentId,
      partnerName,
      rememberMe,
    };

    signup(body);
  };

  render() {
    const { partner } = this.state;
    const { tournament, user, classes } = this.props;

    if (!tournament) {
      return null;
    }

    let content;

    if (partner) {
      content = (
        <>
          <Typography variant="h6">{`Partner: ${partner.firstName} ${
            partner.lastName
          }`}</Typography>
          <Login onLogin={this.onSignup} username={user.volleynetLogin} />
        </>
      );
    } else {
      content = (
        <SearchPlayer
          gender={tournament.gender}
          onSelectPlayer={this.onSelectPlayer}
        />
      );
    }

    return (
      <Layout
        title={{
          text: 'Signup',
          href: `/volleynet/tournament?id=${tournament.id}`,
        }}
      >
        <div className={classes.container}>
          <div className={`${classes.title} ${classes.link}`}>
            <Link href={`/volleynet/tournament?id=${tournament.id}`}>
              <a className={classes.link}>
                <Typography variant="h4">{tournament.name}</Typography>
              </a>
            </Link>
          </div>
          <Card className={classes.card}>{content}</Card>
        </div>
      </Layout>
    );
  }
}

export default withAuth(withStyles(styles)(Signup));
