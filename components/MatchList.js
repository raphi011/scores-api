// @flow 

import React from "react";
import { withStyles } from "material-ui/styles";
import List, { ListItem, ListItemText } from "material-ui/List";
import Typography from 'material-ui/Typography';
import { formatDateTime } from '../utils/dateFormat';
import type { Match, Team } from '../types';

const styles = () => ({
  root: {
    width: "100%",
  },
  listContainer: {
    display: "flex",
    flexDirection: "row",
    alignItems: "center",
    width: "100%"
  },
  team: { flex: "1 1 0" },
  points: { fontWeight: "lighter", flex: "2 2 0" },
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
      {team.Player1.Name}
      <br />
      {team.Player2.Name}
    </span>
  );
}

function MatchList({ matches = [], onMatchClick, classes }: Props) {
  return (
    <List className={classes.root}>
      {matches.map(m => (
        <ListItem key={m.ID} button onClick={() => onMatchClick(m)}>
          <ListItemText
            primary={
              <div className={classes.listContainer}>
                <Typography className={classes.team} type="body">{getTeamName(m.Team1)}</Typography>
                <Typography className={classes.points} type="display2" align="center">
                  {m.ScoreTeam1} - {m.ScoreTeam2}
                </Typography>
                <Typography className={classes.team} type="body" align="right">{getTeamName(m.Team2)}</Typography>
              </div>
            }
            secondary={formatDateTime(new Date(m.CreatedAt))}
          />
        </ListItem>
      ))}
    </List>
  );
}

const StyledMatchList = withStyles(styles)(MatchList);

export default StyledMatchList;
