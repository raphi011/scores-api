import * as actionNames from "../actionNames";

const reducer = (state, action) => {
  switch (action.type) {
    case actionNames.RECEIVE_MATCHES:
      return {
        ...state,
        matches: action.payload
      };
    case actionNames.RECEIVE_PLAYERS:
      const playersMap = {};
      const playerIDs = [];
      action.payload.forEach(p => {
        playerIDs.push(p.ID);
        playersMap[p.ID] = p;
      });

      return {
        ...state,
        playersMap,
        playerIDs
      };
    case actionNames.LOGGEDOUT:
      return {
        ...state,
        user: action.username
      };
    case actionNames.LOGIN:
      return {
        ...state,
        user: action.username
      };
    case actionNames.SET_STATUS:
      return {
        ...state,
        status: action.status
      };
    case actionNames.CLEAR_STATUS:
      return {
        ...state,
        status: ""
      };
    case actionNames.REMOVE_MATCH:
      const matches = state.matches.filter(m => m.ID !== action.ID);
      return {
        ...state,
        matches
      };
    default:
      return state;
  }
};

export default reducer;

export const statusSelector = state => state.status;
export const matchesSelector = state => state.matches;
export const playersSelector = state => ({
  playersMap: state.playersMap,
  playerIDs: state.playerIDs
});
export const userSelector = state => ({
  isLoggedIn: !!state.user,
  user: state.user
});
