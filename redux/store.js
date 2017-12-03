import { createStore, applyMiddleware } from "redux";
import reducer from "./reducers/reducer";
import apiMiddleware from "./apiMiddleware";
import { composeWithDevTools } from "redux-devtools-extension";
import { serverAction } from "./apiMiddleware";

export async function dispatchActions(
  dispatch,
  isServer,
  req,
  res,
  actions = []
) {
  for (let i = 0; i < actions.length; i++) {
    const action = isServer ? serverAction(actions[i], req, res) : actions[i];

    await dispatch(action);
  }
}

const initialState = {
  user: null,
  loginRoute: null,
  status: "",
  matches: [],
  playersMap: {},
  playerIDs: []
};

const initStore = (state = initialState) =>
  createStore(
    reducer,
    state,
    composeWithDevTools(applyMiddleware(apiMiddleware))
  );

export default initStore;
