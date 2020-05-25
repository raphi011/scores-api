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
WHERE p.id = ?