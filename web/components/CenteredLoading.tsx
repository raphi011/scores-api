import React from 'react';
import { withStyles } from '@material-ui/core/styles';
import CircularProgress from '@material-ui/core/CircularProgress';
import { Classes } from 'types';

const styles = () => ({
  container: {
    height: '500px',
    display: 'flex',
    alignItems: 'center',
    justifyContent: 'center',
  },
});

interface Props {
  classes: Classes;
}

const CenteredLoading = ({ classes }: Props) => (
  <div className={classes.container}>
    <CircularProgress />
  </div>
);

export default withStyles(styles)(CenteredLoading);
