import { Action, Dispatch } from 'redux';

type ThunkAction = (dispatch: Dispatch) => Promise<void>;

const thunkMiddleware = (store: any) => (next: any) => async (
  action: Action | ThunkAction,
) => {
  if (typeof action === 'function') {
    await action(store.dispatch);
  } else {
    next(action);
  }
};

export default thunkMiddleware;
