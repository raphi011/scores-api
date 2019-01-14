SELECT
	t.volleynet_tournament_id,
	t.volleynet_player_1_id,
	p1.first_name,
	p1.last_name,
	p1.total_points,
	p1.country_union,
	p1.birthday,
	p1.license,
	p1.gender,
	t.volleynet_player_2_id,
	p2.first_name,
	p2.last_name,
	p2.total_points,
	p2.country_union,
	p2.birthday,
	p2.license,
	p2.gender,
	t.rank,
	t.seed,
	t.total_points,
	t.won_points,
	t.prize_money,
	t.deregistered
FROM volleynet_tournament_teams t
JOIN players p1 on p1.id = t.volleynet_player_1_id
JOIN players p2 on p2.id = t.volleynet_player_2_id
WHERE t.volleynet_tournament_id = ?