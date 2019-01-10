SELECT
    g.id,
    g.created_at,	
    g.name,
    COALESCE(g.image_url, "") as image_url
FROM groups g
JOIN group_players gp on g.id = gp.group_id
WHERE g.deleted_at is null
AND gp.player_id = ?