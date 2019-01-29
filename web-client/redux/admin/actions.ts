import { ApiAction } from '../../redux/api/actions';
import { EntityName } from '../../types';
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

export const updateUserAction = (
  email: string,
  password: string,
): ApiAction => ({
  body: JSON.stringify({ email, password }),
  method: 'POST',
  successStatus: `User ${email} updated`,
  type: actionNames.API,
  url: 'admin/users',
});

export const loadUsersAction = (): ApiAction => ({
  method: 'GET',
  success: actionNames.RECEIVE_ENTITIES,
  successParams: {
    entityName: EntityName.User,
    listOptions: {
      [EntityName.User]: {
        mode: 'replace',
        name: 'all',
      },
    },
  },
  type: actionNames.API,
  url: `admin/users`,
});
