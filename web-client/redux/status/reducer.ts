import * as actionNames from '../actionNames';
import { createReducer } from '../reduxHelper';
import { Store } from '../store';
import { SetStatusAction } from './actions';

export const initialStatusState = {
  status: '',
};

export interface StatusStore {
  status: string;
}

function setStatus(_: StatusStore, action: SetStatusAction): StatusStore {
  return {
    status: action.status,
  };
}

function clearStatus(): StatusStore {
  return {
    status: '',
  };
}

const reducer = createReducer(initialStatusState, {
  [actionNames.SET_STATUS]: setStatus,
  [actionNames.CLEAR_STATUS]: clearStatus,
});

export default reducer;

export const statusSelector = (state: Store) => state.status.status;
