// @flow

import React from 'react';
import { withStyles } from '@material-ui/core/styles';
import List from '@material-ui/core/List';

import PlayerListItem from './PlayerListItem';

import type { VolleynetPlayer } from '../../types';

const styles = () => ({
  root: {
    width: '100%',
  },
});

type Props = {
  players: Array<VolleynetPlayer>,
  onPlayerClick: VolleynetPlayer => void,
  classes: Object,
};

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
