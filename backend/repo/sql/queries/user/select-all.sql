SELECT
    u.id,
    u.created_at,
    u.email,
    u.hash,
    COALESCE(u.iterations, 0) as iterations,
    COALESCE(u.profile_image_url, '') as profile_image_url,
    u.role,
    u.salt,
    u.volleynet_user,
    u.volleynet_user_id
FROM users u