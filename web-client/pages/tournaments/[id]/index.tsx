import React from 'react';

import Link from 'next/link';

import { createStyles, withStyles, WithStyles } from '@material-ui/core/styles';
import Typography from '@material-ui/core/Typography';
import withWidth from '@material-ui/core/withWidth';
import { Breakpoint } from '@material-ui/core/styles/createBreakpoints';

import withAuth from '../../../hoc/next/withAuth';
import Layout from '../../../containers/LayoutContainer';

import { userSelector } from '../../../redux/auth/selectors';
import { loadTournamentAction } from '../../../redux/entities/actions';
import { tournamentSelector } from '../../../redux/entities/selectors';
import { Tournament, User } from '../../../types';
import { Store } from '../../../redux/store';
import TeamList from '../../../components/volleynet/TeamList';
import * as Query from '../../../utils/query';
import withConnect, { Context } from '../../../hoc/next/withConnect';
import TournamentHeader from '../../../components/volleynet/TournamentHeader';
import { isSignedup } from '../../../utils/tournament';
import Grid from '@material-ui/core/Grid';

const styles = createStyles({
  body: {
    marginTop: '30px',
  },
  title: {
    marginBottom: 0,
  },
  columnHeader: {
    marginBottom: '20px',
  },
  notes: {
    '& img': {
      maxWidth: '100%',
    },
  },
});

interface Props extends WithStyles<typeof styles> {
  tournament?: Tournament;
  tournamentId: string;
  user: User;
  width: Breakpoint;
}

class ShowTournament extends React.Component<Props> {
  static async getInitialProps(ctx: Context): Promise<Partial<Props>> {
    const { query } = ctx;

    const tournamentId = Query.one(query, 'id');

    return { tournamentId };
  }

  static buildActions({ tournamentId }: Props) {
    return [loadTournamentAction(tournamentId)];
  }

  static mapStateToProps(state: Store, { tournamentId }: Props) {
    const tournament = tournamentSelector(state, tournamentId);
    const user = userSelector(state);

    return { tournament, user };
  }

  hasNotes = (notes: string) => {
    return notes.trim().length > 0;
  };

  renderBody = () => {
    const { classes, user, tournament, width } = this.props;

    const isMobile = ['xs', 'sm'].includes(width);

    if (!tournament) {
      return null;
    }

    const canSignup =
      tournament.registrationOpen && !isSignedup(tournament, user.playerId);

    const teams = tournament.teams || [];

    return (
      <>
        <TournamentHeader tournament={tournament} showSignup={canSignup} />
        <Grid container spacing={2} className={classes.body}>
          <Grid md={8} xs={12} item>
            <Typography className={classes.columnHeader} variant="h2">
              Notes
            </Typography>
            {this.hasNotes(tournament.htmlNotes) ? (
              <Typography
                variant="body2"
                className={classes.notes}
                dangerouslySetInnerHTML={{ __html: tournament.htmlNotes }}
              />
            ) : (
              <Typography variant="body2">There are no notes yet.</Typography>
            )}
          </Grid>
          <Grid md={4} xs={12} item>
            <Typography className={classes.columnHeader} variant="h2">
              Teams
            </Typography>
            {teams.length ? (
              <TeamList teams={teams} />
            ) : (
              <Typography variant="body2">
                No teams have signed up yet.
              </Typography>
            )}
          </Grid>
        </Grid>
      </>
    );
  };

  render() {
    const { tournament } = this.props;

    if (!tournament) {
      return null;
    }

    return (
      <Layout title={{ text: 'Tournaments', href: '/' }}>
        {this.renderBody()}
      </Layout>
    );
  }
}

export default withAuth(
  withConnect(withWidth()(withStyles(styles)(ShowTournament))),
);
