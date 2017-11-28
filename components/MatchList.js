import React from "react";
import { withStyles } from "material-ui/styles";
import List, { ListItem, ListItemText, ListItemIcon } from "material-ui/List";

const styles = theme => ({
  root: {
    width: "100%",
    maxWidth: 360,
    background: theme.palette.background.paper
  },
  listContainer: {
    display: 'flex',
    flexDirection: 'row',
    justifyContent: 'space-between',
    alignItems: 'center',
    width: '100%',
  },
  team1: {},
  team2: { textAlign: 'right' },
  points: {
    fontSize: "35px"
  }
});

function getTeamName(team) {
  if (team.Name) return team.Name;

  return (
    <span>
      {team.Player1.Name}
      <br />
      {team.Player2.Name}
    </span>
  );
}

function MatchList({ matches = [], onMatchClick, classes }) {
  return (
    <List className={classes.root}>
      {matches.map(m => (
        <ListItem key={m.ID} button onClick={() => onMatchClick(m)}>
          <ListItemText
            inset
            primary={
              <div className={classes.listContainer}>
                <div>{getTeamName(m.Team1)} </div>
                <div className={classes.points}>
                  {m.ScoreTeam1} : {m.ScoreTeam2}
                </div>
                <div className={classes.team2}>{getTeamName(m.Team2)}</div>
              </div>
            }
            secondary={new Date(m.CreatedAt).toUTCString()}
          />
        </ListItem>
      ))}
    </List>
  );
}

const StyledMatchList = withStyles(styles)(MatchList);

export default StyledMatchList;
