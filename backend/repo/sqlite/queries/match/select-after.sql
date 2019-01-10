SELECT
	m.id,
	m.created_at,
	m.team1_player1_id,
	p1.name as team1_player1_name,
	COALESCE(u1.profile_image_url, "") as team1_player1_image_url,
	m.team1_player2_id,
	p2.name as team1_player2_name,
	COALESCE(u2.profile_image_url, "") as team1_player2_image_url,
	m.team2_player1_id,
	p3.name as team2_player1_name,
	COALESCE(u3.profile_image_url, "") as team2_player1_image_url,
	m.team2_player2_id,
	p4.name as team2_player2_name,
	COALESCE(u4.profile_image_url, "") as team2_player2_image_url,
	m.score_team1,
	m.score_team2,
	m.created_by_user_id,
	COALESCE(m.group_id, 0) as group_id
FROM matches m
JOIN players p1 on m.team1_player1_id = p1.id
JOIN players p2 on m.team1_player2_id = p2.id
JOIN players p3 on m.team2_player1_id = p3.id
JOIN players p4 on m.team2_player2_id = p4.id
LEFT JOIN users u1 on p1.user_id = u1.id
LEFT JOIN users u2 on p2.user_id = u2.id
LEFT JOIN users u3 on p3.user_id = u3.id
LEFT JOIN users u4 on p4.user_id = u4.id
WHERE m.deleted_at is null
AND m.created_at < ?
ORDER BY m.created_at DESC
LIMIT ?
