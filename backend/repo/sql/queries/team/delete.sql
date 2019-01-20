UPDATE tournament_teams SET
	deleted_at = :deleted_at
WHERE volleynet_tournament_id = :tournament_id
    AND volleynet_player_1_id = :player1.id
    AND volleynet_player_2_id = :player2.id