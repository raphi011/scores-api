import React from 'react';

import List from '@material-ui/core/List';
import { createStyles, WithStyles, withStyles } from '@material-ui/core/styles';

import DayHeader from './DayHeader';
import MatchListItem from './MatchListItem';

import { Match } from '../types';

const styles = createStyles({
  root: {
    width: '100%',
  },
});

interface Props extends WithStyles<typeof styles> {
  matches: Match[];
  highlightPlayerId: number;

  onMatchClick: (match: Match) => void;
}

class MatchList extends React.PureComponent<Props> {
  render() {
    const {
      matches = [],
      highlightPlayerId,
      onMatchClick,
      classes,
    } = this.props;

    return (
      <List className={classes.root}>
        {matches.map((m, i) => {
          const currentDate = new Date(m.createdAt);
          const lastDate = i ? new Date(matches[i - 1].createdAt) : null;
          const showHeader = !lastDate || !isSameDay(currentDate, lastDate);
          const matchKey = m.id;

          return (
            <React.Fragment key={matchKey}>
              {showHeader ? <DayHeader date={currentDate} /> : null}
              <MatchListItem
                key={matchKey}
                onMatchClick={onMatchClick}
                highlightPlayerId={highlightPlayerId}
                match={m}
              />
            </React.Fragment>
          );
        })}
      </List>
    );
  }
}

function isSameDay(date1: Date, date2: Date): boolean {
  return (
    date1.getFullYear() === date2.getFullYear() &&
    date1.getMonth() === date2.getMonth() &&
    date1.getDate() === date2.getDate()
  );
}

export default withStyles(styles)(MatchList);
