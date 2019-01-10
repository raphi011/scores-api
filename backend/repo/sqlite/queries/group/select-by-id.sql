SELECT
    g.id,
    g.created_at,	
    g.name,
    COALESCE(g.image_url, "") as image_url
FROM groups g
WHERE g.deleted_at is null AND g.id = ?