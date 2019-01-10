SELECT 
    s.player_id,
    COALESCE(u.profile_image_url, "") as profile_image,
    max(s.name) as name,
    cast((sum(s.won) / cast(count(1) as decimal) * 100) as unsigned) as percentage_won,
    sum(s.points_won) as points_won,
    sum(s.points_lost) as points_lost,
    count(1) as played,
    sum(s.won) as games_won,
    sum(1) - sum(s.won) as games_lost
FROM player_statistics s
JOIN players p ON s.player_id = p.id
LEFT JOIN users u ON p.user_id = u.id 
WHERE s.created_at > ?
AND s.group_id = ?
GROUP BY s.player_id 
ORDER BY percentage_won DESC