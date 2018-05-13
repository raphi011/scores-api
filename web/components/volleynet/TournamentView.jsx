// @flow

import React from 'react';
import Link from 'next/link';

import { withStyles } from 'material-ui/styles';
import Typography from 'material-ui/Typography';
import Button from 'material-ui/Button';
import TeamList from '../../components/volleynet/TeamList';

import type { FullTournament } from '../../types';

const styles = () => ({});

type Props = {
  tournament: FullTournament,
};

function TournamentView({ tournament }: Props) {
  if (!tournament) {
    return null;
  }

  //   const body = { __html: tournament.htmlNotes };

  return (
    <div>
      <Typography variant="headline">{tournament.name}</Typography>
      <Typography variant="subheading">
        {tournament.startDate} - {tournament.endDate}
      </Typography>
      {/* <div dangerouslySetInnerHTML={body} /> */}
      <TeamList teams={tournament.teams} />

      <Link
        prefetch
        href={{ pathname: '/volleynet/signup', query: { id: tournament.id } }}
      >
        <Button variant="raised" color="primary">
          Signup
        </Button>
      </Link>
    </div>
  );
}

export default withStyles(styles)(TournamentView);
