import React from 'react';

import List from '@material-ui/core/List';
import { createStyles, WithStyles, withStyles } from '@material-ui/core/styles';

import { Player } from '../../types';
import LadderItem from './LadderItem';

const styles = createStyles({
  root: {
    width: '100%',
  },
});

interface Props extends WithStyles<typeof styles> {
  players: Player[];
}

class SearchPlayerList extends React.PureComponent<Props> {
  render() {
    const { players = [], classes } = this.props;

    return (
      <List className={classes.root}>
        {players.map(p => (
          <LadderItem key={p.id} player={p} />
        ))}
      </List>
    );
  }
}

export default withStyles(styles)(SearchPlayerList);
