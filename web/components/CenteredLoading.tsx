import React from 'react';
import { withStyles, createStyles } from '@material-ui/core/styles';
import CircularProgress from '@material-ui/core/CircularProgress';

const styles = createStyles({
  container: {
    height: '500px',
    display: 'flex',
    alignItems: 'center',
    justifyContent: 'center',
  },
});

interface Props {
  classes: any;
}

const CenteredLoading = ({ classes }: Props) => (
  <div className={classes.container}>
    <CircularProgress />
  </div>
);

export default withStyles(styles)(CenteredLoading);
