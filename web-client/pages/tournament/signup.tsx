import React from 'react';

import Router from 'next/router';

import { Grid, Dialog } from '@material-ui/core';
import { createStyles, WithStyles, withStyles } from '@material-ui/core/styles';
import Typography from '@material-ui/core/Typography';

import Login from '../../components/volleynet/Login';
import SearchPlayer from '../../components/volleynet/SearchPlayer';
import withAuth from '../../hoc/next/withAuth';
import Layout from '../../containers/LayoutContainer';
import {
  loadTournamentAction,
  tournamentSignupAction,
  previousPartnersAction,
} from '../../redux/entities/actions';
import {
  tournamentSelector,
  previousPartnersSelector,
  searchVolleynetplayerSelector,
} from '../../redux/entities/selectors';

import { Tournament, User, Player } from '../../types';
import { Store } from '../../redux/store';
import * as Query from '../../utils/query';
import withConnect, { Context } from '../../hoc/next/withConnect';
import TournamentHeader from '../../components/volleynet/TournamentHeader';
import SearchPlayerList from '../../components/volleynet/SearchPlayerList';

const styles = createStyles({
  body: {
    marginTop: '30px',
  },
  header: {
    marginBottom: '20px',
  },
  info: {
    verticalAlign: 'middle',
    height: '100%',
  },
});

interface Props extends WithStyles<typeof styles> {
  tournamentId: string;
  tournament?: Tournament;
  user: User;
  previousPartners: Player[];
  foundPlayers: Player[];

  loadTeamPartners: (playerId: number) => void;
  signup: (info: {
    username: string;
    password: string;
    partnerId: number;
    tournamentId: number;
    rememberMe: boolean;
  }) => void;
}

interface State {
  partner: Player | null;
}

class Signup extends React.Component<Props, State> {
  static mapDispatchToProps = {
    signup: tournamentSignupAction,
    loadTeamPartners: previousPartnersAction,
  };

  static async getInitialProps(ctx: Context) {
    const { query } = ctx;
    const tournamentId = Query.one(query, 'id');

    return { tournamentId };
  }

  static buildActions({ tournamentId, user }: Props) {
    const actions = [previousPartnersAction(user.playerId)];

    if (tournamentId) {
      actions.push(loadTournamentAction(tournamentId));
    }

    return actions;
  }

  static mapStateToProps(state: Store, { tournamentId, user }: Props) {
    const tournament = tournamentSelector(state, tournamentId);
    const previousPartners = previousPartnersSelector(state, user.playerId);
    const foundPlayers = searchVolleynetplayerSelector(state);

    return { tournament, user, previousPartners, foundPlayers };
  }

  state: State = {
    partner: null,
  };

  onSelectPlayer = (partner: Player | null) => {
    this.setState({ partner });
  };

  onSignup = async (
    username: string,
    password: string,
    rememberMe: boolean,
  ) => {
    const { tournamentId, signup } = this.props;
    const { partner } = this.state;

    const partnerId = partner && partner.id;

    if (!partnerId) {
      return;
    }

    const body = {
      partnerId,
      tournamentId: Number(tournamentId),

      password,
      rememberMe,
      username,
    };

    try {
      await signup(body);
      Router.push({
        pathname: '/tournament',
        query: { id: tournamentId },
      });
    } catch (e) {
      // eslint-disable-next-line
      console.log(e);
    }
  };

  onCloseSignup = () => {
    this.setState({ partner: null });
  };

  renderDialog = () => {
    const { partner } = this.state;
    const { user } = this.props;

    const open = !!partner;

    let body = <span />;

    if (partner) {
      body = (
        <>
          <Typography variant="h6">{`Partner: ${partner.firstName} ${
            partner.lastName
          }`}</Typography>
          <Login onLogin={this.onSignup} username={user.playerLogin} />
        </>
      );
    }

    return (
      <Dialog open={open} onClose={this.onCloseSignup}>
        {body}
      </Dialog>
    );
  };

  render() {
    const { foundPlayers, previousPartners, tournament, classes } = this.props;

    if (!tournament) {
      return null;
    }

    const players =
      foundPlayers && foundPlayers.length ? foundPlayers : previousPartners;

    return (
      <Layout
        title={{
          href: `/tournament?id=${tournament.id}`,
          text: 'Signup',
        }}
      >
        {this.renderDialog()}
        <TournamentHeader tournament={tournament} />
        <div className={classes.body}>
          <Typography color="primary" variant="h1" className={classes.header}>
            Signup
          </Typography>
          <Grid container spacing={16}>
            <Grid item xs={12} sm={4}>
              <Typography variant="h3">Search</Typography>
              <SearchPlayer gender={tournament.gender} />
            </Grid>
            <Grid item xs={12} sm={4}>
              <Typography variant="h3">Partner</Typography>
              <SearchPlayerList
                players={players}
                onPlayerClick={this.onSelectPlayer}
              />
            </Grid>
          </Grid>
        </div>
      </Layout>
    );
  }
}

export default withAuth(withConnect(withStyles(styles)(Signup)));
