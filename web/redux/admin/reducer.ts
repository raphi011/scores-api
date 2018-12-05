import { ScrapeJob } from '../../types';
import * as actionNames from '../actionNames';
import { createReducer } from '../reduxHelper';

export const initialAdminState = {
  scrapeJobs: [],
};

export interface IAdminStore {
  scrapeJobs: ScrapeJob[];
}

function receiveScrapeJobs(_: IAdminStore, action): IAdminStore {
  return {
    scrapeJobs: action.payload,
  };
}

const reducer = createReducer(initialAdminState, {
  [actionNames.RECEIVE_SCRAPE_JOBS]: receiveScrapeJobs,
});

export default reducer;
