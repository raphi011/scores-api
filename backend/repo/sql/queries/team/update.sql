UPDATE tournament_teams SET
	rank = :rank,
	seed = :seed,
	total_points = :total_points,
	won_points = :won_points,
	prize_money = :prize_money,
	deregistered = :deregistered
WHERE volleynet_tournament_id = :tournament_id
    AND volleynet_player_1_id = :player1.id
    AND volleynet_player_2_id = :player2.id