import React from 'react';

import Link from 'next/link';

import Fab from '@material-ui/core/Fab';
import Typography from '@material-ui/core/Typography';
import SignupIcon from "@material-ui/icons/PlaylistAdd";

import withAuth from '../../containers/AuthContainer';
import Layout from '../../containers/LayoutContainer';

import { userSelector } from '../../redux/auth/selectors';
import { loadTournamentAction } from '../../redux/entities/actions';
import { tournamentSelector } from '../../redux/entities/selectors';
import { Tournament, User } from '../../types';


interface Props {
  tournament?: Tournament;
  user: User;
}

class ShowTournament extends React.Component<Props> {
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

  render() {
    const { tournament } = this.props;

    return (
      <Layout title={{ text: 'Tournaments', href: '/' }}>
        <Typography variant="h2">{tournament.name}</Typography>
          <Link href={{ pathname: "/tournament/signup", query: { id: tournament.id }}} >
            <Fab size="medium" color="secondary" aria-label="Add" >
              <SignupIcon />
            </Fab>
          </Link>
      </Layout>
    );
  }
}

export default withAuth(ShowTournament);
