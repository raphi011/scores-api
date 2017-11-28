import React from "react";
import DeleteIcon from "material-ui-icons/Delete";
import CopyIcon from "material-ui-icons/ContentCopy";
import Dialog, { DialogTitle } from 'material-ui/Dialog';
import List, { ListItem, ListItemText, ListItemAvatar } from "material-ui/List";
import Avatar from 'material-ui/Avatar';

const MatchOptionsDialog = ({
  classes,
  onClose,
  onClone,
  onDelete,
  match,
  open
}) => (
  <Dialog onRequestClose={onClose} open={open}>
    <DialogTitle>Menu</DialogTitle>
    <div>
      <List>
        <ListItem button onClick={() => onDelete(match)}>
          <ListItemAvatar>
            <Avatar>
              <DeleteIcon />
            </Avatar>
          </ListItemAvatar>
          <ListItemText primary="Delete" />
        </ListItem>
        <ListItem button onClick={() => onClone(match)}>
          <ListItemAvatar>
            <Avatar>
              <CopyIcon />
            </Avatar>
          </ListItemAvatar>
          <ListItemText primary="Clone" />
        </ListItem>
      </List>
    </div>
  </Dialog>
);

export default MatchOptionsDialog;
