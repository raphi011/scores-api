import { ApiAction } from '../../types';
import * as actionNames from '../actionNames';

export const loadVolleynetScrapeJobsAction = (): ApiAction => ({
  method: 'GET',
  success: actionNames.RECEIVE_SCRAPE_JOBS,
  type: actionNames.API,
  url: 'admin/volleynet/scrape/report',
});
