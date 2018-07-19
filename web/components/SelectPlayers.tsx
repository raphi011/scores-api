import React from 'react';
import Avatar from '@material-ui/core/Avatar';
import { withStyles, Theme, createStyles } from '@material-ui/core/styles';
import List from '@material-ui/core/List';
import ListItem from '@material-ui/core/ListItem';
import ListItemText from '@material-ui/core/ListItemText';
import Badge from '@material-ui/core/Badge';

import { Player, Classes } from '../types';

const styles = (theme: Theme) =>
  createStyles({
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

interface Props {
  players: Player[];
  onUnsetPlayer: (number) => void;
  onSetPlayer: (number, Player, boolean) => void;
  player1: Player;
  player2: Player;
  player3: Player;
  player4: Player;
  classes: Classes;
}

class SelectPlayers extends React.Component<Props> {
  onSelectPlayer = (selected: Player) => {
    const { onUnsetPlayer, onSetPlayer } = this.props;

    let unassigned = 0;
    let assignedCount = 0;
    let selectedNr = 0;

    for (let i = 1; i < 5; i += 1) {
      const player: Player = this.props[`player${i}`];
      const pId = player ? player.id : 0;

      if (pId === selected.id) {
        selectedNr = i;
        break;
      } else if (!unassigned && pId === 0) {
        unassigned = i;
      }

      if (pId) assignedCount += 1;
    }

    if (selectedNr) {
      onUnsetPlayer(selectedNr);
    } else if (unassigned) {
      onSetPlayer(unassigned, selected, assignedCount === 3);
    }
  };

  isSamePlayer = (p1, p2) => {
    if (!p1 || !p2) return false;

    return p1.id === p2.id;
  };

  playerNr = (player: Player): number => {
    const { player1, player2, player3, player4 } = this.props;

    if (this.isSamePlayer(player, player1)) return 1;
    else if (this.isSamePlayer(player, player2)) return 2;
    else if (this.isSamePlayer(player, player3)) return 3;
    else if (this.isSamePlayer(player, player4)) return 4;

    return 0;
  };

  render() {
    const { players = [], classes } = this.props;

    return (
      <List className={classes.root}>
        {players.map(p => (
          <StyledPlayerListItem
            onClick={() => this.onSelectPlayer(p)}
            key={p.id}
            player={p}
            playerNr={this.playerNr(p)}
          />
        ))}
      </List>
    );
  }
}

interface PlayerListProps {
  onClick: (Event) => void;
  player: Player;
  playerNr: number;
}

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

  let avatar = player.profileImageUrl ? (
    <Avatar src={player.profileImageUrl} />
  ) : (
    <Avatar>{player.name.substring(0, 1)}</Avatar>
  );

  if (playerNr) {
    avatar = (
      <Badge badgeContent={team} color={color}>
        {avatar}
      </Badge>
    );
  }

  return (
    <ListItem onClick={onClick} button>
      {avatar}
      <ListItemText inset primary={player.name} />
    </ListItem>
  );
}

const StyledSelectPlayers = withStyles(styles)(SelectPlayers);
const StyledPlayerListItem = withStyles(playerItemStyles)(PlayerListItem);

export default StyledSelectPlayers;
