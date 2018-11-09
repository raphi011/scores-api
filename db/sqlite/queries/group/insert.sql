INSERT INTO groups
(
	created_at,
	name,
	image_url
)
VALUES
(
	CURRENT_TIMESTAMP,
	?,
	?
)