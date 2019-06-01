import React from 'react';

import Button from '@material-ui/core/Button';
import ListItem from '@material-ui/core/ListItem';
import ListItemIcon from '@material-ui/core/ListItemIcon';
import ListItemSecondaryAction from '@material-ui/core/ListItemSecondaryAction';
import ListItemText from '@material-ui/core/ListItemText';
import { createStyles, WithStyles, withStyles } from '@material-ui/core/styles';

import PauseIcon from '@material-ui/icons/Pause';
import PlayIcon from '@material-ui/icons/PlayArrow';
import StopIcon from '@material-ui/icons/Stop';
import WarningIcon from '@material-ui/icons/Warning';

import { ScrapeJob } from '../../types';

const styles = createStyles({
  root: {
    width: '100%',
  },
});

interface Props extends WithStyles<typeof styles> {
  job: ScrapeJob;

  onAction: (jobName: string) => void;
}

class JobListItem extends React.PureComponent<Props> {
  render() {
    const { job, onAction } = this.props;

    return (
      <ListItem key={job.name}>
        <ListItemIcon>{stateToString(job.execution.state)}</ListItemIcon>
        <ListItemText primary={job.name} />
        <ListItemSecondaryAction>
          <Button onClick={() => onAction(job.name)}>run</Button>
        </ListItemSecondaryAction>
      </ListItem>
    );
  }
}

function stateToString(state: number): JSX.Element {
  switch (state) {
    case 0:
      return <StopIcon />;
    case 1:
      return <StopIcon />;
    case 2:
      return <PauseIcon />;
    case 3:
      return <PlayIcon />;
    case 4:
      return <WarningIcon />;
    default:
      return <span />;
  }
}

export default withStyles(styles)(JobListItem);
