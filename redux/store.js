import { createStore, applyMiddleware } from 'redux'
import reducer from './reducers/reducer';
import apiMiddleware from './apiMiddleware';
import { composeWithDevTools } from 'redux-devtools-extension'

const initialState = {
  user: null,
  loginRoute: null,
  status: '',
  matches: [],
  playersMap: {},
  playerIDs: [],
}

const initStore = (state = initialState) => (
  createStore(reducer, state, composeWithDevTools(applyMiddleware(apiMiddleware)))
)

export default initStore;
