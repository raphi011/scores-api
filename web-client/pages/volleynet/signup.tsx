import React from 'react';

import Link from 'next/link';

import { Theme } from '@material-ui/core';
import Card from '@material-ui/core/Card';
import CardContent from '@material-ui/core/CardContent';
import { createStyles, WithStyles, withStyles } from '@material-ui/core/styles';
import Typography from '@material-ui/core/Typography';

import Login from '../../components/volleynet/Login';
import SearchPlayer from '../../components/volleynet/SearchPlayer';
import withAuth from '../../containers/AuthContainer';
import Layout from '../../containers/LayoutContainer';
import { userSelector } from '../../redux/auth/selectors';
import {
  loadTournamentAction,
  tournamentSignupAction,
} from '../../redux/entities/actions';
import { tournamentSelector } from '../../redux/entities/selectors';
import { card, link, title } from '../../styles/shared';

import { TournamentInfo, User, VolleynetPlayer } from '../../types';

const styles = (theme: Theme) =>
  createStyles({
    card,
    container: {
      padding: theme.spacing.unit * 2,
    },
    link,
    title: title(theme),
  });

interface Props extends WithStyles<typeof styles> {
  tournamentId: number;
  tournament?: TournamentInfo;
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
}

type State = {
  partner?: VolleynetPlayer;
};

class Signup extends React.Component<Props, State> {
  static mapDispatchToProps = {
    signup: tournamentSignupAction,
  };
  static getParameters(query) {
    const { id } = query;

    const tournamentId = Number(id);

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
      partnerId,
      partnerName,
      password,
      rememberMe,
      tournamentId,
      username,
    };

    signup(body);
  };

  render() {
    const { partner } = this.state;
    const { tournament, user, classes } = this.props;

    if (!tournament) {
      return null;
    }

    let content: React.ReactNode;

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
          href: `/volleynet/tournament?id=${tournament.id}`,
          text: 'Signup',
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
          <Card className={classes.card}>
            <CardContent>{content}</CardContent>
          </Card>
        </div>
      </Layout>
    );
  }
}

export default withAuth(withStyles(styles)(Signup));
