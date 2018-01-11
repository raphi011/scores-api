package sqlite

import (
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

func Open(filename string) (*sql.DB, error) {

	db, err := sql.Open("sqlite3", filename)

	if err != nil {
		return nil, err
	}

	err = db.QueryRow("SELECT name FROM sqlite_master WHERE type='table' AND name='dbVersion'").Scan()

	if err == sql.ErrNoRows {
		log.Println("Creating sqlite tables")
		err := execMultiple(
			db,
			versionTable,
			userTable,
			playerTable,
			teamTable,
			matchTable,
			playerStatisticsView,
			teamStatisticsView)

		if err != nil {
			return nil, err
		}
	}

	return db, nil
}

func ClearTables(db *sql.DB) error {
	err := execMultiple(db,
		"DELETE FROM matches",
		"DELETE FROM teams",
		"DELETE FROM players",
		"DELETE FROM users",
	)

	return err
}

// func seedDb() {
// 		player1 := Player{Name: "Raphi"}
// 		player2 := Player{Name: "Robert"}
// 		player3 := Player{Name: "Lukas"}
// 		player4 := Player{Name: "Richie"}
// 		player5 := Player{Name: "Dominik"}
// 		player6 := Player{Name: "Roman"}

// 		user1 := User{Email: "raphi011@gmail.com", Player: player1}
// 		user2 := User{Email: "", Player: player2}
// 		user3 := User{Email: "", Player: player3}
// 		user4 := User{Email: "Rb1@outlook.at", Player: player4}
// 		user5 := User{Email: "Rieder.dominik@gmail.com", Player: player5}
// 		user6 := User{Email: "", Player: player6}

// 		a.Db.Create(&user1)
// 		a.Db.Create(&user2)
// 		a.Db.Create(&user3)
// 		a.Db.Create(&user4)
// 		a.Db.Create(&user5)
// 		a.Db.Create(&user6)
// }

func execMultiple(db *sql.DB, statements ...string) error {
	for _, statement := range statements {
		_, err := db.Exec(statement)
		if err != nil {
			return err
		}
	}

	return nil
}

const (
	versionTable = `
		CREATE TABLE "dbVersion" (
			"version" integer NOT NULL
		)
	`

	matchTable = `
		CREATE TABLE "matches" (
			"id" integer PRIMARY KEY autoincrement,
			"created_at" datetime NOT NULL,
			"updated_at" datetime,
			"deleted_at" datetime,
			"team1_id" integer NOT NULL,
			"team2_id" integer NOT NULL,
			"score_team1" integer NOT NULL,
			"score_team2" integer NOT NULL,
			"created_by_user_id" integer NOT NULL,
			FOREIGN KEY(created_by_user_id) REFERENCES user(id),
			FOREIGN KEY(team1_id) REFERENCES team(id),
			FOREIGN KEY(team2_id) REFERENCES team(id)
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
			"id" integer PRIMARY KEY autoincrement,
			"created_at" datetime NOT NULL,
			"updated_at" datetime,
			"deleted_at" datetime,
			"name" varchar(255),
			"player1_id" integer NOT NULL,
			"player2_id" integer NOT NULL,
			FOREIGN KEY(player1_id) REFERENCES players(id),
			FOREIGN KEY(player2_id) REFERENCES players(id)
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
		CREATE VIEW IF NOT EXISTS playerStatistics AS
		SELECT
			p.id,
			p.name,
			m.created_at,
			CASE
				WHEN
					(t1.player1_id = p.id	OR t1.player2_id = p.id)
					AND (m.score_team1 > m.score_team2)
					OR (t2.player1_id = p.id OR t2.player2_id = p.id)
					AND (m.score_team2 > m.score_team1)
				THEN 1
				ELSE 0
			END AS won,
			CASE
				WHEN t1.player1_id = p.id	OR t1.player2_id = p.id THEN m.score_team1
				ELSE m.score_team2
			END AS pointsWon,
			CASE
				WHEN t1.player1_id = p.id
					OR t1.player2_id = p.id THEN m.score_team2
				ELSE m.score_team1
			END AS pointsLost
		FROM matches m
		JOIN teams t1 ON m.team1_id = t1.id
		JOIN teams t2 ON m.team2_id = t2.id
		JOIN players p ON t1.player1_id = p.id
		OR t1.player2_id = p.id
		OR t2.player1_id = p.id
		OR t2.player2_id = p.id
		WHERE m.deleted_at IS NULL
	`
	teamStatisticsView = `
		CREATE VIEW IF NOT EXISTS teamStatistics AS
		SELECT t.id AS team_id,
			t.player1_id,
			t.player2_id,
			m.id AS match_id,
			CASE
				WHEN
					(m.team1_id = t.id AND m.score_team1 > m.score_team2)
					OR (m.team2_id = t.id
					AND m.score_team2 > m.score_team1)
				THEN 1
				ELSE 0
			END AS won,
			CASE
				WHEN m.team1_id = t.id THEN m.score_team1
				ELSE m.score_team2
			END AS pointsWon,
			CASE
				WHEN m.team1_id = t.id
				THEN m.score_team2
				ELSE m.score_team1
			END AS pointsLost
		FROM matches m
		JOIN teams t ON m.team1_id = t.id
		OR m.team2_id = t.id
		WHERE m.deleted_at IS NULL
	`
)
