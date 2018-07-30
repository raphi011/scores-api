import React from 'react';
import Link from 'next/link';

import { withStyles, Theme, createStyles } from '@material-ui/core/styles';
import Tabs from '@material-ui/core/Tabs';
import Tab from '@material-ui/core/Tab';
import Card from '@material-ui/core/Card';
import Typography from '@material-ui/core/Typography';
import Button from '@material-ui/core/Button';
import PhoneIcon from '@material-ui/icons/Phone';
import PeopleIcon from '@material-ui/icons/People';
import LinkIcon from '@material-ui/icons/Link';
import CalendarIcon from '@material-ui/icons/DateRange';
import EmailIcon from '@material-ui/icons/Email';
import LocationIcon from '@material-ui/icons/GpsFixed';
import { tournamentDateString, isSignedup } from '../../utils/tournament';
import { formatDateTime } from '../../utils/dateFormat';

import TeamList from '../../components/volleynet/TeamList';
import CenteredLoading from '../../components/CenteredLoading';

import { Tournament, User } from '../../types';

const styles = (theme: Theme) =>
  createStyles({
    link: {
      textDecoration: 'none',
    },
    title: {
      marginTop: theme.spacing.unit,
      textDecoration: 'none',
      marginBottom: theme.spacing.unit * 2,
      wrapWrap: 'break-word',
    },
    updatedAt: {
      marginTop: theme.spacing.unit,
    },
    header: {
      padding: theme.spacing.unit * 2,
    },
    headerContainer: {
      padding: theme.spacing.unit * 2,
    },
    infoContainer: {
      margin: '20px 0',
    },
    tabContent: {
      background: theme.palette.background.paper,
    },
    descriptionContainer: {
      padding: theme.spacing.unit * 2,
    },
    infoElement: {
      verticalAlign: 'middle',
      fontSize: '1rem',
    },
  });

interface Props {
  tournament: Tournament;
  user: User;
  classes: any;
}

interface State {
  tabOpen: number;
}

class TournamentView extends React.Component<Props, State> {
  state = {
    tabOpen: 0,
  };

  onTabClick = (_, tabOpen) => {
    this.setState({ tabOpen });
  };

  render() {
    const { user, tournament, classes } = this.props;
    const { tabOpen } = this.state;
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
        info: `${(tournament.teams || []).length} / ${tournament.maxTeams}`,
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

    return (
      <div>
        <div className={classes.headerContainer}>
          <div className={classes.title}>
            <a
              className={classes.link}
              href={`${tournament.link}`}
              target="_blank"
            >
              <Typography variant="display1">{tournament.name}</Typography>
            </a>
          </div>
          <Card className={classes.header}>
            <div className={classes.infoContainer}>
              {infos.map((info, i) => (
                <Typography key={i} variant="subheading">
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
                <Button variant="raised" color="primary" fullWidth>
                  {signedup ? 'You are signed up' : 'Signup'}
                </Button>
              </Link>
            )}
          </Card>
          <Typography
            className={classes.updatedAt}
            variant="caption"
            align="center"
            paragraph
          >
            Last updated: {formatDateTime(tournament.updatedAt)}
          </Typography>
        </div>

        <Tabs
          onChange={this.onTabClick}
          value={tabOpen}
          textColor="primary"
          fullWidth
        >
          RegisterI
          <Tab label="Infos" />
          <Tab label="Teams" />
        </Tabs>
        <div className={classes.tabContent}>
          {tabOpen === 0 ? infoText : <TeamList teams={tournament.teams} />}
        </div>
      </div>
    );
  }
}

export default withStyles(styles)(TournamentView);
