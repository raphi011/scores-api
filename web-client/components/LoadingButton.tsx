import Button from '@material-ui/core/Button';
import CircularProgress from '@material-ui/core/CircularProgress';
import { createStyles, withStyles } from '@material-ui/core/styles';
import React from 'react';

const styles = createStyles({
    buttonProgress: {
      left: '50%',
      marginLeft: -12,
      marginTop: -12,
      position: 'absolute',
      top: '50%',
    },
    wrapper: {
      position: 'relative',
    },
  });

const LoadingButton = ({ children, loading, classes }) => (
  <div className={classes.wrapper}>
    <Button
      color="primary"
      fullWidth
      variant="contained"
      disabled={loading}
      type="submit"
    >
      {children}
    </Button>
    {loading && (
      <CircularProgress size={24} className={classes.buttonProgress} />
    )}
  </div>
);

export default withStyles(styles)(LoadingButton);
