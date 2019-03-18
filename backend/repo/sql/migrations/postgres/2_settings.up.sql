CREATE TABLE settings (
	created_at  timestamptz NOT NULL,
	updated_at  timestamptz NOT NULL,
	deleted_at  timestamptz,
    key         text NOT NULL,
    value       text,
	type 		text NOT NULL,	
    user_id     int REFERENCES users(id),

	PRIMARY KEY (key, user_id)
);
