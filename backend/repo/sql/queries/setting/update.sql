UPDATE settings SET
	updated_at = :updated_at,
	s_value = :s_value
WHERE user_id = :user_id and s_key = :s_key
