UPDATE volleynet_tournament_teams SET
	rank = ?,
	seed = ?,
	total_points = ?,
	won_points = ?,
	prize_money = ?,
	deregistered = ?
WHERE volleynet_tournament_id = ?
    AND volleynet_player_1_id = ?
    AND volleynet_player_2_id = ?