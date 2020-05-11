UPDATE players SET
	updated_at = :updated_at,
	first_name = :first_name,
	last_name = :last_name,
	birthday = :birthday,
	gender = :gender,
	total_points = :total_points,
	ladder_rank = :ladder_rank,
	club = :club,
	country_union = :country_union,
	license = :license
WHERE id = :id