import React from 'react';
import DeleteIcon from 'material-ui-icons/Delete';
import CopyIcon from 'material-ui-icons/ContentCopy';
import Divider from 'material-ui/Divider';
import Dialog, { DialogTitle } from 'material-ui/Dialog';
import { withStyles } from 'material-ui/styles';
import List, { ListItem, ListItemText, ListItemAvatar } from 'material-ui/List';
import Avatar from 'material-ui/Avatar';

const styles = () => ({
  root: {
    width: 250,
  },
});

class MatchOptionsDialog extends React.PureComponent {
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

    const playerInfos = open ? [
      { playerID: match.Team1.Player1.ID, name: match.Team1.Player1.Name },
      { playerID: match.Team1.Player2.ID, name: match.Team1.Player2.Name },
      { playerID: match.Team2.Player1.ID, name: match.Team2.Player1.Name },
      { playerID: match.Team2.Player2.ID, name: match.Team2.Player2.Name },
    ] : [];

    return (
      <Dialog onRequestClose={onClose} open={open}>
        <DialogTitle>Options</DialogTitle>
        <List className={classes.root}>
          {playerInfos.map(({ playerID, name }) => (
            <ListItem button key={playerID} onClick={() => onShowPlayer(playerID)}>
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
