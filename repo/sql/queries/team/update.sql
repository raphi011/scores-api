UPDATE tournament_teams SET
	updated_at = :updated_at,
	result = :result,
	seed = :seed,
	total_points = :total_points,
	won_points = :won_points,
	prize_money = :prize_money,
	deregistered = :deregistered
WHERE tournament_id = :tournament_id
    AND player_1_id = :player1.id
    AND player_2_id = :player2.id