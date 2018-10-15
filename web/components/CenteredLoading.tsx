import CircularProgress from '@material-ui/core/CircularProgress';
import { createStyles, withStyles } from '@material-ui/core/styles';
import React from 'react';

const styles = createStyles({
  container: {
    alignItems: 'center',
    display: 'flex',
    height: '500px',
    justifyContent: 'center',
  },
});

interface IProps {
  classes: any;
}

const CenteredLoading = ({ classes }: IProps) => (
  <div className={classes.container}>
    <CircularProgress />
  </div>
);

export default withStyles(styles)(CenteredLoading);
