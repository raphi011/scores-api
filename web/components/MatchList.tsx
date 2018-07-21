import React from 'react';
import { withStyles, createStyles } from '@material-ui/core/styles';
import List from '@material-ui/core/List';

import MatchListItem from './MatchListItem';

import { Match } from '../types';
import DayHeader from './DayHeader';

const styles = createStyles({
  root: {
    width: '100%',
  },
});

function isSameDay(date1: Date, date2: Date): boolean {
  return (
    date1.getFullYear() === date2.getFullYear() &&
    date1.getMonth() === date2.getMonth() &&
    date1.getDate() === date2.getDate()
  );
}

interface Props {
  matches: Match[];
  onMatchClick: (Match) => void;
  highlightPlayerId: number;
  classes: any;
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

export default withStyles(styles)(MatchList);
