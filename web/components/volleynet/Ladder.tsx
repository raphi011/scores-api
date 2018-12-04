import List from '@material-ui/core/List';
import { createStyles, withStyles } from '@material-ui/core/styles';
import React from 'react';

import LadderItem from './LadderItem';

import { VolleynetPlayer } from '../../types';

const styles = createStyles({
  root: {
    width: '100%',
  },
});

interface IProps {
  players: VolleynetPlayer[];
  classes: any;
}

class SearchPlayerList extends React.PureComponent<IProps> {
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
