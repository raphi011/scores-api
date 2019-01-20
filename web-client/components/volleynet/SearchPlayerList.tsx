import React from 'react';

import List from '@material-ui/core/List';
import { createStyles, WithStyles, withStyles } from '@material-ui/core/styles';

import SearchPlayerListItem from './SearchPlayerListItem';

import { SearchPlayer } from '../../types';

const styles = createStyles({
  root: {
    width: '100%',
  },
});

interface Props extends WithStyles<typeof styles> {
  players: SearchPlayer[];

  onPlayerClick: (player: SearchPlayer) => void;
}

class SearchPlayerList extends React.PureComponent<Props> {
  render() {
    const { players = [], onPlayerClick, classes } = this.props;

    return (
      <List className={classes.root}>
        {players.map(p => (
          <SearchPlayerListItem key={p.id} onClick={onPlayerClick} player={p} />
        ))}
      </List>
    );
  }
}

export default withStyles(styles)(SearchPlayerList);
