import List from '@material-ui/core/List';
import { createStyles, Theme, withStyles } from '@material-ui/core/styles';
import React from 'react';

import { Player } from '../types';
import PlayerListItem from './PlayerListItem';

const styles = (theme: Theme) =>
  createStyles({
    root: {
      width: '100%',
      background: theme.palette.background.paper,
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
  classes: any;
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

      if (pId) { assignedCount += 1; }
    }

    if (selectedNr) {
      onUnsetPlayer(selectedNr);
    } else if (unassigned) {
      onSetPlayer(unassigned, selected, assignedCount === 3);
    }
  };

  isSamePlayer = (p1, p2) => {
    if (!p1 || !p2) { return false; }

    return p1.id === p2.id;
  };

  playerNr = (player: Player): number => {
    const { player1, player2, player3, player4 } = this.props;

    if (this.isSamePlayer(player, player1)) { return 1; }
    else if (this.isSamePlayer(player, player2)) { return 2; }
    else if (this.isSamePlayer(player, player3)) { return 3; }
    else if (this.isSamePlayer(player, player4)) { return 4; }

    return 0;
  };

  render() {
    const { players = [], classes } = this.props;

    return (
      <List className={classes.root}>
        {players.map(p => (
          <PlayerListItem
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

export default withStyles(styles)(SelectPlayers);
