import React from "react";
import DeleteIcon from "material-ui-icons/Delete";
import CopyIcon from "material-ui-icons/ContentCopy";
import Dialog from 'material-ui/Dialog';
import { withStyles } from "material-ui/styles";
import List, { ListItem, ListItemText, ListItemAvatar } from "material-ui/List";
import Avatar from 'material-ui/Avatar';

const styles = theme => ({
  root: {
    width: 250,
  },
});

const MatchOptionsDialog = ({
  classes,
  onClose,
  onRematch,
  onDelete,
  match,
  open,
}) => (
  <Dialog onRequestClose={onClose} open={open}>
      <List className={classes.root}>
        <ListItem button onClick={() => onDelete(match)}>
          <ListItemAvatar>
            <Avatar>
              <DeleteIcon />
            </Avatar>
          </ListItemAvatar>
          <ListItemText primary="Delete" />
        </ListItem>
        <ListItem button onClick={() => onRematch(match)}>
          <ListItemAvatar>
            <Avatar>
              <CopyIcon />
            </Avatar>
          </ListItemAvatar>
          <ListItemText primary="Rematch" />
        </ListItem>
      </List>
  </Dialog>
);

export default withStyles(styles)(MatchOptionsDialog);
