import React from 'react';

import { createStyles, WithStyles, withStyles } from '@material-ui/core/styles';
import Typography from '@material-ui/core/Typography';

import { formatDate } from '../utils/date';

const styles = createStyles({
  container: {
    paddingTop: 0,
    marginBottom: '20px',
  },
});

interface Props extends WithStyles<typeof styles> {
  appendix?: string;
  date: Date;
}

const DayHeader = ({ date, appendix, classes }: Props) => (
  <Typography
    component="a"
    variant="h2"
    color="primary"
    className={classes.container}
  >
    {formatDate(date)} {appendix}
  </Typography>
);

export default withStyles(styles)(DayHeader);
