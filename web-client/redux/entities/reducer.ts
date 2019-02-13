import * as actionNames from '../actionNames';
import { norm } from '../entitySchemas';
import { createReducer } from '../reduxHelper';

import { EntityName, EntityType } from '../../types';

export interface EntityStoreValues {
  [key: string]: EntityType;
}

export interface EntityStorePart {
  values: EntityStoreValues;
  list?: {
    all: string[];
    [listName: string]: string[];
  };
  by?: {
    [listName: string]: { [key: string]: string[] };
  };
}

export type EntityStore = { [key in EntityName]: EntityStorePart };

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
  payload: object | string;
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
    let { values, by, list } = state[entityKey];

    let newIds: string[];

    if (entityKey === action.entityName) {
      newIds = result;
    } else {
      newIds = Object.keys(entities[entityKey]);
    }

    values = {
      ...values,
      ...entities[entityKey],
    };

    const options = listOptions[entityKey];

    // STEP 3: append or replace ids
    if (options) {
      let newList = [];

      if (options.mode === 'append') {
        const previousList = options.key
          ? (state[entityKey].by[options.name] || {})[options.key]
          : state[entityKey].list[options.name];

        if (previousList) {
          newList = previousList;
        }
      }

      newList = [...newList, ...newIds];

      if (options.key) {
        by = {
          ...by,
          [options.name]: {
            ...(by[options.name] || {}),
            [options.key]: newList,
          },
        };
      } else {
        list = {
          ...list,
          [options.name]: newList,
        };
      }
    }

    newState[entityKey] = {
      by,
      list,
      values,
    };
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
