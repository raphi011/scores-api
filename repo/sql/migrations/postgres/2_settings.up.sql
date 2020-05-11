CREATE TABLE settings (
	created_at  timestamptz NOT NULL,
	updated_at  timestamptz,
	deleted_at  timestamptz,
    s_key       text NOT NULL,
    s_value     text,
	s_type 		text NOT NULL,	
    user_id     int REFERENCES users(id),

	PRIMARY KEY (s_key, user_id)
);
