/* eslint-disable import/prefer-default-export */
export function createReducer(
  initialState: Object,
  handlers: { [key: string]: (any, Object) => any },
) {
  return function reducer(state = initialState, action) {
    if (handlers.hasOwnProperty(action.type)) {
      return handlers[action.type](state, action);
    }

    return state;
  };
}
