/* eslint-disable import/prefer-default-export */
export function createReducer(
  initialState: object,
  handlers: { [actionType: string]: (state: object, action: object) => object },
) {
  return function reducer(state = initialState, action) {
    if (handlers.hasOwnProperty(action.type)) {
      return handlers[action.type](state, action);
    }

    return state;
  };
}
