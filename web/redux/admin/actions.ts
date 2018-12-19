import { ApiAction } from '../../redux/api/actions';
import * as actionNames from '../actionNames';

export const loadVolleynetScrapeJobsAction = (): ApiAction => ({
  method: 'GET',
  success: actionNames.RECEIVE_SCRAPE_JOBS,
  type: actionNames.API,
  url: 'admin/volleynet/scrape/report',
});

export const runJobAction = (jobName: string): ApiAction => ({
  method: 'POST',
  params: { job: jobName },
  type: actionNames.API,
  url: 'admin/volleynet/scrape/run',
});

export const loadUsersAction = (): ApiAction => ({
  method: 'GET',
  success: actionNames.RECEIVE_ENTITIES,
  successParams: {
    entityName: 'user',
    listOptions: {
      user: {
        mode: 'replace',
        name: 'all',
      },
    },
  },
  type: actionNames.API,
  url: `admin/users`,
});