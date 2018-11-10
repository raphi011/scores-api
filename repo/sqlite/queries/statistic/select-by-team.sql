SELECT 
    MAX(CASE WHEN s.player1_id = ? THEN s.player2_id ELSE s.player1_id END) AS player_id,
    COALESCE(MAX(CASE WHEN s.player1_id = ? THEN u2.profile_image_url ELSE u1.profile_image_url END), "") AS profile_image,
    MAX(CASE WHEN s.player1_id = ? THEN p2.name ELSE p1.name END) AS name,
    CAST((SUM(s.won) / CAST(COUNT(1) AS decimal) * 100) AS unsigned) AS percentage_won,
    SUM(s.points_won) AS points_won,
    SUM(s.points_lost) AS points_lost,
    COUNT(1) AS played,
    SUM(s.won) AS games_won,
    SUM(1) - SUM(s.won) AS games_lost
FROM team_statistics s
JOIN players p1 ON s.player1_id = p1.id
JOIN players p2 ON s.player2_id = p2.id
LEFT JOIN users u1 ON p1.user_id = u1.id 
LEFT JOIN users u2 ON p2.user_id = u2.id 
WHERE (s.player1_id = ? OR s.player2_id = ?) and s.created_at > ?
GROUP BY s.player1_id, s.player2_id 
ORDER BY percentage_won DESC