import { createStyles, withStyles } from '@material-ui/core/styles';
import React from 'react';

import Avatar from '@material-ui/core/Avatar';
import Dialog from '@material-ui/core/Dialog';
import DialogTitle from '@material-ui/core/DialogTitle';
import Divider from '@material-ui/core/Divider';
import List from '@material-ui/core/List';
import ListItem from '@material-ui/core/ListItem';
import ListItemText from '@material-ui/core/ListItemText';
import DeleteIcon from '@material-ui/icons/Delete';
import CopyIcon from '@material-ui/icons/FileCopy';
import Link from 'next/link';

import { Match } from '../types';

const styles = createStyles({
  root: {
    width: 250,
  },
});

interface IProps {
  match: Match;
  onDelete: (Match) => void;
  onClose: (Event) => void;
  onShowPlayer: (playerId: number) => void;
  open: boolean;
  classes: any;
}

class MatchOptionsDialog extends React.Component<IProps> {
  shouldComponentUpdate(nextProps) {
    return this.props.open !== nextProps.open;
  }

  onDelete = () => {
    const { onDelete, match } = this.props;

    onDelete(match);
  };

  render() {
    const { classes, onClose, match, onShowPlayer, open } = this.props;

    const rematchLink = match
      ? `/group/createMatch?groupId=${match.groupId}&rematchId=${match.id}`
      : '';
    const playerInfos = match
      ? [
          {
            name: match.team1.player1.name,
            playerId: match.team1.player1.id,
            profileImageUrl: match.team1.player1.profileImageUrl,
          },
          {
            name: match.team1.player2.name,
            playerId: match.team1.player2.id,
            profileImageUrl: match.team1.player2.profileImageUrl,
          },
          {
            name: match.team2.player1.name,
            playerId: match.team2.player1.id,
            profileImageUrl: match.team2.player1.profileImageUrl,
          },
          {
            name: match.team2.player2.name,
            playerId: match.team2.player2.id,
            profileImageUrl: match.team2.player2.profileImageUrl,
          },
        ]
      : [];

    return (
      <Dialog onClose={onClose} open={open}>
        <DialogTitle>Options</DialogTitle>
        <List className={classes.root}>
          {playerInfos.map(({ playerId, profileImageUrl, name }) => (
            <ListItem
              button
              key={playerId}
              onClick={() => onShowPlayer(playerId)}
            >
              <Avatar src={profileImageUrl} />
              <ListItemText inset primary={name} />
            </ListItem>
          ))}
          <Divider />
          <Link href={rematchLink}>
            <ListItem button>
              <Avatar>
                <CopyIcon />
              </Avatar>
              <ListItemText primary="Rematch" />
            </ListItem>
          </Link>
          <Divider />
          <ListItem button onClick={this.onDelete}>
            <Avatar>
              <DeleteIcon />
            </Avatar>
            <ListItemText primary="Delete" />
          </ListItem>
        </List>
      </Dialog>
    );
  }
}

export default withStyles(styles)(MatchOptionsDialog);
