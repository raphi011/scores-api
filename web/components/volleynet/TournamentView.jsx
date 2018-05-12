// @flow

import React from 'react';
import { withStyles } from 'material-ui/styles';
import Typography from 'material-ui/Typography';
// import Tooltip from 'material-ui/Tooltip';
// import Avatar from 'material-ui/Avatar';
import type { FullTournament } from '../types';

const styles = () => ({});

type Props = {
  tournament: FullTournament,
  //   classes: Object,
};

function TournamentView({ tournament /*, classes */ }: Props) {
  if (!tournament) {
    return null;
  }

  //   const body = { __html: tournament.htmlNotes };

  return (
    <div>
      <Typography variant="headline">{tournament.name}</Typography>
      <Typography variant="subheading">{tournament.startDate}</Typography>
      <Typography variant="subheading">{tournament.startDate}</Typography>
      {/* <div dangerouslySetInnerHTML={body} /> */}
    </div>
  );
}

export default withStyles(styles)(TournamentView);
