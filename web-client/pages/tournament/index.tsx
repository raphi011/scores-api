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
import SignupIcon from '@material-ui/icons/AssignmentTurnedIn';
import Tabs from '@material-ui/core/Tabs';
import Tab from '@material-ui/core/Tab';
import Fab from '@material-ui/core/Fab';
import { Breakpoint } from '@material-ui/core/styles/createBreakpoints';

import withAuth from '../../hoc/next/withAuth';
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
import { formatDate } from '../../utils/date';
import withConnect, { Context } from '../../hoc/next/withConnect';

const styles = (theme: Theme) =>
  createStyles({
    attrContainer: {
      marginRight: '40px',
    },
    attr: {
      fontSize: fontPalette[800],
      fontWeight: 500,
      marginRight: '10px',
    },
    attrValue: {
      fontSize: fontPalette[500],
      fontWeight: 300,
      color: theme.palette.grey[500],
    },
    body: {
      marginTop: '30px',
    },
    fab: {
      position: 'fixed',
      bottom: theme.spacing.unit * 2,
      right: theme.spacing.unit * 2,
    },
    externalIcon: {
      fontSize: '16px',
      marginLeft: '10px',
      verticalAlign: 'middle',
    },
    signupButton: {
      width: '120px',
    },
    stats: {
      display: 'flex',
      flexDirection: 'row',
      [theme.breakpoints.down('sm')]: {
        flexDirection: 'column',
      },
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

type TabOption = 'notes' | 'teams';

const tabOptions: TabOption[] = ['notes', 'teams'];

interface Props extends WithStyles<typeof styles> {
  tournament?: Tournament;
  tournamentId: string;
  user: User;
  tab: TabOption;
  width: Breakpoint;
}

interface State {
  signupPageOpen: boolean;
}

class ShowTournament extends React.Component<Props, State> {
  state = {
    signupPageOpen: false,
  };

  static async getInitialProps(ctx: Context): Promise<Partial<Props>> {
    const { query } = ctx;

    const tab = Query.oneOfDefault(query, 'tab', tabOptions, 'notes');
    const tournamentId = Query.one(query, 'id');

    return { tournamentId, tab };
  }

  static buildActions({ tournamentId }: Props) {
    return [loadTournamentAction(tournamentId)];
  }

  static mapStateToProps(state: Store, { tournamentId }: Props) {
    const tournament = tournamentSelector(state, tournamentId);
    const user = userSelector(state);

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

  renderBody = () => {
    const { classes, tab, tournament, width } = this.props;
    // const { signupPageOpen } = this.state;

    const isMobile = ['xs', 'sm'].includes(width);

    if (!tournament) {
      return null;
    }

    const teams = tournament.teams || [];

    // if (signupPageOpen) {
    //   return (
    //     <Signup
    //   );
    // }

    return (
      <>
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
      </>
    );
  };

  render() {
    const { classes, tournament, width } = this.props;

    if (!tournament) {
      return null;
    }

    const isMobile = ['xs', 'sm'].includes(width);

    const titleRowClassName = classNames(classes.titleRow, {
      [classes.titleRowDesktop]: !isMobile,
      [classes.titleRowMobile]: isMobile,
    });

    let button = null;
    let fabButton = null;

    if (isMobile) {
      fabButton = (
        <>
          <div style={{ height: '30px' }} />
          <Fab className={classes.fab} color="primary">
            <SignupIcon />
          </Fab>
        </>
      );
    } else {
      button = (
        <Button
          variant="contained"
          className={classes.signupButton}
          color="primary"
        >
          Signup
        </Button>
      );
    }

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
          {button}
        </div>
        <div>
          <div className={classes.stats}>
            <span className={classes.attrContainer}>
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
            </span>
            <span className={classes.attrContainer}>
              <Typography inline className={classes.attr}>
                {`${tournament.signedupTeams}/${tournament.maxTeams}`}
              </Typography>
              <Typography inline className={classes.attrValue}>
                Teams signed up
              </Typography>
            </span>
            <span className={classes.attrContainer}>
              <Typography inline className={classes.attr}>
                {tournament.maxPoints}
              </Typography>
              <Typography inline className={classes.attrValue}>
                Max points
              </Typography>
            </span>
          </div>
        </div>
        {this.renderBody()}
        {fabButton}
      </Layout>
    );
  }
}

export default withAuth(
  withConnect(withWidth()(withStyles(styles)(ShowTournament))),
);
