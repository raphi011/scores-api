import React from 'react';
import { withStyles, createStyles } from '@material-ui/core/styles';
import List from '@material-ui/core/List';

import TournamentListItem from './TournamentListItem';

import { Tournament } from '../../types';

const styles = createStyles({
  root: {
    width: '100%',
  },
});

interface Props {
  tournaments: Tournament[];
  onTournamentClick: () => void;
  classes: any;
}

class TournamentList extends React.PureComponent<Props> {
  render() {
    const { tournaments = [], onTournamentClick, classes } = this.props;

    return (
      <List className={classes.root}>
        {tournaments.map(t => (
          <TournamentListItem
            key={t.id}
            onClick={onTournamentClick}
            tournament={t}
          />
        ))}
      </List>
    );
  }
}

export default withStyles(styles)(TournamentList);
