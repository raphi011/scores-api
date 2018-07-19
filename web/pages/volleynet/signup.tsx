import React from 'react';

import { withStyles } from '@material-ui/core/styles';
import Typography from '@material-ui/core/Typography';

import withAuth from '../../containers/AuthContainer';
import Layout from '../../containers/LayoutContainer';
import SearchPlayer from '../../components/volleynet/SearchPlayer';
import Login from '../../components/volleynet/Login';
import {
  loadTournamentAction,
  tournamentSignupAction,
} from '../../redux/actions/entities';
import { tournamentSelector } from '../../redux/reducers/entities';

import { Tournament, VolleynetPlayer } from '../../types';

const styles = () => ({});

interface Props {
  tournamentId: number;
  tournament?: Tournament;
  signup: () => void;
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

    return { tournament };
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
    const { tournament } = this.props;

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
