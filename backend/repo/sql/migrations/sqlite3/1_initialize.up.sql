CREATE TABLE users (
	id integer PRIMARY KEY autoincrement,
	created_at datetime NOT NULL,
	deleted_at datetime,
	email varchar(255) NOT NULL UNIQUE,
	profile_image_url varchar(255) NOT NULL,
	role varchar(32) NOT NULL,
	pw_iterations integer,
	pw_hash blob,
	pw_salt blob,
	updated_at datetime,
	volleynet_user varchar(64),
	volleynet_user_id integer
);

CREATE TABLE tournaments (
	id integer PRIMARY KEY,
	created_at datetime NOT NULL,
	current_points varchar(256) NOT NULL,
	email varchar(128) NOT NULL,
	end_date datetime NOT NULL,
	end_registration datetime NOT NULL,
	entry_link varchar(255) NOT NULL,
	gender varchar(16) NOT NULL,
	league varchar(128) NOT NULL,
	link varchar(255) NOT NULL,
	live_scoring_link varchar(255) NOT NULL,
	loc_lat double NOT NULL,
	loc_lon double NOT NULL,
	location varchar(255) NOT NULL,
	max_points integer NOT NULL,
	max_teams integer NOT NULL,
	min_teams integer NOT NULL,
	mode varchar(64) NOT NULL,
	name varchar(128) NOT NULL,
	organiser varchar(128) NOT NULL,
	phone varchar(128) NOT NULL,
	registration_open integer NOT NULL,
	season integer NOT NULL,
	signedup_teams integer NOT NULL,
	start_date datetime NOT NULL,
	status varchar(255) NOT NULL,
	updated_at datetime NOT NULL,
	website varchar(128) NOT NULL,
	html_notes text NOT NULL
);

CREATE TABLE players (
	id integer PRIMARY KEY,
	created_at datetime NOT NULL,
	updated_at datetime,
	first_name varchar(128) NOT NULL,
	last_name varchar(128) NOT NULL,
	total_points integer NOT NULL,
	ladder_rank integer NOT NULL,
	country_union varchar(255) NOT NULL,
	club varchar(255) NOT NULL,
	birthday date NOT NULL,
	license varchar(32) NOT NULL,
	gender varchar(1) NOT NULL
);

CREATE TABLE tournament_teams (
	tournament_id integer NOT NULL,
	player_1_id integer NOT NULL,
	player_2_id integer NOT NULL,
	result integer NOT NULL,
	seed integer NOT NULL,
	total_points integer NOT NULL,
	won_points integer NOT NULL,
	prize_money real NOT NULL,
	deregistered integer NOT NULL,
	FOREIGN KEY(tournament_id) REFERENCES tournaments(id),
	FOREIGN KEY(player_1_id) REFERENCES players(id),
	FOREIGN KEY(player_2_id) REFERENCES players(id),
	PRIMARY KEY(tournament_id, player_1_id, player_2_id)
);