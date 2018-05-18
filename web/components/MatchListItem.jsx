// @flow
import React from 'react';
import { withStyles } from '@material-ui/core/styles';
import { ListItem, ListItemText } from '@material-ui/core/List';
import Typography from '@material-ui/core/Typography';

import TeamName from './TeamName';

import type { Match, Team } from '../types';

const itemStyles = () => ({
  listContainer: {
    display: 'flex',
    flexDirection: 'row',
    alignItems: 'center',
    width: '100%',
  },
  team: { flex: '1 1 0' },
  points: { fontWeight: 'lighter', flex: '2 2 0' },
});

type Props = {
  match: Match,
  onMatchClick: Match => void,
  highlightPlayerId: number,
  classes: Object,
};

function WinnerAndLoser(
  match: Match,
): {
  winner: Team,
  loser: Team,
  winnerScore: number,
  loserScore: number,
} {
  if (match.scoreTeam1 > match.scoreTeam2) {
    return {
      winner: match.team1,
      loser: match.team2,
      winnerScore: match.scoreTeam1,
      loserScore: match.scoreTeam2,
    };
  }

  return {
    winner: match.team2,
    loser: match.team1,
    winnerScore: match.scoreTeam2,
    loserScore: match.scoreTeam1,
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
            <Typography className={classes.team} variant="body1">
              <TeamName
                team={result.winner}
                highlightPlayerId={highlightPlayerId}
              />
            </Typography>
            <Typography
              className={classes.points}
              variant="display2"
              align="center"
            >
              {score}
            </Typography>
            <Typography className={classes.team} variant="body1" align="right">
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

export default withStyles(itemStyles)(MatchListItem);
