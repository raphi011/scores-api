import * as actionNames from '../actionNames';
import { norm } from '../entitySchemas';
import { createReducer } from '../reduxHelper';

import { EntityName, EntityType } from '../../types';
import { ReceiveEntitiesAction } from './actions';

export interface EntityStoreValues {
  [key: string]: EntityType;
}

export interface EntityStorePart {
  values: EntityStoreValues;
  list: {
    all: string[];
    [listName: string]: string[];
  };
  by?: {
    [listName: string]: { [key: string]: string[] };
  };
}

export type EntityStore = { [key in EntityName]: EntityStorePart };

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

function isEntityName(entity: string): entity is EntityName {
  return ['player', 'tournament', 'user'].indexOf(entity) !== -1;
}

function getList(
  state: EntityStore,
  entityName: EntityName,
  name: string,
  key?: string,
) {
  const statePart = state[entityName];

  if (!statePart) {
    return [];
  }

  if (key) {
    const lists = statePart.by && statePart.by[name];

    if (!lists || !lists[key]) {
      return [];
    }

    return lists[key];
  }

  return statePart.list[name];
}

function receiveEntities(state: EntityStore, action: ReceiveEntitiesAction) {
  // STEP 1: normalize entites
  const { entityName, payload, assignId = false, listOptions = {} } = action;

  const { entities, result } = norm(entityName, payload, assignId);

  const newState = { ...state };

  // STEP 2: add entities to entity map(s)
  for (const entityName in entities) {
    if (!isEntityName(entityName)) {
      continue;
    }

    let { values, by, list } = state[entityName];

    let newIds: string[];

    if (entityName === action.entityName) {
      newIds = result;
    } else {
      newIds = Object.keys(entityName);
    }

    values = {
      ...values,
      ...entities[entityName],
    };

    const options = listOptions[entityName];

    // STEP 3: append or replace ids
    if (options) {
      let newList: string[] = [];

      if (options.mode === 'append') {
        const previousList = getList(
          state,
          entityName,
          options.name,
          options.key,
        );

        if (previousList) {
          newList = previousList;
        }
      }

      newList = [...newList, ...newIds];

      if (options.key) {
        by = {
          ...by,
          [options.name]: {
            ...((by && by[options.name]) || {}),
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

    newState[entityName] = {
      by,
      list,
      values,
    };
  }

  return newState;
}

const reducer = createReducer(initialEntitiesState, {
  [actionNames.RECEIVE_ENTITIES]: receiveEntities,
});

export default reducer;
