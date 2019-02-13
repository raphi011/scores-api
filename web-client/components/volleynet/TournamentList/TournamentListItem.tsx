import React from 'react';

import ListItem from '@material-ui/core/ListItem';
import {
  createStyles,
  Theme,
  WithStyles,
  withStyles,
} from '@material-ui/core/styles';
import Typography from '@material-ui/core/Typography';
import ArrowRight from '@material-ui/icons/KeyboardArrowRight';

import StatusMarker from './StatusMarker';

import { Tournament } from '../../../types';
import TournamentAttribute from './TournamentAttribute';

const styles = (theme: Theme) =>
  createStyles({
    arrow: {
      alignSelf: 'center',
      paddingRight: '5px',
    },
    attributes: {
      '&> *': {
        flex: '1',
      },
      display: 'flex',
      flexDirection: 'row',
    },
    content: {
      display: 'flex',
      flex: '1 0 auto',
      flexDirection: 'column',
      margin: '5px 5px 5px 20px',
      padding: '5px',
      width: '200px',
    },
    gender: {
      color: theme.palette.grey[400],
    },
    item: {
      alignItems: 'stretch',
      flex: '1 0 auto',
      padding: 0,
    },
    name: {
      color: theme.palette.grey[800],
      fontSize: '20px',
      marginBottom: '15px',
      overflow: 'hidden',
      textOverflow: 'ellipsis',
      whiteSpace: 'nowrap',
    },
  });

interface Props extends WithStyles<typeof styles> {
  tournament: Tournament;

  onClick: (tournament: Tournament) => void;
}

const TournamentListItem = ({ tournament, classes, onClick }: Props) => {
  return (
    <ListItem
      className={classes.item}
      button
      divider
      onClick={() => onClick(tournament)}
    >
      <StatusMarker status={tournament.status} />
      <div className={classes.content}>
        <Typography className={classes.name}>
          {tournament.name}{' '}
          <span className={classes.gender}>
            ({tournament.gender.toLowerCase()})
          </span>
        </Typography>
        <div className={classes.attributes}>
          <TournamentAttribute
            label="Teams"
            data={`${tournament.signedupTeams}/${tournament.maxTeams}`}
          />
          <TournamentAttribute label="Status" data={tournament.status} />
          <TournamentAttribute label="League" data={tournament.league} />
        </div>
      </div>
      <ArrowRight fontSize="large" className={classes.arrow} />
    </ListItem>
  );
};

export default withStyles(styles)(TournamentListItem);
