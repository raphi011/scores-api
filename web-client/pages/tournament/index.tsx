import React from 'react';

import Button from '@material-ui/core/Button';
import Card from '@material-ui/core/Card';
import { createStyles, withStyles, WithStyles } from '@material-ui/core/styles';
import Typography from '@material-ui/core/Typography';
import withWidth from '@material-ui/core/withWidth';
import External from '@material-ui/icons/OpenInNew';

import withAuth from '../../containers/AuthContainer';
import Layout from '../../containers/LayoutContainer';

import { userSelector } from '../../redux/auth/selectors';
import { loadTournamentAction } from '../../redux/entities/actions';
import { tournamentSelector } from '../../redux/entities/selectors';
import { link } from '../../styles/shared';

import { Tournament, User } from '../../types';

const styles = createStyles({
  body: {
    display: 'flex',
    flexDirection: 'row',
  },
  externalIcon: {
    fontSize: '16px',
    marginLeft: '10px',
    verticalAlign: 'middle',
  },
  notes: {
    flex: '1 1 0',
    marginRight: '100px',
  },
  signupButton: {
    width: '120px',
  },
  teams: {
    flex: '0 0 400px',
  },
  title: {
    marginBottom: 0,
  },
  titleRow: {
    display: 'flex',
    flexDirection: 'row',
    justifyContent: 'space-between',
    marginBottom: '40px',
  },
  volleynetLink: {
    ...link,
  },
});

interface Props extends WithStyles<typeof styles> {
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
    const { classes, tournament } = this.props;

    if (!tournament) {
      return null; // todo
    }

    // const teams = tournament.teams || [];

    return (
      <Layout title={{ text: 'Tournaments', href: '/' }}>
        <div className={classes.titleRow}>
          <a
            href={tournament.link}
            className={classes.volleynetLink}
            target="_blank"
            rel="noopener noreferrer"
          >
            <Typography variant="h1">
              {tournament.name}
              <External className={classes.externalIcon} />
            </Typography>
          </a>
          <Button
            variant="contained"
            className={classes.signupButton}
            color="primary"
          >
            Signup
          </Button>
        </div>
        <div className={classes.body}>
          <div className={classes.notes}>
            {/* <div>
              <TournamentAttribute
                size="lg"
                label="Start"
                data={formatDate(tournament.start)}
              />
            </div>
            <Divider /> */}
            <Typography
              variant="body2"
              dangerouslySetInnerHTML={{ __html: tournament.htmlNotes }}
            />
          </div>
          <Card className={classes.teams}>
            <Typography variant="h2">Organiser</Typography>
            <Typography variant="body1">{tournament.organiser}</Typography>
            <Typography variant="body1">{tournament.website}</Typography>
            <Typography variant="body1">{tournament.email}</Typography>
            <Typography variant="body1">{tournament.phone}</Typography>
          </Card>
        </div>
        {/* <TeamList teams={teams} /> */}
        {/* <Link
          href={{
            pathname: '/tournament/signup',
            query: { id: tournament.id },
          }}
        >
        </Link> */}
      </Layout>
    );
  }
}

export default withStyles(styles)(withAuth(withWidth()(ShowTournament)));
