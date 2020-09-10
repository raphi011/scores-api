CREATE TABLE players (
	id integer PRIMARY KEY,

	created_at datetime NOT NULL,
	updated_at datetime,
	deleted_at datetime,

	first_name varchar(128) NOT NULL,
	last_name varchar(128) NOT NULL,
	total_points integer NOT NULL,
	ladder_rank integer NOT NULL,
	country_union varchar(255) NOT NULL,
	club varchar(255) NOT NULL,
	birthday date,
	license varchar(32) NOT NULL,
	gender varchar(1) NOT NULL
);

CREATE TABLE users (
	id string PRIMARY KEY,

	created_at datetime NOT NULL,
	updated_at datetime,
	deleted_at datetime,

	email varchar(255) NOT NULL UNIQUE,
	profile_image_url varchar(255) NOT NULL,
	role varchar(32) NOT NULL,
	pw_iterations integer,
	pw_hash blob,
	pw_salt blob,

	player_id integer,
	player_login varchar(64),

	FOREIGN KEY(player_id) REFERENCES players(id)
);

CREATE TABLE tournaments (
	id integer PRIMARY KEY,

	created_at datetime NOT NULL,
	updated_at datetime,
	deleted_at datetime,

	current_points varchar(256) NOT NULL,
	email varchar(128) NOT NULL,
	end_date datetime NOT NULL,
	end_registration datetime,
	entry_link varchar(255) NOT NULL,
	gender varchar(16) NOT NULL,
	league varchar(128) NOT NULL,
	league_key varchar(128) NOT NULL,
	sub_league varchar(128) NOT NULL,
	sub_league_key varchar(128) NOT NULL,
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
	season varchar(16) NOT NULL,
	signedup_teams integer NOT NULL,
	start_date datetime NOT NULL,
	status varchar(255) NOT NULL,
	website varchar(128) NOT NULL,
	html_notes text NOT NULL
);


CREATE TABLE tournament_teams (
	tournament_id integer NOT NULL,
	player_1_id integer NOT NULL,
	player_2_id integer NOT NULL,

	created_at datetime NOT NULL,
	updated_at datetime,
	deleted_at datetime,

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

CREATE TABLE settings (
	created_at          datetime NOT NULL,
	updated_at          datetime,
	deleted_at          datetime,
    s_key               varchar(255) NOT NULL,
    s_value             varchar(255),
	s_type				varchar(255) NOT NULL,
    user_id             uuid REFERENCES users(id),

	PRIMARY KEY(s_key, user_id),
	FOREIGN KEY(user_id) REFERENCES users(id)
);