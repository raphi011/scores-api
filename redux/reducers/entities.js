import * as actionNames from '../actionNames';
import { createReducer } from '../reduxHelper';

function receiveEntities(state, action) {
  
}

const reducer = createReducer(
  {},
  {
    [actionNames.RECEIVE_ENTITIES]: receiveEntities,
    // [actionNames.RECEIVE_PLAYER_MATCHES]: receivePlayerMatches,
    // [actionNames.RECEIVE_STATISTIC]: receiveStatistic,
    // [actionNames.RECEIVE_STATISTICS]: receiveStatistics,
    // [actionNames.RECEIVE_MATCH]: receiveMatch,
    // [actionNames.RECEIVE_PLAYER]: receivePlayer,
    // [actionNames.RECEIVE_MATCHES]: receiveMatches,
    // [actionNames.RECEIVE_PLAYERS]: receivePlayers,
    // [actionNames.LOGGEDOUT]: loggedOut,
    // [actionNames.LOGIN]: login,
    // [actionNames.SET_USER_OR_LOGINROUTE]: setUserOrLoginroute,
    // [actionNames.SET_STATUS]: setStatus,
    // [actionNames.CLEAR_STATUS]: clearStatus,
    // [actionNames.REMOVE_MATCH]: removeMatch,
  },
);

export default reducer;
