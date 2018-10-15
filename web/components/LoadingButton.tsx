import Button from '@material-ui/core/Button';
import CircularProgress from '@material-ui/core/CircularProgress';
import { createStyles, Theme, withStyles } from '@material-ui/core/styles';
import React from 'react';

const styles = (theme: Theme) =>
  createStyles({
    wrapper: {
      margin: theme.spacing.unit,
      position: 'relative',
    },
    buttonProgress: {
      position: 'absolute',
      top: '50%',
      left: '50%',
      marginTop: -12,
      marginLeft: -12,
    },
  });

const LoadingButton = ({ children, loading, classes }) => (
  <div className={classes.wrapper}>
    <Button
      color="primary"
      fullWidth
      variant="raised"
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
