/* eslint-disable import/prefer-default-export */
export function createReducer(
  initialState: object,
  handlers: { [actionType: string]: (state: any, action: object) => any },
) {
  return function reducer(state = initialState, action) {
    if (handlers.hasOwnProperty(action.type)) {
      return handlers[action.type](state, action);
    }

    return state;
  };
}
