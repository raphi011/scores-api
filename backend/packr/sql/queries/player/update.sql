UPDATE players SET
	updated_at = :updatedAt,
	first_name = :firstname,
	last_name = :lastname,
	birthday = :birthday,
	gender = :gender,
	total_points = :totalpoints,
	rank = :rank,
	club = :club,
	country_union = :countryunion,
	license = :license
WHERE id = :id