import { createStore, applyMiddleware } from 'redux';
import { composeWithDevTools } from 'redux-devtools-extension';

import reducer from './reducers/reducer';
import apiMiddleware, { serverAction } from './apiMiddleware';
import type { Store } from './storeTypes';

export async function dispatchActions(
  dispatch,
  isServer,
  req,
  res,
  actions = [],
) {
  for (let i = 0; i < actions.length; i++) {
    const action = isServer ? serverAction(actions[i], req, res) : actions[i];

    await dispatch(action);
  }
}

const initialState = {
  user: null,
  loginRoute: null,
  status: '',
  playersMap: {},
  statisticsMap: {},
  matchesMap: {},
  matchesIds: [],
  playerIds: [],
  statisticIds: [],
};

const initStore = (state: Store = initialState) =>
  createStore(
    reducer,
    state,
    composeWithDevTools(applyMiddleware(apiMiddleware)),
  );

export default initStore;
