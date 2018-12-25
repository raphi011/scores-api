import React from 'react';

import ListItem from '@material-ui/core/ListItem';
import ListItemText from '@material-ui/core/ListItemText';
import { createStyles, WithStyles, withStyles } from '@material-ui/core/styles';
import Typography from '@material-ui/core/Typography';

import TeamName from './TeamName';

import { Match, Team } from '../types';

const styles = createStyles({
  listContainer: {
    alignItems: 'center',
    display: 'flex',
    flexDirection: 'row',
    width: '100%',
  },
  points: { fontWeight: 'lighter', flex: '2 2 0' },
  team: { flex: '1 1 0' },
});

interface Props extends WithStyles<typeof styles> {
  match: Match;
  highlightPlayerId: number;

  onMatchClick: (match: Match) => void;
}

function WinnerAndLoser(
  match: Match,
): {
  winner: Team;
  loser: Team;
  winnerScore: number;
  loserScore: number;
} {
  if (match.scoreTeam1 > match.scoreTeam2) {
    return {
      loser: match.team2,
      loserScore: match.scoreTeam2,
      winner: match.team1,
      winnerScore: match.scoreTeam1,
    };
  }

  return {
    loser: match.team1,
    loserScore: match.scoreTeam1,
    winner: match.team2,
    winnerScore: match.scoreTeam2,
  };
}

const MatchListItem = ({
  onMatchClick,
  match,
  highlightPlayerId,
  classes,
}: Props) => {
  const result = WinnerAndLoser(match);

  const winnerScore = result.winnerScore.toString(); /*.padStart(2, '0');*/
  const loserScore = result.loserScore.toString(); /*.padStart(2, '0'); */

  const score = `${winnerScore} - ${loserScore}`;

  return (
    <ListItem button onClick={() => onMatchClick(match)}>
      <ListItemText
        primary={
          <div className={classes.listContainer}>
            <Typography className={classes.team} variant="body2">
              <TeamName
                team={result.winner}
                highlightPlayerId={highlightPlayerId}
              />
            </Typography>
            <Typography className={classes.points} variant="h3" align="center">
              {score}
            </Typography>
            <Typography className={classes.team} variant="body2" align="right">
              <TeamName
                team={result.loser}
                highlightPlayerId={highlightPlayerId}
              />
            </Typography>
          </div>
        }
      />
    </ListItem>
  );
};

export default withStyles(styles)(MatchListItem);
