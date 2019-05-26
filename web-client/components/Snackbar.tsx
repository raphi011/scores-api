import React from 'react';

import IconButton from '@material-ui/core/IconButton';
import Snackbar from '@material-ui/core/Snackbar';
import {
  createStyles,
  Theme,
  WithStyles,
  withStyles,
} from '@material-ui/core/styles';
import CloseIcon from '@material-ui/icons/Close';

const styles = (theme: Theme) =>
  createStyles({
    close: {
      padding: theme.spacing(0.5),
    },
  });

interface Props extends WithStyles<typeof styles> {
  status: string;
  open: boolean;

  onClose: () => void;
}

const SimpleSnackbar = ({ classes, onClose, status, open }: Props) => (
  <Snackbar
    anchorOrigin={{
      horizontal: 'center',
      vertical: 'bottom',
    }}
    open={open}
    autoHideDuration={6000}
    onClose={handleRequestClose(onClose)}
    message={<span id="message-id">{status}</span>}
    action={
      <IconButton
        key="close"
        aria-label="Close"
        color="inherit"
        className={classes.close}
        onClick={onClose}
      >
        <CloseIcon />
      </IconButton>
    }
  />
);

const handleRequestClose = (onClose: () => void) => (
  _: any,
  reason: string,
) => {
  if (reason === 'clickaway') {
    return;
  }

  onClose();
};

export default withStyles(styles)(SimpleSnackbar);
