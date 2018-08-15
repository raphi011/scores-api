import React from 'react';

import { withStyles, createStyles } from '@material-ui/core/styles';
import Typography from '@material-ui/core/Typography';

import { userSelector } from '../../redux/reducers/auth';
import withAuth from '../../containers/AuthContainer';
import Layout from '../../containers/LayoutContainer';
import SearchPlayer from '../../components/volleynet/SearchPlayer';
import Login from '../../components/volleynet/Login';
import {
  loadTournamentAction,
  tournamentSignupAction,
} from '../../redux/actions/entities';
import { tournamentSelector } from '../../redux/reducers/entities';

import { Tournament, VolleynetPlayer, User } from '../../types';

const styles = createStyles({});

interface Props {
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

class Signup extends React.Component<Props, State> {
  static getParameters(query) {
    const { id } = query;

    const tournamentId = Number.parseInt(id, 10);

    return { tournamentId };
  }

  state = {
    partner: null,
  };

  static buildActions({ tournamentId }) {
    return [loadTournamentAction(tournamentId)];
  }

  static mapStateToProps(state, { tournamentId }) {
    const tournament = tournamentSelector(state, tournamentId);
    const { user } = userSelector(state);

    return { tournament, user };
  }

  static mapDispatchToProps = {
    signup: tournamentSignupAction,
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
    const { tournament, user } = this.props;

    if (!tournament) {
      return null;
    }

    return (
      <Layout title={{ text: 'Signup', href: '' }}>
        <Typography variant="headline">{tournament.name}</Typography>
        {partner ? (
          <>
            <Typography
              variant="title"
              style={{ margin: '20px 0' }}
            >{`Partner: ${partner.firstName} ${partner.lastName}`}</Typography>
            <Login onLogin={this.onSignup} username={user.volleynetLogin} />
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
