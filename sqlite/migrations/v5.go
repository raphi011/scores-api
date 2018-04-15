package migrations

var V5 = []string{
	dropPlayerStatisticsView,
	playerStatisticsGroupID,
}

var ResetV5 = ResetV4

const (
	dropPlayerStatisticsView = `DROP VIEW "playerStatistics"`

	playerStatisticsGroupID = `
		CREATE VIEW "playerStatistics" AS
		SELECT
			p.id as player_id,
			p.name,
			m.created_at,
			CASE
				WHEN
					(m.team1_player1_id = p.id OR m.team1_player2_id = p.id)
					AND (m.score_team1 > m.score_team2)
					OR (m.team2_player1_id = p.id OR m.team2_player2_id = p.id)
					AND (m.score_team2 > m.score_team1)
				THEN 1
				ELSE 0
			END AS won,
			CASE
				WHEN m.team1_player1_id = p.id OR m.team1_player2_id = p.id THEN m.score_team1
				ELSE m.score_team2
			END AS pointsWon,
			CASE
				WHEN m.team1_player1_id = p.id
					OR m.team1_player2_id = p.id THEN m.score_team2
				ELSE m.score_team1
			END AS pointsLost,
			m.group_id
		FROM matches m
		JOIN players p ON 
			m.team1_player1_id = p.id OR
			m.team1_player2_id = p.id OR
			m.team2_player1_id = p.id OR
			m.team2_player2_id = p.id
		WHERE m.deleted_at IS NULL
	`
)
