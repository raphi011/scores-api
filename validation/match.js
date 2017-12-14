// @flow
import type { Match } from '../types';

/* eslint-disable import/prefer-default-export */
export function validateMatch({
  /* player1ID,
  player2ID,
  player3ID,
  player4ID, */
  scoreTeam1,
  scoreTeam2,
  targetScore
}: Match) {
  const errors = {};

  let higherScore = scoreTeam1;
  let lowerScore = scoreTeam2;

  if (lowerScore > higherScore) {
    [higherScore, lowerScore] = [lowerScore, higherScore];
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
