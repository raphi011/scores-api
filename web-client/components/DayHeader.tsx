import React from 'react';

import ListItem from '@material-ui/core/ListItem';
import { createStyles, WithStyles, withStyles } from '@material-ui/core/styles';
import Typography from '@material-ui/core/Typography';

import { formatDate } from '../utils/dateFormat';

const styles = createStyles({
  container: {
    paddingTop: 0,
  },
});

interface Props extends WithStyles<typeof styles> {
  appendix?: string;
  date: Date;
}

const DayHeader = ({ date, appendix, classes }: Props) => (
  <ListItem className={classes.container} disableGutters>
    <Typography variant="subtitle1">
      {formatDate(date)} {appendix}
    </Typography>
  </ListItem>
);

export default withStyles(styles)(DayHeader);
