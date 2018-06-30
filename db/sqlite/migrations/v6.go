package migrations

var V6 = []string{
	alterUserTable,
	volleynetTournaments,
	volleynetPlayers,
	volleynetTournamentTeam,
}

var ResetV6 = ResetV5 // TODO

const (
	alterUserTable = `
		ALTER TABLE "users" ADD COLUMN "role" varchar(32) NOT NULL DEFAULT "user";
		ALTER TABLE "users" ADD COLUMN "volleynet_login" varchar(64) NOT NULL DEFAULT "";
		ALTER TABLE "users" ADD COLUMN "volleynet_user_id" integer NOT NULL DEFAULT 0;
	`
	volleynetTournaments = `
		CREATE TABLE "volleynetTournaments" (
			"id" integer PRIMARY KEY,
			"created_at" datetime NOT NULL,
			"updated_at" datetime NOT NULL,
			"gender" varchar(1) NOT NULL,
			"start" datetime NOT NULL,
			"end" datetime NOT NULL,
			"name" varchar(255) NOT NULL,
			"league" varchar(255) NOT NULL,
			"link" varchar(255) NOT NULL,
			"entry_link" varchar(255) NOT NULL,
			"status" varchar(255) NOT NULL,
			"registration_open" integer NOT NULL,
			"location" varchar(255) NOT NULL,
			"live_scoring_link" varchar(255) NOT NULL,
			"html_notes" varchar NOT NULL,
			"mode" varchar(64) NOT NULL,
			"max_points" integer NOT NULL,
			"min_teams" integer NOT NULL,
			"max_teams" integer NOT NULL,
			"end_registration" datetime NOT NULL,
			"organiser" varchar(128) NOT NULL,
			"phone" varchar(128) NOT NULL,
			"email" varchar(128) NOT NULL,
			"web" varchar(128) NOT NULL,
			"current_points" integer NOT NULL,
			"season" integer NOT NULL,
			"loc_lat" double NOT NULL,
			"loc_lon" double NOT NULL
		)
	`

	volleynetPlayers = `
		CREATE TABLE "volleynetPlayers" (
			"id" integer PRIMARY KEY,
			"created_at" datetime NOT NULL,
			"updated_at" datetime,
			"first_name" varchar(255) NOT NULL,
			"last_name" varchar(255) NOT NULL,
			"total_points" integer NOT NULL,
			"rank" integer NOT NULL,
			"country_union" varchar(255) NOT NULL,
			"club" varchar(255) NOT NULL,
			"birthday" date NOT NULL,
			"license" varchar(32) NOT NULL,
			"gender" varchar(1) NOT NULL,
			"login" varchar(255) NOT NULL
		)
	`

	volleynetTournamentTeam = `
		CREATE TABLE "volleynetTournamentTeams" (
			"volleynet_tournament_id" integer NOT NULL,
			"volleynet_player_1_id" integer NOT NULL,
			"volleynet_player_2_id" integer NOT NULL,
			"rank" integer NOT NULL,
			"seed" integer NOT NULL,
			"total_points" integer NOT NULL,
			"won_points" integer NOT NULL,
			"prize_money" real NOT NULL,
            "deregistered" integer NOT NULL,
			FOREIGN KEY(volleynet_tournament_id) REFERENCES volleynetTournaments(id),
			FOREIGN KEY(volleynet_player_1_id) REFERENCES volleynetPlayers(id),
			FOREIGN KEY(volleynet_player_2_id) REFERENCES volleynetPlayers(id),
			PRIMARY KEY(volleynet_tournament_id, volleynet_player_1_id, volleynet_player_2_id)
	 	)
	`
)
