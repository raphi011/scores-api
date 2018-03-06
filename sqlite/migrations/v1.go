package migrations

var V1 = []string{
	userTable,
	playerTable,
	teamTable,
	matchTable,
	playerStatisticsView,
}

var ResetV1 = []string{
	"matches",
	"teams",
	"players",
	"users",
}

const (
	matchTable = `
		CREATE TABLE "matches" (
			"id" integer PRIMARY KEY autoincrement,
			"created_at" datetime NOT NULL,
			"updated_at" datetime,
			"deleted_at" datetime,
			"team1_player1_id" integer NOT NULL,
			"team1_player2_id" integer NOT NULL,
			"team2_player1_id" integer NOT NULL,
			"team2_player2_id" integer NOT NULL,
			"score_team1" integer NOT NULL,
			"score_team2" integer NOT NULL,
			"created_by_user_id" integer NOT NULL,
			FOREIGN KEY(created_by_user_id) REFERENCES users(id),
			FOREIGN KEY(team1_player1_id, team1_player2_id)
			REFERENCES teams(player1_id, player2_id),
			FOREIGN KEY(team2_player1_id, team2_player2_id)
			REFERENCES teams(player1_id, player2_id)
		)
	`

	playerTable = `
		CREATE TABLE "players" (
			"id" integer PRIMARY KEY autoincrement,
			"created_at" datetime NOT NULL,
			"updated_at" datetime,
			"deleted_at" datetime,
			"name" varchar(255) NOT NULL,
			"user_id" integer,
			FOREIGN KEY(user_id) REFERENCES users(id)
		)
	`

	teamTable = `
		CREATE TABLE "teams" (
			"created_at" datetime NOT NULL,
			"name" varchar(255),
			"player1_id" integer NOT NULL,
			"player2_id" integer NOT NULL,
			FOREIGN KEY(player1_id) REFERENCES players(id),
			FOREIGN KEY(player2_id) REFERENCES players(id),
			PRIMARY KEY (player1_id, player2_id)
		)
	`

	userTable = `
		CREATE TABLE "users" (
			"id" integer PRIMARY KEY autoincrement,
			"created_at" datetime NOT NULL,
			"updated_at" datetime,
			"deleted_at" datetime,
			"email" varchar(255) NOT NULL UNIQUE,
			"profile_image_url" varchar(255)
		)
	`

	playerStatisticsView = `
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
			END AS pointsLost
		FROM matches m
		JOIN players p ON 
			m.team1_player1_id = p.id OR
			m.team1_player2_id = p.id OR
			m.team2_player1_id = p.id OR
			m.team2_player2_id = p.id
		WHERE m.deleted_at IS NULL
	`
)
