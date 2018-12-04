import { createStyles, Theme, withStyles } from '@material-ui/core/styles';
import React from 'react';

import Paper from '@material-ui/core/Paper';
import Typography from '@material-ui/core/Typography';

import PauseIcon from '@material-ui/icons/Pause';
import PlayIcon from '@material-ui/icons/PlayArrow';
import StopIcon from '@material-ui/icons/Stop';
import WarningIcon from '@material-ui/icons/Warning';

import Button from '@material-ui/core/Button';

import List from '@material-ui/core/List';
import ListItem from '@material-ui/core/ListItem';
import ListItemIcon from '@material-ui/core/ListItemIcon';
import ListItemSecondaryAction from '@material-ui/core/ListItemSecondaryAction';
import ListItemText from '@material-ui/core/ListItemText';

import withAuth from '../containers/AuthContainer';
import Layout from '../containers/LayoutContainer';
import {
  loadVolleynetScrapeJobsAction,
  runJobAction,
} from '../redux/actions/admin';
import { scrapeJobsSelector } from '../redux/reducers/admin';
import { ScrapeJob } from '../types';

const styles = (theme: Theme) =>
  createStyles({
    container: {
      margin: theme.spacing.unit,
    },
    paper: {
      margin: '10px 0',
      padding: theme.spacing.unit * 2,
    },
  });

interface IProps {
  jobs: ScrapeJob[];
  loadScrapeJobs: () => void;
  runJob: (jobName: string) => void;
  classes: any;
}

class Home extends React.Component<IProps> {
  static mapDispatchToProps = {
    loadScrapeJobs: loadVolleynetScrapeJobsAction,
    runJob: runJobAction,
  };

  static buildActions() {
    return [loadVolleynetScrapeJobsAction()];
  }

  static mapStateToProps(state) {
    const jobs = scrapeJobsSelector(state);

    return {
      jobs,
    };
  }

  interval: NodeJS.Timer;

  componentDidMount() {
    this.interval = setInterval(this.loadScrapeJobs, 5000);
  }

  componentWillUnmount() {
    clearInterval(this.interval);
  }

  loadScrapeJobs = () => {
    const { loadScrapeJobs } = this.props;

    loadScrapeJobs();
  };

  runJob = (jobName: string) => {
    const { runJob } = this.props;

    runJob(jobName);
  };

  render() {
    const { classes, jobs } = this.props;

    return (
      <Layout title={{ text: 'Settings', href: '' }}>
        <div className={classes.container}>
          <Typography variant="h5">Volleynet scrape jobs</Typography>
          <Paper className={classes.paper}>
            <List dense>
              {jobs.map(j => (
                <ListItem>
                  <ListItemIcon>{stateToString(j.state)}</ListItemIcon>
                  <ListItemText primary={j.job.name} />
                  <ListItemSecondaryAction>
                    <Button
                      onClick={() => this.runJob(j.job.name)}
                      className={classes.button}
                    >
                      run
                    </Button>
                  </ListItemSecondaryAction>
                </ListItem>
              ))}
            </List>
          </Paper>
        </div>
      </Layout>
    );
  }
}

function stateToString(state: number) {
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
  }
}

export default withAuth(withStyles(styles)(Home));
