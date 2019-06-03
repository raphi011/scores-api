import React from 'react';

import ListItem from '@material-ui/core/ListItem';
import { createStyles, WithStyles, withStyles } from '@material-ui/core/styles';
import Typography from '@material-ui/core/Typography';

import { formatDate } from '../utils/date';

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
    <Typography variant="h2" color="primary">
      {formatDate(date)} {appendix}
    </Typography>
  </ListItem>
);

export default withStyles(styles)(DayHeader);
