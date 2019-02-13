/* eslint-disable @typescript-eslint/no-explicit-any */

import { denormalize, normalize, schema } from 'normalizr';

import { EntityName, EntityType } from './../types';

const player = new schema.Entity('player');
const playerList = new schema.Array(player);

const tournament = new schema.Entity('tournament');
const tournamentList = new schema.Array(tournament);

const user = new schema.Entity('user');
const userList = new schema.Array(user);

const entitySchemaMap = {
  player,
  playerList,
  tournament,
  tournamentList,
  user,
  userList,
};

function getSchemaMap(entityName: string, isList: boolean) {
  const key = entityName + (isList ? 'List' : '');
  const entitySchema = entitySchemaMap[key];

  if (!entitySchema) {
    throw new Error(`Unknown schema: ${key}`);
  }

  return entitySchema;
}

function getArtificialId() {
  const id = Math.floor(Math.random() * Math.floor(100000));
  return id.toString();
}

function assignArtificialId(entities: any[] | any) {
  if (Array.isArray(entities)) {
    entities.forEach(e => {
      e.id = getArtificialId();
    });
  } else {
    entities.id = getArtificialId();
  }
}

export function norm(
  entityName: EntityName,
  data: any,
  assignId: boolean,
): {
  result: string[];
  entities: { [entity in EntityName]?: { [id: string]: EntityType } };
} {
  const isList = Array.isArray(data);

  const entitySchema = getSchemaMap(entityName, isList);

  if (assignId) {
    assignArtificialId(data);
  }

  const { result, entities } = normalize(data, entitySchema);

  return { result: isList ? result : [result], entities };
}

export function denorm(
  entityName: EntityName,
  entities,
  ids: string[] | string,
) {
  const isList = Array.isArray(ids);

  const entitySchema = getSchemaMap(entityName, isList);

  const result = denormalize(ids, entitySchema, entities);

  return result;
}
