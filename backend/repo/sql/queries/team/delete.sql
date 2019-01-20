DELETE FROM tournament_teams
WHERE
	tournament_id = :tournament_id AND
	player_1_id = :player1.id AND
	player_2_id = :player2.id