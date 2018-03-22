import { normalize, denormalize, schema } from 'normalizr';

import type { EntityName } from './../types';

const group = new schema.Entity('group');

const groupList = new schema.Array(group);

const player = new schema.Entity('player', { groups: groupList });

const playerList = new schema.Array(player);

const user = new schema.Entity('user', { player });

const userList = new schema.Array(user);

const team = new schema.Entity(
  'team',
  {
    player1: player,
    player2: player,
  },
  { idAttribute: t => Number.parseInt(`${t.player1Id}${t.player2Id}`, 10) },
);

const teamList = new schema.Array(team);

const match = new schema.Entity('match', {
  team1: team,
  team2: team,
});

const matchList = new schema.Array(match);

const statistic = new schema.Entity('statistic');

const statisticList = new schema.Array(statistic);

const entitySchemaMap = {
  group,
  groupList,
  user,
  userList,
  player,
  playerList,
  team,
  teamList,
  match,
  matchList,
  statistic,
  statisticList,
};

function getSchemaMap(entityName: string, isList: boolean) {
  const key = entityName + (isList ? 'List' : '');
  const entitySchema = entitySchemaMap[key];

  if (!entitySchema) throw new Error(`Unknown schema: ${key}`);

  return entitySchema;
}

function getArtificialId(): number {
  const id = Math.floor(Math.random() * Math.floor(100000));
  // TODO: fix this, id collisions possible
  return id;
}

function assignArtificialId(entities: Array<any> | Object) {
  if (Array.isArray(entities)) {
    entities.forEach(e => {
      e.id = getArtificialId();
    });
  } else {
    entities.id = getArtificialId();
  }
}

export function norm(entityName: EntityName, data, assignId: boolean) {
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
  ids: Array<number> | number,
) {
  const isList = Array.isArray(ids);

  const entitySchema = getSchemaMap(entityName, isList);

  const result = denormalize(ids, entitySchema, entities);

  return result;
}
