import { Store } from '../store';
import { Filter } from './reducer';

export const tournamentFilterSelector = (state: Store): Filter =>
  state.tournament.filter;
