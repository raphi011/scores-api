SELECT
	s.created_at,
	s.updated_at,
    s.user_id,
    s.key,
    s.value,
    s.type
FROM settings s
WHERE s.user_id = ?
