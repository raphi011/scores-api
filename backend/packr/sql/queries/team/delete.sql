DELETE FROM tournament_teams
WHERE
	tournament_id = ? AND
	player_1_id = ? AND
	player_2_id = ?