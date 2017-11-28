import * as actionNames from '../actionNames';

export const loadMatchesAction = () => ({
  type: actionNames.API,
  method: 'GET',
  url: 'matches',
  success: actionNames.RECEIVE_MATCHES,
});

export const loadPlayersAction = () => ({
  type: actionNames.API,
  method: 'GET',
  url: 'players',
  success: actionNames.RECEIVE_PLAYERS,
});

export const createNewMatchAction = (match) => ({
  type: actionNames.API,
  method: 'POST',
  url: 'matches',
  body: JSON.stringify(match),
  successStatus: 'New Match created',
});

export const deleteMatchAction = (match) => ({
  type: actionNames.API,
  method: 'DELETE',
  url: `matches/${match.ID}`,
  success: actionNames.REMOVE_MATCH,
  successParams: { ID: match.ID },
  successStatus: 'Match deleted',

});

export const setStatusAction = (status) => ({
  type: actionNames.SET_STATUS,
  status,
});

export const clearStatusAction = () => ({
  type: actionNames.CLEAR_STATUS
});
