import React from 'react';

import List from '@material-ui/core/List';

import { ScrapeJob } from '../../types';
import JobListItem from './JobListItem';

interface Props {
  jobs: ScrapeJob[];
  onAction: (jobName: string) => void;
}

class JobList extends React.PureComponent<Props> {
  render() {
    const { jobs = [], onAction } = this.props;

    return (
      <List dense>
        {jobs.map(j => (
          <JobListItem key={j.job.name} job={j} onAction={onAction} />
        ))}
      </List>
    );
  }
}

export default JobList;
