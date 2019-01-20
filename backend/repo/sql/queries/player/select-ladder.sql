SELECT
	p.id,
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
WHERE p.ladder_rank > 0 AND p.gender = ? ORDER BY p.ladder_rank