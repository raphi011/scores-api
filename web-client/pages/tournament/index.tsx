import React from 'react';

import Router from 'next/router';
import TimeAgo from 'react-timeago';

import Button from '@material-ui/core/Button';
import {
  createStyles,
  withStyles,
  WithStyles,
  Theme,
} from '@material-ui/core/styles';
import Typography from '@material-ui/core/Typography';
import withWidth from '@material-ui/core/withWidth';
import External from '@material-ui/icons/OpenInNew';
import Tabs from '@material-ui/core/Tabs';
import Tab from '@material-ui/core/Tab';

import withAuth from '../../containers/AuthContainer';
import Layout from '../../containers/LayoutContainer';

import { userSelector } from '../../redux/auth/selectors';
import { loadTournamentAction } from '../../redux/entities/actions';
import { tournamentSelector } from '../../redux/entities/selectors';
import { link } from '../../styles/shared';
import { Tournament, User } from '../../types';
import { Store } from '../../redux/store';
import TeamList from '../../components/volleynet/TeamList';
import * as Query from '../../utils/query';
import { fontPalette } from '../../styles/theme';
import classNames from 'classnames';
import { Breakpoint } from '@material-ui/core/styles/createBreakpoints';
import { formatDate } from '../../utils/date';
import { QueryStringMapObject } from 'next';

const styles = (theme: Theme) =>
  createStyles({
    attr: {
      fontSize: fontPalette[800],
      fontWeight: 500,
      marginRight: '10px',
    },
    attrValue: {
      fontSize: fontPalette[500],
      fontWeight: 300,
      color: theme.palette.grey[500],
      marginRight: '40px',
    },
    body: {
      marginTop: '30px',
    },
    externalIcon: {
      fontSize: '16px',
      marginLeft: '10px',
      verticalAlign: 'middle',
    },
    signupButton: {
      width: '120px',
    },
    tabs: {
      marginTop: '50px',
    },
    title: {
      marginBottom: 0,
    },
    titleRow: {
      marginBottom: '20px',
      display: 'flex',
    },
    titleRowDesktop: {
      alignItems: 'flex-start',
      flexDirection: 'row',
      justifyContent: 'space-between',
    },
    titleRowMobile: {
      flexDirection: 'column',
      alignItems: 'stretch',
    },
    volleynetLink: {
      ...link,
    },
  });

type TabOption = 'notes' | 'teams' | 'organiser';

const tabOptions: TabOption[] = ['notes', 'teams', 'organiser'];

interface Props extends WithStyles<typeof styles> {
  tournament?: Tournament;
  tournamentId: string;
  user: User;
  tab: TabOption;
  width: Breakpoint;
}

class ShowTournament extends React.Component<Props> {
  static getParameters(query: QueryStringMapObject): Partial<Props> {
    const tab = Query.oneOfDefault(query, 'tab', tabOptions, 'notes');
    const tournamentId = Query.str(query, 'id');

    return { tournamentId, tab };
  }

  static buildActions({ tournamentId }: Props) {
    return [loadTournamentAction(tournamentId)];
  }

  static mapStateToProps(state: Store, { tournamentId }: Props) {
    const tournament = tournamentSelector(state, tournamentId);
    const { user } = userSelector(state);

    return { tournament, user };
  }

  onSelectTab = (_event: React.ChangeEvent<{}>, tabIndex: number) => {
    const { tournamentId: id } = this.props;

    const tab = tabOptions[tabIndex];

    Router.push({
      pathname: '/tournament',
      query: { tab, id },
    });
  };

  render() {
    const { classes, tab, tournament, width } = this.props;

    if (!tournament) {
      return null;
    }

    const teams = tournament.teams || [];
    const isMobile = ['xs', 'sm'].includes(width);

    const titleRowClassName = classNames(classes.titleRow, {
      [classes.titleRowDesktop]: !isMobile,
      [classes.titleRowMobile]: isMobile,
    });

    return (
      <Layout title={{ text: 'Tournaments', href: '/' }}>
        <div className={titleRowClassName}>
          <div>
            <a
              href={tournament.link}
              className={classes.volleynetLink}
              target="_blank"
              rel="noopener noreferrer"
            >
              <Typography variant="h1" inline>
                {tournament.name}
                <External className={classes.externalIcon} />
              </Typography>
            </a>
            <Typography variant="subtitle1">
              {tournament.subLeague} - {formatDate(tournament.start)}
            </Typography>
          </div>
          <Button
            variant="contained"
            className={classes.signupButton}
            color="primary"
          >
            Signup
          </Button>
        </div>
        <div>
          <TimeAgo
            date={tournament.start}
            formatter={(value, unit, suffix) => (
              <>
                <Typography inline className={classes.attr}>
                  {value}
                </Typography>
                <Typography inline className={classes.attrValue}>
                  {unit}
                  {value > 1 ? 's' : ''} {suffix}
                </Typography>
              </>
            )}
          />
          <Typography inline className={classes.attr}>
            {`${tournament.signedupTeams}/${tournament.maxTeams}`}
          </Typography>
          <Typography inline className={classes.attrValue}>
            Teams signed up
          </Typography>
          <Typography inline className={classes.attr}>
            {tournament.maxPoints}
          </Typography>
          <Typography inline className={classes.attrValue}>
            Max points
          </Typography>
        </div>
        <Tabs
          className={classes.tabs}
          indicatorColor="primary"
          value={tabOptions.indexOf(tab)}
          variant={isMobile ? 'fullWidth' : 'standard'}
          onChange={this.onSelectTab}
        >
          {tabOptions.map(t => (
            <Tab key={t} label={t} />
          ))}
        </Tabs>
        <div className={classes.body}>
          {tab === 'notes' && (
            <Typography
              variant="body2"
              dangerouslySetInnerHTML={{ __html: tournament.htmlNotes }}
            />
          )}
          {tab === 'teams' && <TeamList teams={teams} />}
        </div>
      </Layout>
    );
  }
}

export default withStyles(styles)(withAuth(withWidth()(ShowTournament)));
