import List from '@material-ui/core/List';
import { createStyles, withStyles } from '@material-ui/core/styles';
import React from 'react';

import SearchPlayerListItem from './SearchPlayerListItem';

import { VolleynetSearchPlayer } from '../../types';

const styles = createStyles({
  root: {
    width: '100%',
  },
});

interface IProps {
  players: VolleynetSearchPlayer[];
  onPlayerClick: (VolleynetSearchPlayer) => void;
  classes: any;
}

class SearchPlayerList extends React.PureComponent<IProps> {
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
