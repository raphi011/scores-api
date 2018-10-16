CREATE TABLE db_version ( version integer NOT NULL );

CREATE TABLE users (
	id integer PRIMARY KEY AUTO_INCREMENT,
	created_at datetime NOT NULL,
	updated_at datetime,
	deleted_at datetime,
	role varchar(32) NOT NULL,
	volleynet_login varchar(64),
	volleynet_user_id integer,
	email varchar(255) NOT NULL UNIQUE,
	profile_image_url varchar(255) NOT NULL,
	salt blob,
	hash blob,
	iterations integer
);

CREATE TABLE players (
	id integer PRIMARY KEY AUTO_INCREMENT,
	created_at datetime NOT NULL,
	updated_at datetime,
	deleted_at datetime,
	name varchar(255) CHARSET utf8mb4 NOT NULL,
	user_id integer UNIQUE,
	FOREIGN KEY(user_id) REFERENCES users(id),
	INDEX(name)
);

CREATE TABLE groups (
	id integer PRIMARY KEY AUTO_INCREMENT,
	created_at datetime NOT NULL,
	updated_at datetime,
	deleted_at datetime,
	image_url varchar(255) NOT NULL,
	name varchar(255) CHARSET utf8mb4 NOT NULL,
	INDEX(name)
);

CREATE TABLE group_players (
	group_id integer,
	player_id integer,
	role varchar(32) NOT NULL,
	FOREIGN KEY(group_id) REFERENCES groups(id),
	FOREIGN KEY(player_id) REFERENCES players(id),
	PRIMARY KEY(group_id, player_id)
);

CREATE TABLE teams (
	created_at datetime NOT NULL,
	name varchar(255) CHARSET utf8mb4 NOT NULL,
	player1_id integer NOT NULL,
	player2_id integer NOT NULL,
	FOREIGN KEY(player1_id) REFERENCES players(id),
	FOREIGN KEY(player2_id) REFERENCES players(id),
	PRIMARY KEY (player1_id, player2_id),
	INDEX(name)
);

CREATE TABLE matches (
	id integer PRIMARY KEY AUTO_INCREMENT,
	created_at datetime NOT NULL,
	updated_at datetime,
	deleted_at datetime,
	team1_player1_id integer NOT NULL,
	team1_player2_id integer NOT NULL,
	team2_player1_id integer NOT NULL,
	team2_player2_id integer NOT NULL,
	score_team1 integer NOT NULL,
	score_team2 integer NOT NULL,
	created_by_user_id integer NOT NULL,
	group_id integer,
	FOREIGN KEY(group_id) REFERENCES groups(id),
	FOREIGN KEY(created_by_user_id) REFERENCES users(id),
	FOREIGN KEY(team1_player1_id, team1_player2_id)
	REFERENCES teams(player1_id, player2_id),
	FOREIGN KEY(team2_player1_id, team2_player2_id)
	REFERENCES teams(player1_id, player2_id)
);

CREATE VIEW team_statistics AS
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
	END AS points_won,
	CASE
		WHEN m.team1_player1_id = t.player1_id AND m.team1_player2_id = t.player2_id
		THEN m.score_team2
		ELSE m.score_team1
	END AS points_lost
FROM matches m
JOIN teams t
ON (m.team1_player1_id = t.player1_id AND m.team1_player2_id = t.player2_id)
OR (m.team2_player1_id = t.player1_id AND m.team2_player2_id = t.player2_id)
WHERE m.deleted_at IS NULL;

CREATE VIEW player_statistics AS
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
	END AS points_won,
	CASE
		WHEN m.team1_player1_id = p.id
			OR m.team1_player2_id = p.id THEN m.score_team2
		ELSE m.score_team1
	END AS points_lost,
	m.group_id
FROM matches m
JOIN players p ON 
	m.team1_player1_id = p.id OR
	m.team1_player2_id = p.id OR
	m.team2_player1_id = p.id OR
	m.team2_player2_id = p.id
WHERE m.deleted_at IS NULL;

CREATE TABLE volleynet_tournaments (
	id integer PRIMARY KEY,
	created_at datetime NOT NULL,
	updated_at datetime NOT NULL,
	gender varchar(1) NOT NULL,
	signedup_teams integer NOT NULL,
	start datetime NOT NULL,
	end datetime NOT NULL,
	name varchar(255) CHARSET utf8mb4 NOT NULL,
	league varchar(255) NOT NULL,
	link varchar(255) NOT NULL,
	entry_link varchar(255) NOT NULL,
	status varchar(255) NOT NULL,
	registration_open integer NOT NULL,
	location varchar(255) NOT NULL,
	live_scoring_link varchar(255) NOT NULL,
	html_notes text CHARSET utf8mb4 NOT NULL,
	mode varchar(64) NOT NULL,
	max_points integer NOT NULL,
	min_teams integer NOT NULL,
	max_teams integer NOT NULL,
	end_registration datetime NOT NULL,
	organiser varchar(128) NOT NULL,
	phone varchar(128) NOT NULL,
	email varchar(128) NOT NULL,
	web varchar(128) NOT NULL,
	current_points varchar(256) NOT NULL,
	season integer NOT NULL,
	loc_lat double NOT NULL,
	loc_lon double NOT NULL,
	INDEX(season),
	INDEX(gender),
	INDEX(league),
	INDEX(name),
	INDEX(start)
);

CREATE TABLE volleynet_players (
	id integer PRIMARY KEY,
	created_at datetime NOT NULL,
	updated_at datetime,
	first_name varchar(255) CHARSET utf8mb4 NOT NULL,
	last_name varchar(255) CHARSET utf8mb4 NOT NULL,
	total_points integer NOT NULL,
	rank integer NOT NULL,
	country_union varchar(255) NOT NULL,
	club varchar(255) NOT NULL,
	birthday date NOT NULL,
	license varchar(32) NOT NULL,
	gender varchar(1) NOT NULL,
	INDEX(first_name),
	INDEX(last_name),
	INDEX(birthday),
	INDEX(gender),
	INDEX(rank)
);

CREATE TABLE volleynet_tournament_teams (
	volleynet_tournament_id integer NOT NULL,
	volleynet_player_1_id integer NOT NULL,
	volleynet_player_2_id integer NOT NULL,
	rank integer NOT NULL,
	seed integer NOT NULL,
	total_points integer NOT NULL,
	won_points integer NOT NULL,
	prize_money real NOT NULL,
	deregistered integer NOT NULL,
	FOREIGN KEY(volleynet_tournament_id) REFERENCES volleynet_tournaments(id),
	FOREIGN KEY(volleynet_player_1_id) REFERENCES volleynet_players(id),
	FOREIGN KEY(volleynet_player_2_id) REFERENCES volleynet_players(id),
	PRIMARY KEY(volleynet_tournament_id, volleynet_player_1_id, volleynet_player_2_id)
);