import { createStyles, Theme, withStyles } from '@material-ui/core/styles';
import React from 'react';

import Paper from '@material-ui/core/Paper';
import Typography from '@material-ui/core/Typography';

import withAuth from '../containers/AuthContainer';
import Layout from '../containers/LayoutContainer';
import { loadVolleynetScrapeJobsAction } from '../redux/actions/admin';
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
  classes: any;
}

class Home extends React.Component<IProps> {
  static mapDispatchToProps = {
    loadScrapeJobs: loadVolleynetScrapeJobsAction,
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

  render() {
    const { classes, jobs } = this.props;

    return (
      <Layout title={{ text: 'Settings', href: '' }}>
        <div className={classes.container}>
          <Typography variant="h5">Volleynet scrape jobs</Typography>
          <Paper className={classes.paper}>
            {jobs.map(j => (
              <Typography variant="h6">
                {j.job.name} - {stateToString(j.state)}
              </Typography>
            ))}
          </Paper>
        </div>
      </Layout>
    );
  }
}

function stateToString(state: number) {
  switch (state) {
    case 0:
      return 'stopped';
    case 1:
      return 'stopping';
    case 2:
      return 'waiting';
    case 3:
      return 'running';
    case 4:
      return 'errored';
    default:
      return 'unknown';
  }
}

export default withAuth(withStyles(styles)(Home));
