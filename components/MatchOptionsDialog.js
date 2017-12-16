// @flow

import React from 'react';
import DeleteIcon from 'material-ui-icons/Delete';
import CopyIcon from 'material-ui-icons/ContentCopy';
import Divider from 'material-ui/Divider';
import Dialog, { DialogTitle } from 'material-ui/Dialog';
import { withStyles } from 'material-ui/styles';
import List, { ListItem, ListItemText, ListItemAvatar } from 'material-ui/List';
import Avatar from 'material-ui/Avatar';
import type { Match } from '../types';

const styles = () => ({
  root: {
    width: 250,
  },
});

type Props = {
  onRematch: Match => void,
  match: Match,
  onDelete: Match => void,
  onClose: Event => void,
  onShowPlayer: number => void,
  open: boolean,
  classes: Object,
};

class MatchOptionsDialog extends React.PureComponent<Props> {
  onRematch = () => {
    const { onRematch, match } = this.props;

    onRematch(match);
  };

  onDelete = () => {
    const { onDelete, match } = this.props;

    onDelete(match);
  };

  render() {
    const { classes, onClose, match, onShowPlayer, open } = this.props;

    const playerInfos = open
      ? [
          { playerId: match.team1.player1.id, name: match.team1.player1.name },
          { playerId: match.team1.player2.id, name: match.team1.player2.name },
          { playerId: match.team2.player1.id, name: match.team2.player1.name },
          { playerId: match.team2.player2.id, name: match.team2.player2.name },
        ]
      : [];

    return (
      <Dialog onRequestClose={onClose} open={open}>
        <DialogTitle>Options</DialogTitle>
        <List className={classes.root}>
          {playerInfos.map(({ playerId, name }) => (
            <ListItem
              button
              key={playerId}
              onClick={() => onShowPlayer(playerId)}
            >
              <ListItemText inset primary={name} />
            </ListItem>
          ))}
          <Divider />
          <ListItem button onClick={this.onRematch}>
            <ListItemAvatar>
              <Avatar>
                <CopyIcon />
              </Avatar>
            </ListItemAvatar>
            <ListItemText primary="Rematch" />
          </ListItem>
          <Divider />
          <ListItem button onClick={this.onDelete}>
            <ListItemAvatar>
              <Avatar>
                <DeleteIcon />
              </Avatar>
            </ListItemAvatar>
            <ListItemText primary="Delete" />
          </ListItem>
        </List>
      </Dialog>
    );
  }
}

export default withStyles(styles)(MatchOptionsDialog);
