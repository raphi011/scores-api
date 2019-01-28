import React from 'react';

import { Button } from '@material-ui/core';
import Paper from '@material-ui/core/Paper';
import {
  createStyles,
  Theme,
  WithStyles,
  withStyles,
} from '@material-ui/core/styles';
import Typography from '@material-ui/core/Typography';
import AddIcon from '@material-ui/icons/Add';

import JobList from '../../components/admin/JobList';
import UserList from '../../components/admin/UserList';
import EditUserDialog from '../../containers/admin/EditUserDialogContainer';
import withAuth from '../../containers/AuthContainer';
import Layout from '../../containers/LayoutContainer';
import {
  loadUsersAction,
  loadVolleynetScrapeJobsAction,
  runJobAction,
} from '../../redux/admin/actions';
import { scrapeJobsSelector } from '../../redux/admin/selectors';
import { allUsersSelector } from '../../redux/entities/selectors';
import { Store } from '../../redux/store';
import { ScrapeJob, User } from '../../types';

const styles = (theme: Theme) =>
  createStyles({
    container: {
      margin: theme.spacing.unit,
    },
    paper: {
      margin: '10px 0',
      padding: theme.spacing.unit * 2,
    },
    title: { marginTop: '25px' },
    userHeader: {
      display: 'flex',
      flexDirection: 'row',
      justifyContent: 'space-between',
      width: '100%',
    },
  });

interface Props extends WithStyles<typeof styles> {
  jobs: ScrapeJob[];
  users: User[];

  loadScrapeJobs: () => void;
  runJob: (jobName: string) => void;
}

type State = {
  isEditUserOpen: boolean;
  editUser?: User;
};

class Home extends React.Component<Props, State> {
  static mapDispatchToProps = {
    loadScrapeJobs: loadVolleynetScrapeJobsAction,
    runJob: runJobAction,
  };

  static buildActions() {
    return [loadVolleynetScrapeJobsAction(), loadUsersAction()];
  }

  static mapStateToProps(state: Store) {
    const jobs = scrapeJobsSelector(state);
    const users = allUsersSelector(state);

    return {
      jobs,
      users,
    };
  }

  state = {
    editUser: null,
    isEditUserOpen: false,
  };

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

  newUser = () => {
    this.setState({
      editUser: null,
      isEditUserOpen: true,
    });
  };

  editUser = (user: User) => {
    this.setState({
      editUser: user,
      isEditUserOpen: true,
    });
  };

  onCloseEditUser = () => {
    this.setState({
      isEditUserOpen: false,
    });
  };

  render() {
    const { classes, users, runJob, jobs } = this.props;

    return (
      <Layout title={{ text: 'Settings', href: '' }}>
        <div className={classes.container}>
          <Typography variant="h5" className={classes.title}>
            Volleynet scrape jobs
          </Typography>
          <Paper className={classes.paper}>
            <JobList jobs={jobs} onAction={runJob} />
          </Paper>
          <div className={classes.userHeader}>
            <Typography variant="h5" className={classes.container}>
              Users
            </Typography>
            <Button color="primary" onClick={this.newUser}>
              <AddIcon />
              Add
            </Button>
          </div>
          <Paper className={classes.paper}>
            <UserList onClick={this.editUser} users={users} />
          </Paper>
        </div>
        <EditUserDialog
          onClose={this.onCloseEditUser}
          open={this.state.isEditUserOpen}
          user={this.state.editUser}
        />
      </Layout>
    );
  }
}

export default withAuth(withStyles(styles)(Home));
