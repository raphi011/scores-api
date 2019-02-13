import React from 'react';

import {
  createStyles,
  Theme,
  WithStyles,
  withStyles,
} from '@material-ui/core/styles';
import Typography from '@material-ui/core/Typography';

const styles = (theme: Theme) =>
  createStyles({
    data: {
      color: theme.palette.grey[700],
      fontSize: '16px',
      textTransform: 'capitalize',
    },
    label: {
      color: theme.palette.grey[400],
      fontSize: '14px',
      textTransform: 'uppercase',
    },
    root: {
      display: 'flex',
      flexDirection: 'column',
    },
  });

interface Props extends WithStyles<typeof styles> {
  label: string;
  data: string;
  size?: 'sm' | 'lg';
}

export default withStyles(styles)(({ classes, label, data }: Props) => {
  return (
    <div className={classes.root}>
      <Typography className={classes.label}>{label}</Typography>
      <Typography className={classes.data}>{data.toLowerCase()}</Typography>
    </div>
  );
});
