// @flow 

import React from 'react';
import { withStyles } from 'material-ui/styles';
import Snackbar from 'material-ui/Snackbar';
import IconButton from 'material-ui/IconButton';
import CloseIcon from 'material-ui-icons/Close';

const styles = theme => ({
  close: {
    width: theme.spacing.unit * 4,
    height: theme.spacing.unit * 4,
  },
});

const handleRequestClose = onClose => (event, reason) => {
  if (reason === 'clickaway') {
    return;
  }

  onClose();
};

type Props = {
  onClose: () => void,
  status: string,
  open: boolean,
  classes: Object,
}

const SimpleSnackbar = ({ classes, onClose, status, open }: Props) => (
  <Snackbar
    anchorOrigin={{
      vertical: 'bottom',
      horizontal: 'center',
    }}
    open={open}
    autoHideDuration={6000}
    onRequestClose={handleRequestClose(onClose)}
    SnackbarContentProps={{
      'aria-describedby': 'message-id',
    }}
    message={<span id="message-id">{status}</span>}
    action={[
      <IconButton
        key="close"
        aria-label="Close"
        color="inherit"
        className={classes.close}
        onClick={onClose}
      >
        <CloseIcon />
      </IconButton>,
    ]}
  />
);

export default withStyles(styles)(SimpleSnackbar);
