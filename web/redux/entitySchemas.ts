import { denormalize, normalize, schema } from 'normalizr';

import { EntityName } from './../types';

const player = new schema.Entity('player');

const playerList = new schema.Array(player);

const tournament = new schema.Entity('tournament');

const tournamentList = new schema.Array(tournament);

const user = new schema.Entity('user');

const userList = new schema.Array(user);

const team = new schema.Entity(
  'team',
  {
    player1: player,
    player2: player,
  },
  {
    idAttribute: (t: any) => `${t.player1Id}${t.player2Id}`,
  },
);

const teamList = new schema.Array(team);

const match = new schema.Entity('match', {
  team1: team,
  team2: team,
});

const matchList = new schema.Array(match);

const statistic = new schema.Entity('statistic');

const statisticList = new schema.Array(statistic);

const group = new schema.Entity('group', {
  matches: matchList,
  players: playerList,
});

const groupList = new schema.Array(group);

const volleynetplayer = new schema.Entity('volleynetplayer');

const volleynetplayerList = new schema.Array(volleynetplayer);

const entitySchemaMap = {
  group,
  groupList,
  match,
  matchList,
  player,
  playerList,
  statistic,
  statisticList,
  team,
  teamList,
  tournament,
  tournamentList,
  user,
  userList,
  volleynetplayer,
  volleynetplayerList,
};

function getSchemaMap(entityName: string, isList: boolean) {
  const key = entityName + (isList ? 'List' : '');
  const entitySchema = entitySchemaMap[key];

  if (!entitySchema) {
    throw new Error(`Unknown schema: ${key}`);
  }

  return entitySchema;
}

function getArtificialId(): number {
  const id = Math.floor(Math.random() * Math.floor(100000));
  // TODO: fix this, id collisions possible
  return id;
}

function assignArtificialId(entities: any[] | object) {
  if (Array.isArray(entities)) {
    entities.forEach(e => {
      e.id = getArtificialId();
    });
  } else {
    entities.id = getArtificialId();
  }
}

export function norm(entityName: EntityName, data, assignId: boolean) {
  try {
    const isList = Array.isArray(data);

    const entitySchema = getSchemaMap(entityName, isList);

    if (assignId) {
      assignArtificialId(data);
    }

    const { result, entities } = normalize(data, entitySchema);

    return { result: isList ? result : [result], entities };
  } catch (e) {
    console.log(e);
  }
}

export function denorm(
  entityName: EntityName,
  entities,
  ids: number[] | number,
) {
  const isList = Array.isArray(ids);

  const entitySchema = getSchemaMap(entityName, isList);

  const result = denormalize(ids, entitySchema, entities);

  return result;
}
