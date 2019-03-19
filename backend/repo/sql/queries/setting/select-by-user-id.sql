SELECT
	s.created_at,
	s.updated_at,
    s.user_id,
    s.s_key,
    s.s_value,
    s.s_type
FROM settings s
WHERE s.user_id = ?
