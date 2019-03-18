INSERT INTO settings 
(
	created_at,
	updated_at,
    user_id,
    key,
    value,
    type
)
VALUES
(
	:created_at,
	:updated_at,
    :user_id,
    :key,
    :value,
    :type
)
