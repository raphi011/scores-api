UPDATE tournament_teams SET
	rank = :rank,
	seed = :seed,
	total_points = :totalpoints,
	won_points = :wonpoints,
	prize_money = :prizemoney,
	deregistered = :deregistered
WHERE volleynet_tournament_id = :tournamentid
    AND volleynet_player_1_id = :player1Id
    AND volleynet_player_2_id = :player2Id