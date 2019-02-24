import { Action } from 'redux';
// import { Store } from './store';

/* eslint-disable import/prefer-default-export */
export function createReducer(
  initialState: any,
  handlers: { [actionType: string]: (state: any, action: any) => object },
) {
  return function reducer(state = initialState, action: Action) {
    if (handlers.hasOwnProperty(action.type)) {
      return handlers[action.type](state, action);
    }

    return state;
  };
}
