// @flow

import React from 'react';
import { withStyles } from 'material-ui/styles';
import List, { ListItem, ListItemText } from 'material-ui/List';
import Chip from 'material-ui/Chip';
import Typography from 'material-ui/Typography';
import { formatDate } from '../utils/dateFormat';
import type { Match, Team } from '../types';

const styles = () => ({
  root: {
    width: '100%',
  },
  listContainer: {
    display: 'flex',
    flexDirection: 'row',
    alignItems: 'center',
    width: '100%',
  },
  team: { flex: '1 1 0' },
  points: { fontWeight: 'lighter', flex: '2 2 0' },
});

type Props = {
  matches: Array<Match>,
  onMatchClick: Match => void,
  classes: Object,
};

function getTeamName(team: Team) {
  if (team.Name) return team.Name;

  return (
    <span>
      {team.player1.name}
      <br />
      {team.player2.name}
    </span>
  );
}

function isSameDay(date1: Date, date2: Date): boolean {
  return (
    date1.getFullYear() === date2.getFullYear() &&
    date1.getMonth() === date2.getMonth() &&
    date1.getDate() === date2.getDate()
  );
}

const DayHeader = ({ date }) => (
  <ListItem dense style={{ justifyContent: 'center' }}>
    <Chip label={formatDate(date)} />
  </ListItem>
);

function MatchList({ matches = [], onMatchClick, classes }: Props) {
  return (
    <List className={classes.root}>
      {matches.map((m, i) => {
        const currentDate = new Date(m.createdAt);
        const lastDate = i ? new Date(matches[i - 1].createdAt) : null;

        return (
          <React.Fragment>
            {!lastDate || !isSameDay(currentDate, lastDate) ? (
              <DayHeader date={currentDate} />
            ) : // <ListItem key={currentDate.getTime()}>
            //   <ListItemText />
            // </ListItem>
            null}
            <ListItem divider key={m.id} button onClick={() => onMatchClick(m)}>
              <ListItemText
                primary={
                  <div className={classes.listContainer}>
                    <Typography className={classes.team} type="body1">
                      {getTeamName(m.team1)}
                    </Typography>
                    <Typography
                      className={classes.points}
                      type="display2"
                      align="center"
                    >
                      {m.scoreTeam1} - {m.scoreTeam2}
                    </Typography>
                    <Typography
                      className={classes.team}
                      type="body1"
                      align="right"
                    >
                      {getTeamName(m.team2)}
                    </Typography>
                  </div>
                }
                // secondary={formatDateTime(currentDate)}
              />
            </ListItem>
          </React.Fragment>
        );
      })}
    </List>
  );
}

const StyledMatchList = withStyles(styles)(MatchList);

export default StyledMatchList;
