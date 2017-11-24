import React from "react";
import { withStyles } from "material-ui/styles";
import List, { ListItem, ListItemText, ListItemIcon } from "material-ui/List";

const styles = theme => ({
  root: {
    width: "100%",
    maxWidth: 360,
    background: theme.palette.background.paper
  }
});

function MatchList({ matches = [], classes }) {
  return (
    <List className={classes.root}>
      {matches.map(m => (
        <ListItem key={m.ID} button>
          <ListItemText
            inset
            primary={
              <span>
                {m.Team1.Name}{" "}
                <b>
                  {m.ScoreTeam1} : {m.ScoreTeam2}
                </b>{" "}
                {m.Team2.Name}
              </span>
            }
          />
        </ListItem>
      ))}
    </List>
  );
}

const StyledMatchList = withStyles(styles)(MatchList);

export default StyledMatchList;