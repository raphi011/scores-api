// @flow

import React from 'react';
import Link from 'next/link';

import Tabs, { Tab } from '@material-ui/core/Tabs';
import { withStyles } from '@material-ui/core/styles';
import Card from '@material-ui/core/Card';
import Typography from '@material-ui/core/Typography';
import Button from '@material-ui/core/Button';
import PhoneIcon from '@material-ui/icons/Phone';
import PeopleIcon from '@material-ui/icons/People';
import LinkIcon from '@material-ui/icons/Link';
import CalendarIcon from '@material-ui/icons/DateRange';
import EmailIcon from '@material-ui/icons/Email';
import LocationIcon from '@material-ui/icons/GpsFixed';

import TeamList from '../../components/volleynet/TeamList';
import CenteredLoading from '../../components/CenteredLoading';

import type { FullTournament } from '../../types';

const styles = theme => ({
  headerContainer: {
    margin: theme.spacing.unit,
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

type Props = {
  tournament: FullTournament,
};

type State = {
  tabOpen: number,
};

class TournamentView extends React.Component<Props, State> {
  state = {
    tabOpen: 0,
  };

  onTabClick = (event, tabOpen) => {
    this.setState({ tabOpen });
  };

  render() {
    const { tournament, classes } = this.props;
    const { tabOpen } = this.state;
    if (!tournament) {
      return <CenteredLoading />;
    }
    const body = { __html: tournament.htmlNotes };

    const infos = [
      {
        icon: <LocationIcon className={classes.infoElement} />,
        info: tournament.location,
      },
      {
        icon: <PeopleIcon className={classes.infoElement} />,
        info: `${tournament.teams.length} / ${tournament.maxTeams}`,
      },
      {
        icon: <CalendarIcon className={classes.infoElement} />,
        info: tournament.startDate,
      },
      {
        icon: <PhoneIcon className={classes.infoElement} />,
        info: tournament.phone,
      },
      {
        icon: <EmailIcon className={classes.infoElement} />,
        info: <a href={`emailto:${tournament.email}`}>{tournament.email}</a>,
      },
      {
        icon: <LinkIcon className={classes.infoElement} />,
        info: <a href={tournament.web}>{tournament.web}</a>,
      },
    ];

    return (
      <div>
        <Card className={classes.headerContainer}>
          <Typography variant="title">{tournament.name}</Typography>
          <div className={classes.infoContainer}>
            {infos.map((info, i) => (
              <Typography key={i} variant="subheading">
                {info.icon}{' '}
                <span className={classes.infoElement}>{info.info}</span>
              </Typography>
            ))}
          </div>
          <Link
            prefetch
            href={{
              pathname: '/volleynet/signup',
              query: { id: tournament.id },
            }}
          >
            <Button variant="raised" color="primary">
              Signup
            </Button>
          </Link>
        </Card>

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
          {tabOpen === 0 ? (
            <div
              className={classes.descriptionContainer}
              dangerouslySetInnerHTML={body}
            />
          ) : (
            <TeamList teams={tournament.teams} />
          )}
        </div>
      </div>
    );
  }
}

export default withStyles(styles)(TournamentView);
