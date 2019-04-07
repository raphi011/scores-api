CREATE TABLE settings (
	created_at          datetime NOT NULL,
	updated_at          datetime,
	deleted_at          datetime,
    s_key               varchar(255) NOT NULL,
    s_value             varchar(255),
	s_type				varchar(255) NOT NULL,
    user_id             int REFERENCES users(id),

	PRIMARY KEY(s_key, user_id),
	FOREIGN KEY(user_id) REFERENCES users(id)
);
