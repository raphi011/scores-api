SELECT
	p.id,
	p.first_name,
	p.last_name,
	p.birthday,
	p.gender,
	p.total_points,
	p.rank,
	p.club,
	p.country_union,
	p.license
FROM volleynet_players p
WHERE p.id = ?