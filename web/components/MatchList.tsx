import React from 'react';
import { withStyles } from '@material-ui/core/styles';
import List from '@material-ui/core/List';
import ListItem from '@material-ui/core/ListItem';
import Chip from '@material-ui/core/Chip';

import MatchListItem from './MatchListItem';
import { formatDate } from '../utils/dateFormat';

import { Match, Classes } from '../types';

const styles = () => ({
  root: {
    width: '100%',
  },
});

interface Props {
  matches: Match[];
  onMatchClick: (Match) => void;
  highlightPlayerId: number;
  classes: Classes;
}

function isSameDay(date1: Date, date2: Date): boolean {
  return (
    date1.getFullYear() === date2.getFullYear() &&
    date1.getMonth() === date2.getMonth() &&
    date1.getDate() === date2.getDate()
  );
}

type DayHeaderProps = {
  date: Date;
};

const DayHeader = ({ date }: DayHeaderProps) => (
  <ListItem dense style={{ justifyContent: 'center' }}>
    <Chip label={formatDate(date)} />
  </ListItem>
);

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

const StyledMatchList = withStyles(styles)(MatchList);

export default StyledMatchList;
