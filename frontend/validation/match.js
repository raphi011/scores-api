// @flow
import type { NewMatch } from '../types';

/* eslint-disable import/prefer-default-export */
export function validateMatch({
  player1Id,
  player2Id,
  player3Id,
  player4Id,
  scoreTeam1,
  scoreTeam2,
  targetScore,
}: NewMatch) {
  const errors = {};

  let higherScore = scoreTeam1;
  let lowerScore = scoreTeam2;

  if (lowerScore > higherScore) {
    [higherScore, lowerScore] = [lowerScore, higherScore];
  }

  if (!player1Id || !player2Id || !player3Id || !player4Id) {
    errors.all = "Four players have to be selected";
  }

  if (scoreTeam1 <= 0) {
    errors.scoreTeam1 = "Invalid score";
  }
  if (scoreTeam2 <= 0) {
    errors.scoreTeam2 = "Invalid score";
  }

  if (higherScore < targetScore) {
    errors.all = `One team has to score atleast ${targetScore}`;
  }

  if (
    (higherScore === targetScore && lowerScore === targetScore - 1) ||
    (higherScore > targetScore && lowerScore !== higherScore - 2)
  ) {
    errors.all = "Scores have to be two points apart";
  }

  errors.valid = Object.keys(errors).length === 0;

  return errors;
}
