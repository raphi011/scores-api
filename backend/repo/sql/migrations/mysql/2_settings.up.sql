CREATE TABLE settings (
	created_at          datetime NOT NULL,
	updated_at          datetime NOT NULL,
	deleted_at          datetime,
    key                 varchar(255) NOT NULL,
    value               varchar(255),
	type				varchar(255) NOT NULL,
    user_id             int REFERENCES users(id),

	PRIMARY KEY(key, user_id),
	FOREIGN KEY(user_id) REFERENCES users(id)
);
