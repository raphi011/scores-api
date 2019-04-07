UPDATE settings SET
	updated_at = :updated_at,
	s_value = :s_value
WHERE
	user_id = :user_id AND
	s_key = :s_key
