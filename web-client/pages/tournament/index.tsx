import React from 'react';

import Router from 'next/router';

import { createStyles, withStyles, WithStyles } from '@material-ui/core/styles';
import Typography from '@material-ui/core/Typography';
import withWidth from '@material-ui/core/withWidth';
import Tab from '@material-ui/core/Tab';
import Tabs from '@material-ui/core/Tabs';
import { Breakpoint } from '@material-ui/core/styles/createBreakpoints';

import withAuth from '../../hoc/next/withAuth';
import Layout from '../../containers/LayoutContainer';

import { userSelector } from '../../redux/auth/selectors';
import { loadTournamentAction } from '../../redux/entities/actions';
import { tournamentSelector } from '../../redux/entities/selectors';
import { Tournament, User } from '../../types';
import { Store } from '../../redux/store';
import TeamList from '../../components/volleynet/TeamList';
import * as Query from '../../utils/query';
import withConnect, { Context } from '../../hoc/next/withConnect';
import TournamentHeader from '../../components/volleynet/TournamentHeader';

const styles = createStyles({
  body: {
    marginTop: '30px',
  },
  tabs: {
    marginTop: '50px',
  },
  title: {
    marginBottom: 0,
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

class ShowTournament extends React.Component<Props> {
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

  onSignup = () => {
    const { tournamentId: id } = this.props;

    Router.push({
      pathname: '/tournament/signup',
      query: { id },
    });
  };

  renderBody = () => {
    const { classes, tab, tournament, width } = this.props;

    const isMobile = ['xs', 'sm'].includes(width);

    if (!tournament) {
      return null;
    }

    const teams = tournament.teams || [];

    return (
      <>
        <TournamentHeader tournament={tournament} onSignup={this.onSignup} />
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
