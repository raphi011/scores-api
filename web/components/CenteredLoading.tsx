import CircularProgress from '@material-ui/core/CircularProgress';
import { createStyles, WithStyles, withStyles } from '@material-ui/core/styles';
import React from 'react';

const styles = createStyles({
  container: {
    alignItems: 'center',
    display: 'flex',
    height: '500px',
    justifyContent: 'center',
  },
});

interface Props extends WithStyles<typeof styles> {}

const CenteredLoading = ({ classes }: Props) => (
  <div className={classes.container}>
    <CircularProgress />
  </div>
);

export default withStyles(styles)(CenteredLoading);
