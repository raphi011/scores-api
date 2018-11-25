import Card from '@material-ui/core/Card';
import List from '@material-ui/core/List';
import { createStyles, withStyles } from '@material-ui/core/styles';
import React from 'react';

import TournamentListItem from './TournamentListItem';

import { Tournament } from '../../types';

const styles = createStyles({
  root: {
    padding: '0',
    width: '100%',
  },
});

interface IProps {
  tournaments: Tournament[];
  onTournamentClick: (t: Tournament) => void;
  classes: any;
}

class TournamentList extends React.PureComponent<IProps> {
  render() {
    const { tournaments = [], onTournamentClick, classes } = this.props;

    return (
      <Card>
        <List className={classes.root}>
          {tournaments.map(t => (
            <TournamentListItem
              key={t.id}
              onClick={onTournamentClick}
              tournament={t}
            />
          ))}
        </List>
      </Card>
    );
  }
}

export default withStyles(styles)(TournamentList);
