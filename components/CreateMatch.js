import React from "react";
import { withStyles } from "material-ui/styles";
import List, { ListItem, ListItemText, ListItemIcon } from "material-ui/List";
import Badge from "material-ui/Badge";
import PersonIcon from "material-ui-icons/Person";

const styles = theme => ({
  root: {
    width: "100%",
    background: theme.palette.background.paper
  }
});

const playerItemStyles = theme => ({
  default: {
    background: theme.palette.background.paper
  },
  team1: {
    background: "red"
  },
  team2: {
    background: "green"
  }
});

class CreateMatch extends React.Component {
  playerNr = playerID => {
    const { player1ID, player2ID, player3ID, player4ID } = this.props;

    if (playerID === player1ID) return 1;
    else if (playerID === player2ID) return 2;
    else if (playerID === player3ID) return 3;
    else if (playerID === player4ID) return 4;

    return 0;
  };

  onSelectPlayer = ID => {
    const { onUnsetPlayer, onSetPlayer } = this.props;

    let unassigned;
    let assignedCount = 0;
    let selected;

    for (let i = 1; i < 5; i++) {
      const pID = this.props[`player${i}ID`];

      if (pID === ID) {
        selected = i;
        break;
      } else if (!unassigned && pID === 0) {
        unassigned = i;
      }

      if (pID) assignedCount++;
    }

    if (selected) {
      onUnsetPlayer(selected);
    } else if (unassigned) {
      onSetPlayer(unassigned, ID, assignedCount === 3);
    }
  };

  render() {
    const {
      players = [],
      classes,
      player1,
      player2,
      player3,
      player4
    } = this.props;

    return (
      <List className={classes.root}>
        {players.map(p => (
          <StyledPlayerListItem
            onClick={() => this.onSelectPlayer(p.ID)}
            key={p.ID}
            player={p}
            playerNr={this.playerNr(p.ID)}
          />
        ))}
      </List>
    );
  }
}

function PlayerListItem({ player, onClick, playerNr, classes }) {
  let color;

  switch (playerNr) {
    case 1:
    case 2:
      color = "primary";
      break;
    case 3:
    case 4:
      color = "accent";
      break;
  }

  return (
    <ListItem onClick={onClick} button>
      {playerNr ? (
        <ListItemIcon>
          <Badge badgeContent={playerNr} color={color}>
            <PersonIcon />
          </Badge>
        </ListItemIcon>
      ) : null}
      <ListItemText inset primary={player.Name} />
    </ListItem>
  );
}

const StyledCreateMatch = withStyles(styles)(CreateMatch);
const StyledPlayerListItem = withStyles(playerItemStyles)(PlayerListItem);

export default StyledCreateMatch;
