SELECT
    p.id,
    p.name,
    p.user_id,
    u.profile_image_url,
    p.created_at
FROM players p
LEFT JOIN users u on u.id = p.user_id 
JOIN group_players gp on gp.player_id = p.id
WHERE p.deleted_at is null
AND gp.group_id = ?	