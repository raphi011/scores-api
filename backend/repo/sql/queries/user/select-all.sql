SELECT
    u.id,
    u.email,
    COALESCE(u.profile_image_url, "") as profile_image_url,
    COALESCE(p.id, 0) as player_id,
    u.created_at,
    u.salt,
    u.hash,
    COALESCE(u.iterations, 0) as iterations,
    u.volleynet_user_id,
    u.volleynet_login,
    u.role
FROM users u
LEFT JOIN players p on u.id = p.user_id