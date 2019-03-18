import { Store } from '../store';
import { Filter, FilterOptions } from './reducer';

export const tournamentFilterSelector = (state: Store): Filter | null =>
  state.tournament.filter;

export const filterOptionsSelector = (state: Store): FilterOptions | null =>
  state.tournament.options;
