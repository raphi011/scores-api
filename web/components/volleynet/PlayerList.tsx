import List from '@material-ui/core/List';
import { createStyles, withStyles } from '@material-ui/core/styles';
import React from 'react';

import PlayerListItem from './PlayerListItem';

import { VolleynetSearchPlayer } from '../../types';

const styles = createStyles({
  root: {
    width: '100%',
  },
});

interface Props {
  players: VolleynetSearchPlayer[];
  onPlayerClick: (VolleynetSearchPlayer) => void;
  classes: any;
}

class PlayerList extends React.PureComponent<Props> {
  render() {
    const { players = [], onPlayerClick, classes } = this.props;

    return (
      <List className={classes.root}>
        {players.map(p => (
          <PlayerListItem key={p.id} onClick={onPlayerClick} player={p} />
        ))}
      </List>
    );
  }
}

export default withStyles(styles)(PlayerList);
