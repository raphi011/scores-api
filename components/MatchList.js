// @flow

import React from 'react';
import { withStyles } from 'material-ui/styles';
import List, { ListItem, ListItemText } from 'material-ui/List';
import Typography from 'material-ui/Typography';
import { formatDateTime } from '../utils/dateFormat';
import type { Match, Team } from '../types';

const styles = () => ({
  root: {
    width: '100%',
  },
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
  matches: Array<Match>,
  onMatchClick: Match => void,
  classes: Object,
};

function getTeamName(team: Team) {
  if (team.Name) return team.Name;

  return (
    <span>
      {team.player1.name}
      <br />
      {team.player2.name}
    </span>
  );
}

function MatchList({ matches = [], onMatchClick, classes }: Props) {
  return (
    <List className={classes.root}>
      {matches.map(m => (
        <ListItem key={m.id} button onClick={() => onMatchClick(m)}>
          <ListItemText
            primary={
              <div className={classes.listContainer}>
                <Typography className={classes.team} type="body1">
                  {getTeamName(m.team1)}
                </Typography>
                <Typography
                  className={classes.points}
                  type="display2"
                  align="center"
                >
                  {m.scoreTeam1} - {m.scoreTeam2}
                </Typography>
                <Typography className={classes.team} type="body1" align="right">
                  {getTeamName(m.team2)}
                </Typography>
              </div>
            }
            secondary={formatDateTime(new Date(m.createdAt))}
          />
        </ListItem>
      ))}
    </List>
  );
}

const StyledMatchList = withStyles(styles)(MatchList);

export default StyledMatchList;
