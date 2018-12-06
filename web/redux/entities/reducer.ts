import * as actionNames from '../actionNames';
import { norm } from '../entitySchemas';
import { createReducer } from '../reduxHelper';

import { EntityName, EntityType } from '../../types';

export interface IEntityStore {
  [key: string]: { values: { [key: number]: EntityType } };
}

export interface IEntityMap {
  [key: string]: { [key: number]: EntityType };
}

export interface IReceiveEntityAction {
  payload: any[];
  entityName: EntityName;
  assignId?: boolean;
  listOptions?: {
    [key: string]: {
      name: string;
      key?: number;
      mode?: 'replace' | 'append';
    };
  };
}

export interface IDeleteEntityAction {
  payload: any;
  entityName: EntityName;
  listNames: string[];
}

export const initialEntitiesState = {
  group: { values: {}, all: [] },
  match: { values: {}, all: [], byPlayer: {}, byGroup: {} },
  player: { values: {}, all: [], byGroup: {} },
  statistic: {
    all: [],
    byGroup: {},
    byPlayer: {},
    byPlayerTeam: {},
    values: {},
  },
  team: { values: {} },
  tournament: { values: {}, all: [], byLeague: {} },
  user: { values: {}, all: [] },
  volleynetplayer: {
    all: [],
    ladder: {},
    search: [],
    values: {},
  },
};

function deleteEntities(state, action: IDeleteEntityAction) {
  const { entityName, listNames, payload } = action;

  const deletedId = payload.id;
  // TODO: delete from values
  // TODO: delete from [listName][listKey] lists
  const newLists = listNames.reduce((listObj, name) => {
    listObj[name] = state[entityName][name].filter(id => id !== deletedId);
    return listObj;
  }, {});

  return {
    ...state,
    [entityName]: {
      ...state[entityName],
      ...newLists,
    },
  };
}

function receiveEntities(state: IEntityStore, action: IReceiveEntityAction) {
  // STEP 1: normalize entites
  const { entityName, payload, assignId = false, listOptions = {} } = action;

  const { entities, result } = norm(entityName, payload, assignId);

  const newState = { ...state };

  // STEP 2: add entities to entity map(s)
  Object.keys(entities).forEach((entityKey: EntityName) => {
    const statePart = { ...state[entityKey] };

    let newIds;

    if (entityKey === action.entityName) {
      newIds = result;
    } else {
      newIds = Object.keys(entities[entityKey]).map(n =>
        Number.parseInt(n, 10),
      );
    }

    statePart.values = {
      ...state[entityKey].values,
      ...entities[entityKey],
    };

    const options = listOptions[entityKey];

    // STEP 3: append or replace ids
    if (options) {
      let list = [];

      if (options.mode === 'append') {
        const previousList = options.key
          ? (state[entityKey][options.name] || {})[options.key]
          : state[entityKey][options.name];

        if (previousList) {
          list = previousList;
        }
      }

      list = [...list, ...newIds];

      statePart[options.name] = options.key
        ? { ...(state[entityKey][options.name] || {}), [options.key]: list }
        : list;
    }

    newState[entityKey] = statePart;
  });

  return newState;
}

function receiveUser(state: IEntityStore, action) {
  const { user } = action.payload;

  if (!user || !user.player) {
    return state;
  }

  return receiveEntities(state, {
    entityName: 'player',
    payload: user.player,
  });
}

const reducer = createReducer(initialEntitiesState, {
  [actionNames.RECEIVE_ENTITIES]: receiveEntities,
  [actionNames.DELETE_ENTITIES]: deleteEntities,
  [actionNames.SET_USER_OR_LOGINROUTE]: receiveUser,
});

export default reducer;
