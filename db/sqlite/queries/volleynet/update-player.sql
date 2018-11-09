UPDATE volleynet_players SET
	updated_at = CURRENT_TIMESTAMP,
	first_name = ?,
	last_name = ?,
	birthday = ?,
	gender = ?,
	total_points = ?,
	rank = ?,
	club = ?,
	country_union = ?,
	license = ?
WHERE id = ?