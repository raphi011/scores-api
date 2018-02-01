// @flow

import React from 'react';
import { withStyles } from 'material-ui/styles';
import List, { ListItem, ListItemText, ListItemIcon } from 'material-ui/List';
import Badge from 'material-ui/Badge';
import PersonIcon from 'material-ui-icons/Person';
import type { Player } from '../types';

const styles = theme => ({
  root: {
    width: '100%',
    background: theme.palette.background.paper,
  },
});

const playerItemStyles = theme => ({
  default: {
    background: theme.palette.background.paper,
  },
  team1: {
    background: 'red',
  },
  team2: {
    background: 'green',
  },
});

type Props = {
  players: Array<Player>,
  onUnsetPlayer: number => void,
  onSetPlayer: (number, number, boolean) => void,
  player1Id: number,
  player2Id: number,
  player3Id: number,
  player4Id: number,
  classes: Object,
};

class SelectPlayers extends React.Component<Props> {
  onSelectPlayer = id => {
    const { onUnsetPlayer, onSetPlayer } = this.props;

    let unassigned;
    let assignedCount = 0;
    let selected;

    for (let i = 1; i < 5; i += 1) {
      const pId = this.props[`player${i}Id`];

      if (pId === id) {
        selected = i;
        break;
      } else if (!unassigned && pId === 0) {
        unassigned = i;
      }

      if (pId) assignedCount += 1;
    }

    if (selected) {
      onUnsetPlayer(selected);
    } else if (unassigned) {
      onSetPlayer(unassigned, id, assignedCount === 3);
    }
  };

  playerNr = playerId => {
    const { player1Id, player2Id, player3Id, player4Id } = this.props;

    if (playerId === player1Id) return 1;
    else if (playerId === player2Id) return 2;
    else if (playerId === player3Id) return 3;
    else if (playerId === player4Id) return 4;

    return 0;
  };

  render() {
    const { players = [], classes } = this.props;

    return (
      <List className={classes.root}>
        {players.map(p => (
          <StyledPlayerListItem
            onClick={() => this.onSelectPlayer(p.id)}
            key={p.id}
            player={p}
            playerNr={this.playerNr(p.id)}
          />
        ))}
      </List>
    );
  }
}

type PlayerListProps = {
  onClick: Event => void,
  player: Player,
  playerNr: number,
};

function PlayerListItem({ player, onClick, playerNr }: PlayerListProps) {
  let color;
  let team;

  switch (playerNr) {
    case 1:
    case 2:
      color = 'primary';
      team = 1;
      break;
    case 3:
    case 4:
      color = 'secondary';
      team = 2;
      break;
    default:
      color = '';
      team = null;
  }

  return (
    <ListItem onClick={onClick} button>
      {playerNr ? (
        <ListItemIcon>
          <Badge badgeContent={team} color={color}>
            <PersonIcon />
          </Badge>
        </ListItemIcon>
      ) : null}
      <ListItemText inset primary={player.name} />
    </ListItem>
  );
}

const StyledSelectPlayers = withStyles(styles)(SelectPlayers);
const StyledPlayerListItem = withStyles(playerItemStyles)(PlayerListItem);

export default StyledSelectPlayers;
