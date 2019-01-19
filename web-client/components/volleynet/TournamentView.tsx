import React from 'react';

import Link from 'next/link';
import TimeAgo from 'react-timeago';

import Button from '@material-ui/core/Button';
import Card from '@material-ui/core/Card';
import CardContent from '@material-ui/core/CardContent';
import Grid from '@material-ui/core/Grid';
import {
  createStyles,
  Theme,
  WithStyles,
  withStyles,
} from '@material-ui/core/styles';
import Typography from '@material-ui/core/Typography';
import CalendarIcon from '@material-ui/icons/DateRange';
import EmailIcon from '@material-ui/icons/Email';
import LocationIcon from '@material-ui/icons/GpsFixed';
import LinkIcon from '@material-ui/icons/Link';
import PeopleIcon from '@material-ui/icons/People';
import PhoneIcon from '@material-ui/icons/Phone';

import CenteredLoading from '../../components/CenteredLoading';
import TeamList from '../../components/volleynet/TeamList';
import { card, link, title } from '../../styles/shared';
import { isSignedup, tournamentDateString } from '../../utils/tournament';

import { TournamentInfo, User } from '../../types';

const styles = (theme: Theme) =>
  createStyles({
    card,
    container: {
      padding: theme.spacing.unit * 2,
    },
    descriptionContainer: {
      padding: theme.spacing.unit * 2,
    },
    headerContainer: {
      padding: theme.spacing.unit * 2,
    },
    infoElement: {
      fontSize: '1rem',
      verticalAlign: 'middle',
    },
    link,
    tabContent: {
      background: theme.palette.background.paper,
    },
    title: title(theme),
    updatedAt: {
      marginTop: theme.spacing.unit,
    },
  });

interface Props extends WithStyles<typeof styles> {
  tournament: TournamentInfo;
  user: User;
}

interface State {
  tabOpen: number;
}

class TournamentView extends React.Component<Props, State> {
  render() {
    const { user, tournament, classes } = this.props;

    if (!tournament) {
      return <CenteredLoading />;
    }

    const hasInfos = !!tournament.htmlNotes.trim();

    const infoText = hasInfos ? (
      <div
        className={classes.descriptionContainer}
        dangerouslySetInnerHTML={{ __html: tournament.htmlNotes }}
      />
    ) : (
      'No infos'
    );

    const infos = [
      {
        icon: <LocationIcon className={classes.infoElement} />,
        info: tournament.location,
        show: !!tournament.location,
      },
      {
        icon: <PeopleIcon className={classes.infoElement} />,
        info: `${tournament.signedupTeams} / ${tournament.maxTeams}`,
        show: tournament.maxTeams !== -1,
      },
      {
        icon: <CalendarIcon className={classes.infoElement} />,
        info: tournamentDateString(tournament),
        show: true,
      },
      {
        icon: <PhoneIcon className={classes.infoElement} />,
        info: <a href={`tel:${tournament.phone}`}>{tournament.phone}</a>,
        show: !!tournament.phone,
      },
      {
        icon: <EmailIcon className={classes.infoElement} />,
        info: <a href={`emailto:${tournament.email}`}>{tournament.email}</a>,
        show: !!tournament.email,
      },
      {
        icon: <LinkIcon className={classes.infoElement} />,
        info: (
          <a href={`//${tournament.web}`} target="_blank">
            {tournament.web}
          </a>
        ),
        show: !!tournament.web,
      },
    ].filter(detail => detail.show);

    const signedup = isSignedup(tournament, user.volleynetUserId);

    const showSignup = signedup || tournament.registrationOpen;

    const registered = tournament.teams.filter(t => !t.deregistered);
    const deregistered = tournament.teams.filter(t => t.deregistered);

    return (
      <div className={classes.container}>
        <Grid container spacing={24}>
          <Grid item xs={6}>
            <Typography className={classes.title} variant="h4">
              {tournament.name}
            </Typography>
            <Card className={classes.card}>
              <CardContent>
                <div>
                  {infos.map((info, i) => (
                    <Typography key={i} variant="subtitle1">
                      {info.icon}{' '}
                      <span className={classes.infoElement}>{info.info}</span>
                    </Typography>
                  ))}
                </div>
                {!showSignup || (
                  <Link
                    prefetch
                    href={{
                      pathname: '/volleynet/signup',
                      query: { id: tournament.id },
                    }}
                  >
                    <Button variant="contained" color="primary" fullWidth>
                      {signedup ? 'You are signed up' : 'Signup'}
                    </Button>
                  </Link>
                )}
              </CardContent>
            </Card>
            <Typography
              className={classes.updatedAt}
              variant="caption"
              align="center"
              paragraph
            >
              Last updated: <TimeAgo date={tournament.updatedAt} />
            </Typography>
            <Typography className={classes.title} variant="h4">
              Notes
            </Typography>
            <Card className={classes.card}>
              <CardContent>{infoText}</CardContent>
            </Card>
          </Grid>

          <Grid item xs={6}>
            <Typography className={classes.title} variant="h4">
              Teams
            </Typography>
            <Card>
              <CardContent>
                <TeamList
                  emptyMessage="No teams are signed up."
                  teams={registered}
                />
              </CardContent>
            </Card>
          </Grid>
        </Grid>
        {deregistered.length ? (
          <>
            <Typography className={classes.title} variant="h4">
              Deregistered
            </Typography>
            <Card className={classes.card}>
              <CardContent>
                <TeamList teams={deregistered} />
              </CardContent>
            </Card>
          </>
        ) : null}
      </div>
    );
  }
}

export default withStyles(styles)(TournamentView);
