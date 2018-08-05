import React from 'react';
import Router from 'next/router';

import withAuth from '../../containers/AuthContainer';
import CenteredLoading from '../../components/CenteredLoading';
import Layout from '../../containers/LayoutContainer';
// import { loadPlayerAction } from '../../redux/actions/entities';

import { Player } from '../../types';

interface Props {
  players: Player[];
  loadPlayers: (
    filters: { gender: string },
  ) => void;
  gender: string;
  classes: any;
}

const thisYear = new Date().getFullYear().toString();

class Ranking extends React.Component<Props> {
//   static buildActions({ gender }: Props) {
//     return [
//       loadAction({
//         gender,
//       }),
//     ];
//   }

  static mapDispatchToProps = {
    // loadTournaments: loadTournamentsAction,
  };

//   static getParameters(query) {
//     let { league = 'AMATEUR TOUR' } = query;

//     if (!leagues.includes(league)) {
//       league = leagues[0];
//     }

//     return { league };
//   }

//   static mapStateToProps(state, { league }: Props) {
//     const tournaments = tournamentsByLeagueSelector(state, league);

//     return { tournaments };
//   }

//   componentDidUpdate(prevProps) {
//     const { loadTournaments, league } = this.props;

//     if (league !== prevProps.league) {
//       loadTournaments({ gender: 'M', league, season: thisYear });
//     }
//   }

  render() {
    return (
      <Layout title={{ text: 'Rankings', href: '' }}>
        TODO
      </Layout>
    );
  }
}

export default withAuth(Ranking);

