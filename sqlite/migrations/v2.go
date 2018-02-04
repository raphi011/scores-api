package migrations

var V2 = []string{
	teamStatisticsView,
	groupTable,
	groupPlayerTable,
	alterMatchTable,
}

var ResetV2 = []string{
	"matches",
	"teams",
	"players",
	"users",
	"groups",
	"groupPlayers",
}

const (
	groupTable = `
		CREATE TABLE "groups" (
			"id" integer PRIMARY KEY autoincrement,
			"created_at" datetime NOT NULL,
			"updated_at" datetime,
			"deleted_at" datetime,
			"image_url" varchar(255),
			"name" varchar(255) NOT NULL
		)
	`

	groupPlayerTable = `
	 CREATE TABLE "groupPlayers" (
		 "group_id" integer,
		 "player_id" integer,
		 "role" varchar(32) NOT NULL,
			FOREIGN KEY(group_id) REFERENCES groups(id),
			FOREIGN KEY(player_id) REFERENCES players(id),
			PRIMARY KEY(group_id, player_id)
	 	)
	`

	alterMatchTable = `
		ALTER TABLE "matches"
			ADD "group_id" integer REFERENCES groups(id)
	`

	teamStatisticsView = `
		CREATE VIEW IF NOT EXISTS teamStatistics AS
		SELECT
			t.player1_id,
			t.player2_id,
			m.created_at,
			CASE
				WHEN
					(m.team1_player1_id = t.player1_id AND m.team1_player2_id = t.player2_id
						AND m.score_team1 > m.score_team2)
					OR 
					(m.team2_player1_id = t.player1_id AND m.team2_player2_id = t.player2_id
						AND m.score_team2 > m.score_team1)
				THEN 1
				ELSE 0
			END AS won,
			CASE
				WHEN m.team1_player1_id = t.player1_id AND m.team1_player2_id = t.player2_id
				THEN m.score_team1
				ELSE m.score_team2
			END AS pointsWon,
			CASE
				WHEN m.team1_player1_id = t.player1_id AND m.team1_player2_id = t.player2_id
				THEN m.score_team2
				ELSE m.score_team1
			END AS pointsLost
		FROM matches m
		JOIN teams t
		ON (m.team1_player1_id = t.player1_id AND m.team1_player2_id = t.player2_id)
		OR (m.team2_player1_id = t.player1_id AND m.team2_player2_id = t.player2_id)
		WHERE m.deleted_at IS NULL
	`
)
