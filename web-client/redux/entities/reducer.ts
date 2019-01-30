import * as actionNames from '../actionNames';
import { norm } from '../entitySchemas';
import { createReducer } from '../reduxHelper';

import { EntityName, EntityType } from '../../types';

export type EntityStoreValues = { [key: string]: EntityType };

export type EntityStore = {
  [key in EntityName]: {
    values: EntityStoreValues;
    list?: {
      all: string[];
      [listName: string]: string[];
    };
    by?: {
      [listName: string]: { [key: string]: string[] };
    };
  }
};

export interface ReceiveEntityParams {
  payload?: EntityType[];
  entityName: EntityName;
  assignId?: boolean;
  listOptions?: {
    [key in EntityName]?: {
      name: string;
      key?: string;
      mode?: 'replace' | 'append';
    }
  };
}

export interface DeleteEntityAction {
  payload: any;
  entityName: EntityName;
  listNames: string[];
}

export const initialEntitiesState: EntityStore = {
  player: {
    by: {
      ladder: {
        M: [],
        W: [],
      },
    },
    list: {
      all: [],
      search: [],
    },
    values: {},
  },
  tournament: {
    list: {
      all: [],
      filter: [],
    },
    values: {},
  },
  user: {
    list: {
      all: [],
    },
    values: {},
  },
};

function receiveEntities(state: EntityStore, action: ReceiveEntityParams) {
  // STEP 1: normalize entites
  const { entityName, payload, assignId = false, listOptions = {} } = action;

  const { entities, result } = norm(entityName, payload, assignId);

  const newState = { ...state };

  // STEP 2: add entities to entity map(s)
  Object.keys(entities).forEach((entityKey: EntityName) => {
    const statePart = { ...state[entityKey] };

    let newIds: string[];

    if (entityKey === action.entityName) {
      newIds = result;
    } else {
      newIds = Object.keys(entities[entityKey]);
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
          ? (state[entityKey].by[options.name] || {})[options.key]
          : state[entityKey].list[options.name];

        if (previousList) {
          list = previousList;
        }
      }

      list = [...list, ...newIds];

      if (options.key) {
        statePart.by[options.name] = {
          ...(state[entityKey].by[options.name] || {}),
          [options.key]: list,
        };
      } else {
        statePart.list[options.name] = list;
      }
    }

    newState[entityKey] = statePart;
  });

  return newState;
}

function receiveUser(state: EntityStore, action) {
  const { user } = action.payload;

  if (!user || !user.player) {
    return state;
  }

  return receiveEntities(state, {
    entityName: EntityName.Player,
    payload: user.player,
  });
}

const reducer = createReducer(initialEntitiesState, {
  [actionNames.RECEIVE_ENTITIES]: receiveEntities,
  [actionNames.SET_USER_OR_LOGINROUTE]: receiveUser,
});

export default reducer;
