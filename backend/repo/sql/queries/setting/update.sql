UPDATE settings SET
	updated_at = :updated_at,
	value = :value
WHERE user_id = :user_id and key = :key
