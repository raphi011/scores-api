import ListItem from '@material-ui/core/ListItem';
import ListItemText from '@material-ui/core/ListItemText';
import { createStyles, withStyles } from '@material-ui/core/styles';
import React from 'react';

import { Tournament } from '../../types';

const styles = createStyles({
  item: {
    borderBottom: '1px solid rgba(0, 0, 0, 0.075)',
  },
});

interface IProps {
  tournament: Tournament;
  onClick: (Tournament) => void;
  classes: any;
}

function buildSubtitle(tournament: Tournament) {
  let st = tournament.signedupTeams.toString();

  if (tournament.maxTeams >= 0) {
    st += ` / ${tournament.maxTeams} teams`;
  }

  st += ` • ${tournament.league}`;

  return st;
}

const TournamentListItem = ({ tournament, classes, onClick }: IProps) => {
  let primary: string | JSX.Element = tournament.name;

  if (tournament.status === 'canceled') {
    primary = (
      <span>
        <span style={{ textDecoration: 'line-through' }}>{primary}</span>{' '}
        (canceled)
      </span>
    );
  }

  return (
    <ListItem
      className={classes.item}
      button
      onClick={() => onClick(tournament)}
    >
      <ListItemText primary={primary} secondary={buildSubtitle(tournament)} />
    </ListItem>
  );
};

export default withStyles(styles)(TournamentListItem);
