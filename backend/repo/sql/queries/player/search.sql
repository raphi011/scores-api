SELECT
	p.id,
	p.created_at,
	p.updated_at,
	p.first_name,
	p.last_name,
	p.birthday,
	p.gender,
	p.total_points,
	p.ladder_rank,
	p.club,
	p.country_union,
	p.license
FROM players p
WHERE
    (? = '' OR p.first_name LIKE ?) AND
    (? = '' OR p.last_name LIKE ?) AND
    (? = '' OR p.gender = ?)
