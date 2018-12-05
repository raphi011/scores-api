import { IStore } from '../store';

export const scrapeJobsSelector = (state: IStore) => state.admin.scrapeJobs;
