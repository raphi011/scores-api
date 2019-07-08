import { Action } from 'redux';

export function createReducer(
  initialState: any,
  handlers: { [actionType: string]: (state: any, action: any) => object },
) {
  return function reducer(state = initialState, action: Action) {
    if (Object.prototype.hasOwnProperty.call(handlers, action.type)) {
      return handlers[action.type](state, action);
    }

    return state;
  };
}
